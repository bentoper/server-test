FROM alpine

RUN apk -U upgrade && \
      apk --update --no-cache add \
      ca-certificates && \
      apk --update add xtables-addons && \
	    apk --update add iptables && \
      apk --no-cache add curl && \
      apk --update add sudo && \  
      rm -rf /tmp/src && \
      rm -rf /var/cache/apk/* 

WORKDIR /Users/bentoper/go/src/server
COPY server .

EXPOSE 8080 8081

ENTRYPOINT [ "./server" ]
CMD ["./server"]
