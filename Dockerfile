FROM golang:1.10.0-stretch
LABEL maintainer="jjm3333@gmail.com"

COPY . ./src/TacIt/
RUN cd src/TacIt/ && \
    go install

CMD ["./bin/TacIt"]
