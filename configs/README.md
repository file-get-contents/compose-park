# [configs](https://docs.docker.com/reference/compose-file/services/#configs)
this plaything use configs property.  
  
    cd configs
    docker compose up --build
    docker exec -it configs bash
    
    root@7c990870cd72:/usr/local/apache2# ls -l ./conf
    drwxr-xr-x 2 root root  4096 Jul  1 02:14 extra
    -rwxrwxrwx 1 root root 21401 Jul  6 04:23 httpd.conf  << 
    -rw-r--r-- 1 root root 13064 Jul  1 02:14 magic
    -rw-r--r-- 1 root root 60970 Jul  1 02:14 mime.types
    drwxr-xr-x 3 root root  4096 Jul  1 02:14 original


> [!WARNING]
> Does not support uid, gid, and mode properties.
> [[BUG] Configs uid, gid and mode is will not take effect with configs of type file ](https://github.com/docker/compose/issues/12270)
