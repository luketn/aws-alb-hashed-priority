FROM golang as build
WORKDIR /app
ADD main.go .
RUN go build main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/main ./
ENTRYPOINT ["/app/main"]