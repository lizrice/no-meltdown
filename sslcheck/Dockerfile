FROM ubuntu:trusty

RUN apt-get -y update && apt-get -y install nginx openssl
RUN mkdir -p /etc/nginx/conf.d && cd /etc/nginx/conf.d && openssl req -subj '/CN=localhost' -x509 -newkey rsa:4096 -nodes -keyout key.pem -out cert.pem -days 365
ADD server.conf /etc/nginx/conf.d/server.conf
ENTRYPOINT ["nginx", "-g", "daemon off;"]
