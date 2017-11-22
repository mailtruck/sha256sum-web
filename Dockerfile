FROM golang

ADD hello ./hello

ENTRYPOINT ["./hello"]