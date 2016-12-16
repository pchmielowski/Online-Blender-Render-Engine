FROM pchmielowski/go_and_blender

MAINTAINER Piotr Chmielowski

COPY index.html /index.html
COPY result.html /result.html
COPY run.bash /run.bash
RUN mkdir -p /go/src/app
COPY main.go /go/src/app/main.go

WORKDIR /go/src/app
RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]
