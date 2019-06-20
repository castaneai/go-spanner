# go + cloud spanner example

## Usage

### Run on docker

```sh
docker build -t go-spanner .
docker run --rm -v ${GOOGLE_APPLICATION_CREDENTIALS}:/key.json -e GOOGLE_APPLICATION_CREDENTIALS=/key.json -e SPANNER_DSN=projects/xxx/instances/xxx/databases/xxx go-spanner
```

## Tips

### Enable gRPC logging

```sh
-e GRPC_GO_LOG_SEVERITY_LEVEL=info -e GRPC_GO_LOG_VERBOSITY_LEVEL=99
```
