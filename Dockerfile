FROM scratch
ARG APP_PORT=8081
COPY ./cmd/web-scraper/.bin/web-scraper /web-scraper
ENV GIN_MODE=release
EXPOSE 8081
ENTRYPOINT ["./web-scraper"]`