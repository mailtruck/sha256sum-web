FROM alpine

ADD sha256sum-web ./sha256sum-web
ADD static ./static

ENTRYPOINT ["./sha256sum-web"]