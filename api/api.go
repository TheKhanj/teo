package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ApiController struct {
	config *Config
}

func (this *ApiController) AddRoutes(router *httprouter.Router) {
	router.GET("/cameras/", this.auth(this.Cameras))
	router.GET("/cameras/:camera/live", this.auth(this.CamerasLive))
}

func (this *ApiController) getPreset(
	camera *ConfigCamera, preset string,
) (ConfigPreset, error) {
	ret := ConfigPreset{}

	if preset == PrimaryStream {
		ret.Stream = camera.Primary
	} else if preset == SecondaryStream {
		ret.Stream = camera.Secondary
	} else {
		p, presetExists := (*this.config.Api.Presets)[preset]
		if !presetExists {
			return ret, fmt.Errorf("preset %s does not exist", preset)
		}
		ret = p
		if p.Stream == PrimaryStream {
			ret.Stream = camera.Primary
		} else if p.Stream == SecondaryStream {
			ret.Stream = camera.Secondary
		} else {
			return ret, fmt.Errorf("stream %s is not valid, valid options are \"primary\" or \"secondary\"", p.Stream)
		}
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

	ffmpeg := Ffmpeg{
		Context:    r.Context(),
		CameraName:    cameraName,
		Preset: p,
		W:      w,
		Audio:  active,
	}
	err = ffmpeg.Live()
	if err != nil {
		log.Println(err)
	}
}
