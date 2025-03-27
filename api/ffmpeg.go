package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/fsnotify/fsnotify"
)

type Ffmpeg struct {
	Context    context.Context
	CameraName string
	Preset     ConfigPreset
	W          http.ResponseWriter
	Audio      bool
}

func (this *Ffmpeg) Live() error {
	if this.Preset.Fps == nil && this.Preset.Resolution == nil {
		return this.generateVideo()
	}

	return this.generateImageStream()
}

func (this *Ffmpeg) streamNewJpeg(filePath string) error {
	_, err := this.W.Write([]byte(""))
	if err != nil {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(this.W, file)
	if err != nil {
		return err
	}

	_, err = this.W.Write([]byte("\r\n--boundarydonotcross\r\nContent-Type: image/jpeg\r\n\r\n"))
	if err != nil {
		return err
	}

	flusher, ok := this.W.(http.Flusher)
	if !ok {
		return errors.New("failed converting response write to flusher")
	}
	flusher.Flush()

	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}

func (this *Ffmpeg) watchAndStreamImageFiles(dir string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	err = watcher.Add(dir)
	if err != nil {
		return err
	}

	log.Printf("watching for new files in: %s", dir)

	_, err = this.W.Write([]byte("--boundarydonotcross\r\nContent-Type: image/jpeg\r\n\r\n"))
	if err != nil {
		return err
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return errors.New("watcher events not ok")
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				info, err := os.Stat(event.Name)
				if err == nil && !info.IsDir() {
					log.Printf("new image found for streaming: %s", event.Name)
					err = this.streamNewJpeg(event.Name)
					if err != nil {
						return err
					}
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return errors.New("watcher errors not ok")
			}
			return err
		case <-this.Context.Done():
			return errors.New("context closed early")
		}
	}
}

func (this *Ffmpeg) generateImageStream() error {
	this.W.Header().Add("Content-Type", "multipart/x-mixed-replace; boundary=boundarydonotcross")
	this.W.Header().Add("Cache-Control", "no-cache")
	this.W.Header().Add("Transfer-Encoding", "chunked")

	if this.Audio {
		log.Printf("adding audio to image stream is not possible")
	}
	args := make([]string, 0)
	args = append(args, "-timeout", "5", "-rtsp_transport",
		"tcp", "-i", this.Preset.Stream)
	if this.Preset.Fps != nil {
		args = append(args, "-vf", fmt.Sprintf("fps=%.3f", *this.Preset.Fps))
	}
	args = append(args, "%6d.jpeg")

	cmd := exec.CommandContext(this.Context, "ffmpeg", args...)
	tmpDir, err := os.MkdirTemp("", "teo-api-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)
	log.Println("tmp directory is %s", tmpDir)

	cmd.Dir = tmpDir
	_, err = cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(stderr)
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			log.Printf("%s: %s", this.CameraName, line)
		}
	}()

	log.Printf("starting ffmpeg (%s)", cmd.String())
	if err = cmd.Start(); err != nil {
		return err
	}

	go func() {
		err := this.watchAndStreamImageFiles(tmpDir)
		log.Println(err)
	}()

	return cmd.Wait()
}

func (this *Ffmpeg) generateVideo() error {
	this.W.Header().Add("Content-Type", "video/mp4")

	args := make([]string, 0)
	args = append(args,
		"-timeout", "5", "-rtsp_transport", "tcp",
		"-i", this.Preset.Stream, "-c:v", "copy")
	if this.Audio == true {
		args = append(args, "-c:a", "copy")
	} else {
		args = append(args, "-an")
	}
	args = append(args,
		"-f", "mp4", "-movflags", "+faststart+frag_keyframe+empty_moov",
		"-",
	)
	cmd := exec.CommandContext(this.Context, "ffmpeg", args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	log.Printf("starting ffmpeg (%s)", cmd.String())
	if err = cmd.Start(); err != nil {
		return err
	}

	go func() {
		_, err := io.Copy(this.W, stdout)
		if err != nil {
			log.Printf("%s: %s", this.CameraName, err.Error())
		}
	}()

	scanner := bufio.NewScanner(stderr)

	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			log.Printf("%s: %s", this.CameraName, line)
		}
	}()

	return cmd.Wait()
}
