FROM alpine:3.11 as build
ARG ss_version
WORKDIR /home
RUN apk add --update -t deps curl \
    && curl -L https://github.com/shadowsocks/shadowsocks-rust/releases/download/${ss_version}/shadowsocks-v1.8.10-stable.x86_64-unknown-linux-musl.tar.xz -o /home/ss.tar.xz \
    && tar -xvf ss.tar.xz

FROM alpine:3.11
COPY --from=build /home/ssserver /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/ssserver"]
CMD ["-c", "/etc/shadowsocks-server.json"]
