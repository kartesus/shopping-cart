FROM golang:1.19.4-bullseye

RUN apt update && apt upgrade -y && \
    apt install -y git \
    make openssh-client

WORKDIR /src

EXPOSE 8001

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air