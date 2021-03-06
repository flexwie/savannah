# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.192.0/containers/ubuntu/.devcontainer/base.Dockerfile

ARG VARIANT=focal
FROM mcr.microsoft.com/vscode/devcontainers/base:${VARIANT}

# [Optional] Uncomment this section to install additional OS packages.
RUN apt update && export DEBIAN_FRONTEND=noninteractive \
  && apt -y install --no-install-recommends build-essential git \
  #protobuf-compiler\
  apt-transport-https \
  ca-certificates \
  curl \
  gnupg \
  lsb-release


RUN /bin/bash -c "mkdir -p /usr/local/include/google/api && wget -P /usr/local/include/google/api https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/{annotations,http}.proto"
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v3.19.1/protoc-3.19.1-linux-x86_64.zip && \
  mkdir -p /usr/local/protoc && \
  unzip protoc-3.19.1-linux-x86_64.zip -d /usr/local bin/protoc && \
  chmod +x /usr/local/bin/protoc && \
  unzip protoc-3.19.1-linux-x86_64.zip -d /usr/local include/* && \
  chmod -R 775 /usr/local/include/google && \
  rm protoc-3.19.1-linux-x86_64.zip


# Install Golang
RUN wget -c https://go.dev/dl/go1.18beta1.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local

USER vscode
ENV PATH="${PATH}:/usr/local/go/bin:/home/vscode/go/bin"

# Install vscode tools
RUN go install golang.org/x/tools/gopls@latest && \
  go install github.com/go-delve/delve/cmd/dlv@latest && \
  go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest && \
  go install github.com/ramya-rao-a/go-outline@latest && \
  go install honnef.co/go/tools/cmd/staticcheck@latest && \
  go install github.com/gobuffalo/packr/v2/packr2@latest

# Install Golang proto tools
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
  go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest && \
  go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest


USER root


# Install Node.js
RUN curl -fsSL https://deb.nodesource.com/setup_16.x | sudo -E bash - && \
  apt install -y nodejs && \
  npm install --global yarn



# Install buf
RUN BIN="/usr/local/bin" && \
  VERSION="0.54.1" && \
  BINARY_NAME="buf" && \
  curl -sSL \
  "https://github.com/bufbuild/buf/releases/download/v${VERSION}/${BINARY_NAME}-$(uname -s)-$(uname -m)" \
  -o "${BIN}/${BINARY_NAME}" && \
  chmod +x "${BIN}/${BINARY_NAME}"


ENTRYPOINT [ "/usr/local/share/docker-init.sh" ]