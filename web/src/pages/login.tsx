export function LoginPage() {
  return (
    <div class="vh-100 vw-100 d-flex align-items-center justify-content-center p-2">
      <div class="card overflow-hidden col-11 col-sm-6 col-md-4 col-lg-3 d-flex flex-column">
        <h3 class="d-flex align-items-center justify-content-center text-primary py-3">
          Teo 🐶
        </h3>
        <div class="border-bottom mx-3 mb-3"></div>
        <div class="input-group mb-3 px-3">
          <input
            class="form-control"
            type="text"
            id="username"
            placeholder="Username"
          />
        </div>
        <div class="input-group mb-3 px-3">
          <input
            class="form-control"
            type="password"
            id="username"
            placeholder="Password"
          />
        </div>
        <div class="input-group mb-3 px-3">
          <button class="btn btn-primary">Login</button>
        </div>
      </div>
    </div>
  );
}
