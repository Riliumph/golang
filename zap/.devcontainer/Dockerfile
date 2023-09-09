FROM golang:bullseye

RUN sed -i 's@archive.ubuntu.com@ftp.jaist.ac.jp/pub/Linux@g' /etc/apt/sources.list \
  && apt-get update \
  && DEBIAN_FRONTEND=noninteractive apt-get install --no-install-recommends -y\
  sudo\
  wget\
  && apt-get clean\
  && rm -rf /var/lib/apt/lists/*

# hadolint for docker
ARG HADOLINT_URL
RUN wget -q -O /usr/local/bin/hadolint ${HADOLINT_URL} --no-check-certificate \
  && chmod +x /usr/local/bin/hadolint

# Config for Go
RUN go install -v golang.org/x/tools/gopls@latest\
  && go install -v github.com/go-delve/delve/cmd/dlv@latest\
  && go install -v github.com/ramya-rao-a/go-outline@latest\
  && chmod -R o+rwx /go

# New user
SHELL ["/bin/bash", "-o", "pipefail", "-c"]
ARG USERNAME=devcontainer
RUN groupadd -g 1000 ${USERNAME}\
  && useradd -l -u 1000 -g ${USERNAME} -G sudo -m -s /bin/bash ${USERNAME}\
  && echo "${USERNAME}:${USERNAME}" | chpasswd\
  && echo 'ALL ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers
USER 1000
RUN printf "PS1='\[\$(tput setaf 4)\]\u\[\$(tput sgr0)\]@\[\$(tput setaf 2)\]\h(\$(hostname -i))\[\$(tput sgr0)\]:\[\$(tput setaf 3)\]\w\\n\[\$(tput sgr0)\]\\$ '" >> ~/.bashrc
WORKDIR /workspace

ENV LOG_DIR="./"
ENV LOG_LEVEL="info"

ENTRYPOINT [ "/bin/bash" ]
