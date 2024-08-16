FROM alpine:3.19

ARG ALPINE_VERSION=3.19

LABEL Maintainer="Ahmad Mohammadi <ahmadmohammadi940@gmail.com>" \
      Description="Lightweight container with Nginx 1.24 & Go based on Alpine Linux."

RUN echo https://mirrors.pardisco.co/alpine/v$ALPINE_VERSION/main > /etc/apk/repositories
RUN echo https://mirrors.pardisco.co/alpine/v$ALPINE_VERSION/community >> /etc/apk/repositories

# Install packages and remove default server definition


RUN apk add --no-cache --virtual openssl \
    curl \
    nano \
    git \
    vim \
    htop \
    bash \
    gcc \
    musl-dev

# Configure nginx
# COPY .docker/dev/config/nginx.conf /etc/nginx/nginx.conf

RUN set -x \
	&& adduser -u 1000 -D -S -G www-data www-data

COPY --from=golang:1.23 /usr/local/go/ /usr/local/go/
# ENV PATH="/usr/local/go/bin:${PATH}"
# RUN echo "export PATH=$PATH" >> /etc/profile
RUN ln -s /usr/local/go/bin/go /usr/bin/go
RUN chown -R www-data.www-data /usr/bin/go

# download go tar 
# RUN wget -O go.tgz https://dl.google.com/go/go1.10.3.src.tar.gz 
# RUN tar -C /usr/local -xzf go.tgz 
# RUN export PATH="/usr/local/go/bin:$PATH" \
#     export GOPATH=/opt/go/  \
#     export PATH=$PATH:$GOPATH/bin 
# Switch to use a non-root user from here on
USER www-data
# RUN go version
# Add application

# Expose the port nginx is reachable on
EXPOSE 8080

# ENTRYPOINT ["./.docker/dev/docker-entrypoint.sh"]
