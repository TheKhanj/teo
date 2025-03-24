package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

type ApiController struct {
	config *Config
}

func (this *ApiController) AddRoutes(router *httprouter.Router) {
	router.GET("/:camera/live", this.Live)
}

func runFfmpegLiveView(
	ctx context.Context,
	cam string, url string, output io.Writer,
) error {
	cmd := exec.CommandContext(
		ctx,
		"ffmpeg", "-timeout", "5", "-rtsp_transport", "tcp",
		"-i", url, "-c", "copy", "-f", "mp4", "-movflags", "+faststart+frag_keyframe+empty_moov",
		"-",
	)

	log.Printf("starting ffmpeg (%s)", cmd.String())

	stdout, err := cmd.StdoutPipe()
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}

	go func() {
		_, err := io.Copy(output, stdout)
		if err != nil {
			log.Printf("%s: %s", cam, err.Error())
		}
	}()

	scanner := bufio.NewScanner(stderr)

	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			log.Printf("%s: %s", cam, line)
		}
	}()

	return cmd.Wait()
}

func (this *ApiController) Live(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	cameraName := ps.ByName("camera")
	cam, ok := this.config.Cameras[cameraName]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "video/mp4")
	err := runFfmpegLiveView(r.Context(), cameraName, cam.Primary, w)
	if err != nil {
		log.Println(err)
	}
}
