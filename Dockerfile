FROM golang:1.14.0-alpine3.11 AS build

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -o /bin/tracer ./cmd/tracer/

FROM scratch
WORKDIR /workspace/
COPY --from=build /bin/tracer /bin/tracer
#COPY --from=build /src/configs/settings.toml /workspace/settings.toml

ENTRYPOINT ["/bin/tracer", "-t", "/workspace/settings.toml"]