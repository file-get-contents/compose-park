FROM httpd:latest

RUN touch test.txt

CMD ["/usr/local/bash", "-c", "echo \"dockerfile CMD\"."]
ENTRYPOINT ["echo", "dockerfile ENTRYPOINT."]
