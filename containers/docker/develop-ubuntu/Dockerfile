FROM ubuntu:xenial

RUN \
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang git make gcc libc-dev ca-certificates && \
  git clone --depth 1 https://github.com/elementrem/go-elementrem && \
  (cd go-elementrem && make gele) && \
  cp go-elementrem/build/bin/gele /gele && \
  apt-get remove -y golang git make gcc libc-dev && apt autoremove -y && apt-get clean && \
  rm -rf /go-elementrem

EXPOSE 7075
EXPOSE 30707

ENTRYPOINT ["/gele"]
