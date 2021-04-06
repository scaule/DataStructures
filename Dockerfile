FROM golang:1.16 as builder


COPY . /tmp/go/.
# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
#ARG SKAFFOLD_GO_GCFLAG

RUN cd /tmp/go && go build .


FROM alpine:3
# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=single
CMD ["./app"]
COPY --from=builder /app .
