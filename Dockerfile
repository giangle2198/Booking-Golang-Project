# syntax=docker/dockerfile:1

FROM golang:1.17.6-alpine3.15
WORKDIR /app
ARG MONGO__URL=mongodb://172.27.42.11:27017
ARG BIM_CLIENT__HOST=172.27.42.11
ARG ENABLED_CORS=true
ENV MONGO__URL $MONGO__URL
ENV BIM_CLIENT__HOST $BIM_CLIENT__HOST
ENV ENABLED_CORS $ENABLED_CORS
COPY . /app
RUN go mod download
RUN go mod vendor
RUN cd cmd/server && go build -o bcs .
RUN mkdir -p statik && cp -p cmd/server/bcs . && cp -p cmd/server/statik/statik.go statik
CMD ["./bcs"]