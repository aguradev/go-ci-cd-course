FROM golang:alpine3.18 as dev

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod tidy
RUN go build -o main .

FROM alpine:3.18
WORKDIR /root/
COPY --from=dev /app/main .
EXPOSE 1234
CMD [ "./main" ]

