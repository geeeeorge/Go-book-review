FROM golang:1.20 as builder

RUN mkdir /artifact

WORKDIR /workspace

COPY ../.. .

RUN go build -o /artifact/app ./cmd/go-book-review

FROM gcr.io/distroless/base-debian10

COPY --from=builder /artifact/app /app

EXPOSE 8080
CMD ["/app"]