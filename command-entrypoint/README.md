# command
this plaything use command and entrypoint property.  
  
    cd command-entrypoint
    docker compose up --build

    ✔ base                    Built        0.0s 
    ✔ Container entrypointer  Recreated        0.1s 
    Attaching to entrypointer
    entrypointer  | compose entrypoint.  // the command is not executed and the entrypoint takes precedence.
    entrypointer exited with code 0

**the compose "command" is executed instead of the dockerfile "CMD".**  
**the compose "entrypoint" is executed instead of the dockerfile "ENTRYPOINT".**
