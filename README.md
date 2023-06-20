# grpc_nmap_wrapper

The repository contains the [gRPC](https://grpc.io/docs/languages/go/quickstart/) service, which is a wrapper over [Nmap](https://nmap.org/) using the [vulners](https://github.com/vulnersCom/nmap-vulners) script for vulnerability detection.

Before you can run the application, you must install it:  
-> `nmap` [installation instructions](https://nmap.org/download.html);  
-> `golangci-lint` [installation instruction](https://golangci-lint.run/usage/install/).

Creating a binary file and running the application:
```golang
make buid
```
To run the client: 
```golang
make client
```
To run the tests:
```golang
make test
```
To run lint:
```golang
make lint
```
