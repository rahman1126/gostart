# GoStart
Golang starter kit to build powerfull microservice with go clean architecture

### For Local

There is no need build version, so just run the main file.
For an example :

```sh
$ go run cmd/main.go
```

The configuration file will be found in the configuration folder

### For Development, Staging & Production

```sh
$ docker build . -t [SERVICE_NAME]:[SERVICE_VERSION]
$ docker run --rm -d -p [HOST_PORT]:[SERVICE_PORT] --name=[SERVICE_CONTAINER_NAME] [SERVICE_TAG]
```
