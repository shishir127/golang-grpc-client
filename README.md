Spike for implementing a golang grpc client with server streaming.

* Refer to the [quick start guide](https://grpc.io/docs/quickstart/go.html) for directions on installing grpc dependencies.
* Edit `spike/spike.proto`.
* Run `make grpc` to generate code.
* Run `make build` to generate the binary.
* Binary needs SSL cert path, port and access token as environment variables to run e.g. `PORT=443 CERT=public.cert TOKEN=test ./grpc-client`
