FROM golang:1.12
WORKDIR /go/src/app
COPY . .

RUN \
apt-get update && \
apt-get install -y unzip

# Install protoc
RUN \
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip && \
unzip protoc-3.6.1-linux-x86_64.zip && \
mv include/* /usr/include/ && \
mv bin/* /usr/bin/ && \
chmod +x /usr/bin/protoc

RUN make init && make

FROM debian:9-slim
COPY --from=0 /go/bin/rdb /rdb
RUN apt-get update && apt-get install -y ca-certificates
CMD ["/rdb"]