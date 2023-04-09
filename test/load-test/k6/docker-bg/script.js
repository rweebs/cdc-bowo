import http from "k6/http";
import { check, sleep } from "k6";

export default function () {
  // Make an HTTP request to the specified endpoint
  let response = http.get("http://app.default.svc.cluster.local");

  console.log(response.body)
  
  // Check that the response status code is 200
  check(response, {
    "status is 200": (r) => r.status === 200,
  });

  // Sleep for 1 second
}