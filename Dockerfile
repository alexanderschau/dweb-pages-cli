FROM golang:1.16.3-alpine AS build

WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -o dweb-pages
RUN CGO_ENABLED=0 go build -o waitForIpfs action/waitForIpfs/waitForIpfs.go

FROM ipfs/go-ipfs:v0.9.0

USER root

COPY --from=build /app/dweb-pages /bin/dweb-pages
COPY --from=build /app/waitForIpfs /bin/waitForIpfs

COPY action/entrypoint.sh /entrypoint.sh

WORKDIR /project

ENV TAGS=
ENV ENDPOINT=
ENV ACCESS_TOKEN=
ENV NAME=

ENTRYPOINT [ "/bin/sh" ]
CMD [ "/entrypoint.sh" ]