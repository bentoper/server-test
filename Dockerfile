FROM alpine

RUN apk -U upgrade && \
    apk --update --no-cache add \
      ca-certificates && \
    rm -rf /tmp/src && \
    rm -rf /var/cache/apk/* 

WORKDIR /Users/bentoper/go/src/server
COPY server .

RUN apk --no-cache add curl
EXPOSE 8080 8081

ENTRYPOINT [ "./server" ]
CMD ["./server"]
