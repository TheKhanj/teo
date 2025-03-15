get_sidebar() {
	local active="$1"

	cat <<-EOF
		<div class="col-2 vh-100 nav nav-pills p-2 d-flex flex-column bg-body-secondary" style="min-width: 200px">
		  <h3 class="text-center text-primary py-3">Teo 🐶</h3>
			<a class="nav-item nav-link m-1 $([ "$active" = "live" ] && echo -n active)" href="#">Live</a>
		  <a class="nav-item nav-link m-1" href="#">Recordings</a>
		  <a class="nav-item nav-link m-1" href="#">Settings</a>
		</div>
	EOF
}

get_dashboard() {
	local active="$1"
	local view_port="$2"

	cat <<-EOF
		<div class="d-flex d-print-table">
		$(get_sidebar "$active")
		$(cat "$view_port")
		</div>
	EOF
}
