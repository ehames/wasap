FROM ubuntu:18.04

RUN apt update \
    && apt install -y libgtk-3-dev libwebkit2gtk-4.0-dev \
    && rm -rf /var/lib/apt/lists/*

ADD wasap /usr/bin/wasap

ENTRYPOINT [ "/usr/bin/wasap" ]