FROM ubuntu:22.04
COPY . /root
RUN apt update -y
RUN apt-get install tzdata
ENV TZ Europe/Moscow
RUN apt-get install postgresql golang-go ca-certificates -y
RUN echo "listen_addresses = '*'" >> /etc/postgresql/14/main/postgresql.conf
WORKDIR /root/cmd
RUN awk -v cmd='openssl x509 -noout -subject' '/BEGIN/{close(cmd)};{print | cmd}' < /etc/ssl/certs/ca-certificates.crt
RUN go build -i ./application.go
EXPOSE 8080
CMD ["/root/cmd/application"]
