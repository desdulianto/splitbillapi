FROM alpine:3.6 as builder

ADD https://releases.hashicorp.com/envconsul/0.6.2/envconsul_0.6.2_linux_amd64.tgz /tmp/envconsul.tgz
RUN tar -xvzf /tmp/envconsul.tgz -C /usr/local/bin/
COPY deploy/_output/web/web /usr/local/bin/web

FROM alpine:3.6

COPY --from=builder /usr/local/bin/web /usr/local/bin/envconsul /usr/local/bin/
EXPOSE 8080
ENTRYPOINT ["web"]
