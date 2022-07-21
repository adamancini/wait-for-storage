FROM gcr.io/distroless/static-debian11
COPY ./app /app
CMD ["./app"]
