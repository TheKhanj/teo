#!/usr/bin/bash

get_cam() {
	local camera_name="$1"

	cat <<-EOF
		<div class="col-12 col-md-6 p-2">
		  <div class="card overflow-hidden">
		    <div class="d-flex p-2 bg-body-tertiary text-primary">
		      <div class="flex-grow-1 d-flex flex-column justify-content-center px-2">${camera_name}</div>
		      <div class="controls">
		        <button class="btn btn-sm">⚙️</button>
		      </div>
		    </div>
		    <div class="d-flex" style="height: 300px">
		      <video class="flex-grow-1 mw-100 mh-100" style="object-fit: cover;" controls autoplay muted>
		        <source src="http://192.168.40.200:8081/$camera_name/live" type="video/mp4">
		        Your browser does not support the video tag.
		      </video>
		    </div>
		  </div>
		</div>
	EOF
}

get_view_port() {
	cat <<-EOF
		<div class="vh-100 d-flex flex-column overflow-auto bg-body-tertiary">
		  <h4 class="bg-body text-primary border-bottom p-3">Cameras</h4>
		  <div class="d-flex flex-wrap p-2 overflow-auto">
		    $(get_cam "cam2")
		    $(get_cam "cam3")
		    $(get_cam "cam4")
		  </div>
		</div>
	EOF
}

get_page() {
	cat <<-EOF
		<!doctype html>
		<html lang="en">
		  $(get_header)
		  <body>
		    $(get_dashboard "live" <(get_view_port))
		    <style>
		      * {
		        transition: all 0.2s ease;
		      }
		      body {
		        height: 100vh;
		        overflow: hidden;
		      }
		    </style>
		  </body>
		</html>
	EOF
}

main() {
	source ../common/header.bash
	source ../common/dashboard.bash

	echo 'Content-Type: text/html'
	echo ''
	get_page
}

main
