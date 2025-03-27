package main

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
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
	router.GET("/cameras/", this.auth(this.Cameras))
	router.GET("/cameras/:camera/live", this.auth(this.CamerasLive))
}

func runFfmpegLiveView(
	ctx context.Context,
	cam string, preset ConfigPreset, output io.Writer,
) error {
	cmd := exec.CommandContext(
		ctx,
		"ffmpeg", "-timeout", "5", "-rtsp_transport", "tcp",
		"-i", preset.Stream, "-c", "copy", "-f", "mp4", "-movflags", "+faststart+frag_keyframe+empty_moov",
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

func (this *ApiController) getPreset(
	camera *ConfigCamera, preset string,
) (ConfigPreset, error) {
	ret := ConfigPreset{}

	if preset == "primary" {
		ret.Stream = camera.Primary
	} else if preset == "secondary" {
		ret.Stream = camera.Secondary
	} else {
		p, presetExists := (*this.config.Api.Presets)[preset]
		if !presetExists {
			return ret, fmt.Errorf("preset %s does not exist", preset)
		}
		ret = p
	}

	return ret, nil
}

func (this *ApiController) auth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		username, password, ok := r.BasicAuth()
		passwordHash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

		isInvalid := !ok || this.config.Users == nil

		if !isInvalid {
			_, userExists := (*this.config.Users)[username]
			isInvalid = isInvalid || !userExists
		}

		if isInvalid ||
			(*this.config.Users)[username].Password != passwordHash {
			w.Header().Set("WWW-Authenticate", `Basic realm="Api"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r, ps)
	}
}

func (this *ApiController) Cameras(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	cameras := make([]string, 0, len(this.config.Cameras))

	for key := range this.config.Cameras {
		cameras = append(cameras, key)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(cameras); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (this *ApiController) CamerasLive(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	cameraName := ps.ByName("camera")
	cam, ok := this.config.Cameras[cameraName]

	qp := r.URL.Query()
	active := qp.Get("active") == "true"
	preset := qp.Get("preset")
	if preset == "" {
		if active {
			preset = *this.config.Api.DefaultActiveCameraPreset
		} else {
			preset = *this.config.Api.DefaultNonActiveCameraPreset
		}
	}

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	p, err := this.getPreset(&cam, preset)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "video/mp4")
	err = runFfmpegLiveView(r.Context(), cameraName, p, w)
	if err != nil {
		log.Println(err)
	}
}
