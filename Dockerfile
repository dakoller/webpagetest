FROM iron/base
WORKDIR /app
# copy binary into image
COPY main /app/
ENV URL "https://google.es"
ENTRYPOINT ["./main"]