FROM golang:1.23 as base
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 go build -v -o main .

FROM ubuntu:24.10

RUN apt-get update \
    && apt-get install -y python3 wget \
    && apt-get install -y fonts-wqy-zenhei fonts-wqy-microhei \
    && apt-get install -y libegl1  libopengl0 libxcb-cursor0 libfreetype6 xz-utils chromium-browser
#RUN wget -nv -O- https://download.calibre-ebook.com/linux-installer.sh | sh /dev/stdin
COPY calibre.txz /opt/calibre/calibre.txz
RUN cd /opt/calibre && xz -d calibre.txz

ENV PATH="$PATH:/opt/calibre/"
ENV TZ=Asia/Shanghai
ENV LANG en_US.utf8

VOLUME [ "/app/cache" ]
VOLUME [ "/app/uploads" ]

WORKDIR /app

COPY --from=base /app/main ./bookstack

COPY conf ./conf
COPY dictionary ./dictionary
COPY static ./static
COPY views ./views

RUN ./bookstack install

CMD ["./bookstack"]