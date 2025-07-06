FROM httpd:latest
RUN apt-get update && apt-get install -y openssl
RUN mkdir -p ssl && cd ssl
RUN openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout server.key -out server.crt -subj "/C=US/ST=Test/L=Test/O=Demo/CN=localhost"
COPY --chown=www-data:www-data --chmod=770 ./conf /usr/local/apache2/conf
CMD ["httpd-foreground"]
