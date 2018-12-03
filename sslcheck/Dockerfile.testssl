FROM alpine

RUN apk add --update bash curl tar bind-tools coreutils tini

RUN curl -fsSL -o /usr/local/bin/testssl.sh https://testssl.sh/testssl.sh && chmod +x /usr/local/bin/testssl.sh
RUN curl -fsSL https://testssl.sh/etc+doc.tar.gz | tar xzvf -
RUN curl -fsSL https://testssl.sh/openssl-1.0.2k-chacha.pm.ipv6.Linux+FreeBSD.201705.tar.gz | tar xzvf - && chmod +x /bin/openssl.Linux.x86_64.static && mv /bin/openssl.Linux.x86_64.static /bin/openssl

ENTRYPOINT ["/sbin/tini", "/usr/local/bin/testssl.sh"]
