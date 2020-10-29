FROM alpine

RUN apk update

RUN mkdir /app
ADD server /app
EXPOSE 5000
WORKDIR /app
ENTRYPOINT ["/app/server"]
