FROM alpine

WORKDIR /app

COPY ratelimiter ./

CMD ["/app/ratelimiter"]
