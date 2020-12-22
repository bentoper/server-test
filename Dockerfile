FROM alpine

RUN apk -U upgrade && \
    apk --update --no-cache add \
      ca-certificates && \
    rm -rf /tmp/src && \
    rm -rf /var/cache/apk/*

WORKDIR /Users/bentoper/go/src/server
COPY server .

EXPOSE 8080 8081

ENV DEST 8081

RUN apk --no-cache add curl

ENTRYPOINT [ "./server" ]
CMD ["./server"]
