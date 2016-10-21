FROM scratch
EXPOSE 8080

COPY server /
COPY assets/ /assets/
COPY templates /templates/
COPY images.csv /
ENTRYPOINT ["/server"]
