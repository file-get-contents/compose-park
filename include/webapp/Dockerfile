FROM debian:stable-slim AS test

ARG G_NAME
ARG G_GID
ARG G_UID


RUN apt-get update && apt-get install -y curl git
RUN groupadd -g ${G_GID} ${G_NAME} && \
    useradd -m -s /bin/bash -d /home/${G_NAME} -u ${G_UID} -g ${G_NAME} ${G_NAME}
RUN curl -OL https://go.dev/dl/go1.24.4.linux-amd64.tar.gz
RUN rm -rf /usr/local/go &&\
    tar -C /usr/local -xzf go1.24.4.linux-amd64.tar.gz &&\
    rm go1.24.4.linux-amd64.tar.gz 
RUN echo "export PATH=\$PATH:/usr/local/go/bin" >> /etc/bash.bashrc
ENV PATH=$PATH:/usr/local/go/bin
RUN . ~/.bashrc
RUN mkdir /app
COPY --chown=${G_NAME}:${G_NAME} --chmod=770 . /app

EXPOSE 3000
