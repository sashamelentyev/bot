FROM gcr.io/distroless/static

ADD bot /usr/local/bin/bot

ENTRYPOINT ["/usr/local/bin/bot"]
