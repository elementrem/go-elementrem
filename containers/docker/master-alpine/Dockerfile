FROM alpine:3.3

RUN \
  apk add --update go git make gcc musl-dev gmp-dev gmp && \
  git clone https://github.com/elementrem/go-elementrem     && \
  (cd go-elementrem && make gele)                         && \
  cp go-elementrem/build/bin/gele /gele                   && \
  apk del go git make gcc musl-dev gmp-dev              && \
  rm -rf /go-elementrem && rm -rf /var/cache/apk/*

EXPOSE 7075
EXPOSE 30707

ENTRYPOINT ["/gele"]
