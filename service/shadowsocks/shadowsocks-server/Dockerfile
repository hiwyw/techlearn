FROM alpine:3.16.2 as build
ARG ss_version
WORKDIR /home
RUN apk add --update -t deps curl \
    && curl -L https://github.com/shadowsocks/shadowsocks-rust/releases/download/${ss_version}/shadowsocks-${ss_version}.x86_64-unknown-linux-musl.tar.xz -o /home/ss.tar.xz \
    && tar -xvf ss.tar.xz

FROM alpine:3.16.2
COPY --from=build /home/ssserver /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/ssserver"]
CMD ["-c", "/etc/shadowsocks-server.json"]

