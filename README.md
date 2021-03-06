# consul-registration
This project is for registering a service with a consul cluster. 


## Usage
This is mainly meant for use as a docker container sidekick to a primary service. See the [consul agent registration](https://www.consul.io/docs/agent/http/agent.html#agent_service_register) documentation for the API

```
docker run -d --name webapp organization/webapp 
docker run --link webapp:webapp micahhausler/consul-registration \
	-consul http://consul.example.com \
	-container webapp \
	-id webapp1 \
	-name webapp \
	-tag web \
	-tag webapp
```

For help output: just run with `-h`

```
$ docker run micahhausler/consul-registration -h
Usage of /bin/consul-registration:
  -consul="http://consul.service.consul": The address or IP for consul
  -container="": The container name to watch
  -http="": See https://www.consul.io/docs/agent/checks.html
  -id="": The service ID for consul
  -interval="45s": Interval for consul's HTTP check
  -name="": The service name for consul
  -note="": A note to pass along with service checks
  -once=false: Only register the service, then exit
  -script="": Script on consul server to execute
  -sleep=30: How long to wait between checking in with consul.
  -tag=[]: A tag to be applied to the service. Repeat option for multiple tags
  -ttl="45s": TTL for the service. Make this larget than -sleep
```


## Build

### Compile inside the Docker container

If you just want to build the app, but not run it in a docker container, run:

```bash
docker run --rm -v $(pwd):/go/src/github.com/micahhausler/consul-registration -w /go/src/github.com/micahhausler/consul-registration golang:1.4.2 go build -v

```
If you want to build for busybox and have a mini-container:

```bash
docker run --rm -v $(pwd):/go/src/github.com/micahhausler/consul-registration -w /go/src/github.com/micahhausler/consul-registration golang:1.4.2 go build -v -a -tags netgo -installsuffix netgo -ldflags '-w'
```

### Cross-compile inside the Docker container

If you need to compile for a platform other than linux/amd64 (such as windows/386), this can be easily accomplished with the provided cross tags:

```bash
docker run --rm -v $(pwd):/go/src/github.com/micahhausler/consul-registration -w /go/src/github.com/micahhausler/consul-registration -e GOOS=windows -e GOARCH=386 golang:1.4.2-cross go build -v
```

Alternatively, you can build for multiple platforms at once:

```bash
docker run --rm -it -v $(pwd):/go/src/github.com/micahhausler/consul-registration -w /go/src/github.com/micahhausler/consul-registration golang:1.4.2-cross bash
$ for GOOS in darwin linux; do
>   for GOARCH in 386 amd64; do
>     go build -v -o myapp-$GOOS-$GOARCH
>   done
> done
```

## License
See [License](/LICENSE)
