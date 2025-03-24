import { A } from "@solidjs/router";
import { Title } from "@solidjs/meta";

function ErrorPage(title: string, message: string) {
  return (
    <>
      <Title>{title}</Title>
      <div
        class="text-center vh-100 d-flex flex-column justify-content-center align-items-center"
        style="background-color: #f8f9fa"
      >
        <h1 class="display-1 text-danger">404</h1>
        <p class="lead">{message}</p>
        <A href="/" class="btn btn-primary">
          Go Home
        </A>
      </div>
    </>
  );
}

export function NotFoundPage() {
  return ErrorPage(
    "Not Found",
    "Oops! The page you're looking for doesn't exist.",
  );
}
