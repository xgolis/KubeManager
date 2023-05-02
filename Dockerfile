FROM golang:latest as build
WORKDIR /app
COPY . .
RUN go mod download
RUN cd cmd/KubeManager && \
    CGO_ENABLE=0 GOOS=linux go build -o ../../KubeManager && \
    cd ../..

FROM redhat/ubi8:latest
COPY --from=build /app/KubeManager .
EXPOSE 8085
ENTRYPOINT [ "./KubeManager" ]