# ingrid-coding-assignment
## Description

This repository contains the implementation of a small web service that takes the source and a list of destinations
and returns a list of routes between source and each destination.
Both source and destination are defined as a pair of latitude and longitude.
The returned list of routes is sorted by driving time and distance (if time is equal).

## Build and Run
In order to build and run the app normally you would need to do the following steps if you prefer using docker:

```bash

git clone git@github.com:fkryvyts-codete/ingrid-coding-assignment.git
cd ingrid-coding-assignment
make docker-build
make docker-run

```

After that you can open the following URL in browser or fetch it by CURL to see that the app is working

``` bash
curl 'http://localhost:8080/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219'
```

If it works you should see the following response

```json

{
  "source": "13.388860,52.517037",
  "routes": [
    {
      "Destination": "13.397634,52.529407",
      "Duration": 251.5,
      "Distance": 1884.8
    },
    {
      "Destination": "13.428555,52.523219",
      "Duration": 394.2,
      "Distance": 3841.7
    }
  ]
}

```

Alternatively, you can build and run the app without using docker, natively. For this execute the following commands:

```bash

git clone git@github.com:fkryvyts-codete/ingrid-coding-assignment.git
cd ingrid-coding-assignment
make build
make run

```

## Troubleshooting

Very often [router.project-osrm.org](`http://router.project-osrm.org`) is down and not responding
due to maintenance or other reasons. In this case the repository includes 2 scripts to run it locally
by using its official images from the Docker Hub

For this execute the following commands:

```bash

make osrm-dowload
make osrm-launch

```

First command downloads database for Berlin and the second one starts OSRM API on
http://0.0.0.0:5000

In order to make web service using this OSRM instance modify the following file

`ingrid-coding-assignment/app/config.yaml`

Copy and paste the following content there

```yaml
server:
  listen: "0.0.0.0:8080"

osrm:
  url: "http://0.0.0.0:5000"

```

After doing this, leave current bash window open (in order not to stop OSRM service)
and then in the new window execute

```bash

make build
make run

```

## Repository structure:

The `app/` directory contains source code of the application

The `app/cmd` directory contains CLI application boilerplate code that uses [spf13/cobra](https://github.com/spf13/cobra) library

The `app/pkg` directory contains Go packages with business logic

The `app/pkg/osrm` directory contains implementation of the client for calling OSRM API

The `app/pkg/routes` directory contains web service implementation

The `app/pkg/routes/transport/http` directory contains code responsible for handling http requests/responses. Uses [go-kit/kit](https://github.com/go-kit/kit)

The `app/pkg/routes/entities` directory contains data structures that are not related to any particular transport but are required by business logic

The `app/pkg/routes/service/service.go` file contains service business logic that is decoupled from any particular transport

The `app/pkg/routes/service/logging_middleware.go` file contains middleware that adds logging of requests/responses/errors to the service

The `app/config.yaml` file contains application configuration that can be changed without rebuilding the app

The `app/.golangci.yaml` file contains rules for golang-ci linter

The `app/Dockerfile` is required when building the app by using docker

The `scripts/` directory contains scripts for running OSRM API locally. Those scripts almost never needed to be run directly since `Makefile` has helper methods for calling them

The `data/` directory contains downloaded database files needed for OSRM. If local OSRM instance is not used, this folder will be empty

The `Makefile` file contains all helper commands needed when developing, building and running the app

The `.pre-commit-config.yaml` file contains configuration for pre-commit hooks that can be installed by using `make pre-commit`
