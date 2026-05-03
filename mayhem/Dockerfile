FROM golang:1.25 AS builder

COPY . /tinet
WORKDIR /tinet

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN GOFLAGS=-mod=mod go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt-get update && apt-get install -y clang

RUN go-fuzz-build -libfuzzer -o mayhem/fuzzTinet.a github.com/tinynetwork/tinet/mayhem && \
    clang -fsanitize=fuzzer mayhem/fuzzTinet.a -o mayhem/fuzzTinet.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /tinet/mayhem/fuzzTinet.libfuzzer /
