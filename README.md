# Redis
This is an implementation of the server-client solution for storing KV data, lightweight analog of Redis done on pure Go (only standard libraries). Keys and values are utf8-encoded strings.

The project can be executed with the help of [Makefile](./Makefile) and Docker containers. Makefile has `build` and `run` targets.

1. `build` target is necessary to compile the binaries for [server](./main/server.go) and [client](./main/client.go) and create a Docker __image__ (according to [Dockerfile](./Dockerfile))
2. `run` target is used for starting Docker containers with default configurations.

## Makefile usage
|Command|Description|
|:-------|:-----------|
|`build`|`make build` to create the image with installed tools and build binaries inside containers with default configurations which can be changed in [Makefile](./Makefile) by adding `-p`, `--port` and `-m`, `--mode` flags with required value (default for port is "9090", for mode is "disk").|
|`run`|`make run` to start the containers: server in a detached mode and client with allocating a pseudo-TTY connected to the container’s stdin and creating an interactive bash shell in the container. Default configurations can be changed with `-h` or `--host` and `-p` or `--port` flags (default value for host is "127.0.0.1", port "9090"). `$(CURDIR)/main` directory on the host machine is mounted into the container.|
|`check`|`make check` to run subsequently "go vet", "goimports", "golint" on the project.|
 ## Client commands
 |Command|Description|
 |:---|:---|
 |`set`|Updates one key at a time with the given value|.
 |`get`|Returns tuple of the value and the key state. The state either present or absent.|
 |`del`|Removes one key at a time and returns the state of the resource. The state either ignored or absent. A key isignored if it does not exist.|
 |`keys`|Returns all keys matching pattern. Pattern could include \* symbol which matches zero or more characters.|
 |`stop`| To stop the client.|
