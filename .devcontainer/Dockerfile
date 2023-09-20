FROM mcr.microsoft.com/devcontainers/go:1-1.21-bullseye

RUN apt-get update \
  && DEBIAN_FRONTEND=noninteractive apt-get install --no-install-recommends -y\
  peco\
  silversearcher-ag\
  && apt-get clean\
  && rm -rf /var/lib/apt/lists/*

# hadolint for docker
ARG HADOLINT_URL
RUN wget -q -O /usr/local/bin/hadolint ${HADOLINT_URL} --no-check-certificate \
  && chmod a+x /usr/local/bin/hadolint

# Config for Go
RUN go install -v golang.org/x/tools/gopls@latest\
  && go install -v github.com/go-delve/delve/cmd/dlv@latest\
  && go install -v github.com/ramya-rao-a/go-outline@latest\
  && chmod -R g+rwx /go

USER 1000
