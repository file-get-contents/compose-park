# command
this plaything use entrypoint property.  
  
    cd entrypoint
    docker compose up --build

    ✔ base                        Built        0.0s 
    ✔ Network entrypoint_default  Created      0.1s 
    ✔ Container entrypointer      Created      0.1s 
    Attaching to entrypointer
    entrypointer  | compose entrypoint. // the compose "entrypoint" is executed instead of the dockerfile "ENTRYPOINT".
    entrypointer exited with code 0

