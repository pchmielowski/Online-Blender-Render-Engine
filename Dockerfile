FROM golang:1.6-onbuild
COPY index.html /index.html
COPY result.html /result.html
COPY run.bash /run.bash
RUN apt-get update
RUN apt-get install -y blender
