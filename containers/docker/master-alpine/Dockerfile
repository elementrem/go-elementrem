FROM alpine:3.5

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/1.5 https://github.com/elementrem/go-elementrem && \
  (cd go-elementrem && make gele) && \
  cp go-elementrem/build/bin/gele /gele && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /go-elementrem && rm -rf /var/cache/apk/*

EXPOSE 7075
EXPOSE 30707

ENTRYPOINT ["/gele"]
