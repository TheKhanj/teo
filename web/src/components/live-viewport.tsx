function Camera(props: { name: string }) {
  const url = `http://192.168.40.200:8081/${props.name}/live`;

  return (
    <div class="col-12 col-md-6 p-2">
      <div class="card overflow-hidden">
        <div class="d-flex p-2 bg-body-tertiary text-primary">
          <div class="flex-grow-1 d-flex flex-column justify-content-center px-2">
            {props.name}
          </div>
          <div class="controls">
            <button class="btn btn-sm">⚙️</button>
          </div>
        </div>
        <div class="d-flex" style="height: 300px">
          <video
            class="flex-grow-1 mw-100 mh-100"
            style="object-fit: cover;"
            controls
            autoplay
            muted
          >
            <source src={url} type="video/mp4" />
            Your browser does not support the video tag.
          </video>
        </div>
      </div>
    </div>
  );
}

export function LiveViewport() {
  return (
    <div class="vh-100 d-flex flex-column overflow-auto bg-body-tertiary">
      <h4 class="bg-body text-primary border-bottom p-3">Cameras</h4>
      <div class="d-flex flex-wrap p-2 overflow-auto">
        <Camera name="cam2" />
        <Camera name="cam3" />
        <Camera name="cam4" />
      </div>
    </div>
  );
}
