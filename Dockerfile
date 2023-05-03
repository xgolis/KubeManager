FROM golang:latest as build
WORKDIR /app
COPY . .
RUN go mod download
RUN cd cmd/KubeManager && \
    CGO_ENABLE=0 GOOS=linux go build -o ../../KubeManager && \
    cd ../..

FROM xgolis/deployimage:latest as deploy
RUN chmod 400 /root/.ssh/id_rsa
RUN ssh-keyscan 35.240.30.14 >> /root/.ssh/known_hosts
RUN mkdir ~/.kube
RUN scp -o StrictHostKeyChecking=no -i /root/.ssh/id_rsa xgolis@35.240.30.14:/home/xgolis/.kube/config ~/.kube
RUN kubectl config set-cluster kubernetes --server=https://35.240.30.14:6443
RUN kubectl config set-cluster kubernetes --insecure-skip-tls-verify

FROM redhat/ubi8:latest
COPY --from=build /app/KubeManager .
COPY --from=deploy /root/.ssh/ /root/.ssh
RUN mkdir /root/.kube
COPY --from=deploy /root/.kube /root/.kube
COPY --from=deploy /usr/local/bin/helm /usr/local/bin
COPY --from=deploy /usr/local/bin/kubectl /usr/local/bin

EXPOSE 8085
ENTRYPOINT ["./KubeManager"]

          