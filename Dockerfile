FROM golang:1.23 as base
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 go build -v -o main .

FROM ubuntu:24.10

RUN apt-get update && apt-get install -y python3 wget
RUN apt-get install -y libegl1  libopengl0 libxcb-cursor0 libfreetype6 xz-utils
RUN wget -nv -O- https://download.calibre-ebook.com/linux-installer.sh | sh /dev/stdin
ENV PATH="$PATH:/opt/calibre/"
ENV TZ=Asia/Shanghai

VOLUME [ "/app/cache" ]
VOLUME [ "/app/uploads" ]

WORKDIR /app

COPY --from=base /app/main ./bookstack

COPY conf ./conf
COPY dictionary ./dictionary
COPY resources ./resources
COPY static ./static
COPY views ./views

CMD ["./bookstack"]