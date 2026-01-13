FROM alpine:3.23 AS build
RUN apk add -U go
COPY . /src
WORKDIR /src
RUN go install ./...

FROM scratch
COPY --from=build /root/go/bin/funk /funk
COPY --from=build /root/go/bin/stank /stank
COPY --from=build /root/go/bin/stink /stink
ENTRYPOINT ["/stank"]
