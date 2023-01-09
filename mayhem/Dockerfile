FROM golang:1.18 as builder

COPY . /tinet
WORKDIR /tinet/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzTinet.a && \
    clang -fsanitize=fuzzer fuzzTinet.a -o fuzzTinet.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /tinet/mayhem/fuzzTinet.libfuzzer /