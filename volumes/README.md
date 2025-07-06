# [volumes](https://docs.docker.com/reference/compose-file/volumes/)
this plaything use volumes property.  
  
    cd volumes
    docker compose up --build
    docker exec -it volumes bash
    
    root@5f32f8ec104b:/usr/local/apache2# ls -l ./conf/extra
    -rw-r--r-- 1 root     root      2881 Jul  1 02:14 httpd-autoindex.conf
    -rw-r--r-- 1 root     root      1950 Jul  1 02:14 httpd-dav.conf
    -rw-r--r-- 1 root     root      2942 Jul  1 02:14 httpd-default.conf
    -rw-r--r-- 1 root     root      1119 Jul  1 02:14 httpd-info.conf
    -rw-r--r-- 1 root     root      5078 Jul  1 02:14 httpd-languages.conf
    -rw-r--r-- 1 root     root      1395 Jul  1 02:14 httpd-manual.conf
    -rw-r--r-- 1 root     root      4444 Jul  1 02:14 httpd-mpm.conf
    -rw-r--r-- 1 root     root      2222 Jul  1 02:14 httpd-multilang-errordoc.conf
    -rw-r--r-- 1 root     root     13304 Jul  1 02:14 httpd-ssl.conf
    -rw-r--r-- 1 root     root       694 Jul  1 02:14 httpd-userdir.conf
    -rwxrwx--- 1 www-data www-data  1060 Jul  6 03:40 httpd-vhosts.conf     << modified
    -rw-r--r-- 1 root     root      3161 Jul  1 02:14 proxy-html.conf

    root@5f32f8ec104b:/usr/local/apache2# ls -l /dummy/extra
    -rwxrwxrwx 1 root root 1060 Jul  6 03:40 httpd-vhosts.conf


> [!IMPORTANT]
> the dockerfile copy command allows you to reflect differences while maintaining the directory structure, but this is not possible with compose volumes. because the purpose of a compose volume is to maintain space.
