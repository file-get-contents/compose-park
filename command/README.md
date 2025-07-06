# [command](https://docs.docker.com/reference/compose-file/services/#command)
this plaything use "command" property.  
  
    cd command
    docker compose up --build

    ✔ base                             Built                      0.0s 
    ✔ Network      command_default     Created                    0.1s 
    ✔ Container    commander           Created                    0.1s 
    Attaching to commander
    commander  | compose commnad.   <<
    commander exited with code 0


> [!IMPORTANT]
> the compose "command" is executed instead of the dockerfile "CMD".


