# go + cloud spanner example

### Run on docker

```sh
docker build -t go-spanner .
docker run --rm -v ${GOOGLE_APPLICATION_CREDENTIALS}:/key.json -e GOOGLE_APPLICATION_CREDENTIALS=/key.json -e SPANNER_DSN=projects/xxx/instances/xxx/databases/xxx go-spanner
```
