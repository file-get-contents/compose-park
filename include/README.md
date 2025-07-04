# include
this plaything use reverse-proxy and webapp(Go).  
include each item in the "include" property of compose.yml.  
this means that they exist on the same network and can recognize each other by their service-name.
```
cd include
docker compose up --build

access to ->  http://localhost:8080
```