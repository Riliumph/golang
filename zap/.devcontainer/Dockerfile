#FROM debian:bookworm-slim
FROM golang:bullseye

ARG HADOLINT_URL
#ENV DEBIAN_FRONTEND=noninteractive
RUN sed -i 's@archive.ubuntu.com@ftp.jaist.ac.jp/pub/Linux@g' /etc/apt/sources.list \
  && apt-get update \
  && apt-get install --no-install-recommends -y\
  wget\
  && apt-get clean\
  && rm -rf /var/lib/apt/lists/*
# hadolint for docker
ARG HADOLINT_URL
RUN wget --progress=dot:giga -O /usr/local/bin/hadolint ${HADOLINT_URL} --no-check-certificate \
  && chmod +x /usr/local/bin/hadolint

# Config for Go
RUN go install -v golang.org/x/tools/gopls@latest\
  && go install -v github.com/go-delve/delve/cmd/dlv@latest\
  && go install -v github.com/ramya-rao-a/go-outline@latest\
  && chmod -R o+rwx /go

# New user
SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN groupadd -g 1000 docker\
  && useradd -l -u 1000 -g docker -G sudo -m -s /bin/bash docker\
  && echo 'docker:docker' | chpasswd\
  && echo 'ALL ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers
USER docker
RUN printf "PS1='\u@\h(\$(hostname -i)):\w \n\\$ '" >> ~/.bashrc
WORKDIR /workspace

ENV LOG_DIR="./"
ENV LOG_LEVEL="info"

ENTRYPOINT [ "/bin/bash" ]
