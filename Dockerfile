FROM golang

ADD hello ./hello
ADD static ./static

ENTRYPOINT ["./hello"]