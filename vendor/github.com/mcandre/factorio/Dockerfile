FROM alpine:3.23 AS build
RUN apk add -U go
COPY . /src
WORKDIR /src
RUN go install ./...

FROM scratch
COPY --from=build /root/go/bin/factorio /factorio
ENTRYPOINT ["/factorio"]
