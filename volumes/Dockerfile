FROM httpd:latest

# createing dummy for compose volume
RUN cp -r ./conf /dummy

COPY --chown=www-data:www-data --chmod=770 ./conf /usr/local/apache2/conf
#CMD ["httpd-foreground"]
