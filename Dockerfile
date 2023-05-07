From couchbase/centos7-systemd

RUN mkdir -p /home/general/conf

COPY general-service /home/general

COPY conf /home/general/conf

WORKDIR /home/general

EXPOSE 8080

CMD ["./general-service"]