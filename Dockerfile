FROM ubuntu:24.10

RUN apt-get update && apt-get install -y python3 wget
RUN apt-get install -y libegl1  libopengl0 libxcb-cursor0 libfreetype6 xz-utils
RUN wget -nv -O- https://download.calibre-ebook.com/linux-installer.sh | sh /dev/stdin
ENV PATH="$PATH:/opt/calibre/"
ENV TZ=Asia/Shanghai

VOLUME [ "/app/cache" ]
VOLUME [ "/app/uploads" ]

WORKDIR /app

COPY bookstack ./
COPY conf ./conf
COPY dictionary ./dictionary
COPY resources ./resources
COPY static ./static
COPY views ./views

CMD ["./bookstack"]