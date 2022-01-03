# CTFd

```sh
cd ~
git clone https://github.com/CTFd/CTFd.git
cd CTFd
sudo apt install python3-pip -y
```

## Nginx

1. Update Nginx conf, by removing the outer `http` scope:

```diff
diff --git a/conf/nginx/http.conf b/conf/nginx/http.conf
index ff40613..2e00f41 100644
--- a/conf/nginx/http.conf
+++ b/conf/nginx/http.conf
@@ -1,16 +1,7 @@
-worker_processes 4;
-
-events {
-
-  worker_connections 1024;
-}
-
-http {
-
   # Configuration containing list of application servers
   upstream app_servers {

-    server ctfd:8000;
+    server localhost:8000;
   }

   server {

     listen 80;
+    server_name 2021.santahack.xyz;

     client_max_body_size 4G;

@@ -46,4 +38,3 @@ http {
       proxy_set_header X-Forwarded-Host $server_name;
     }
   }
-}
diff --git a/docker-compose.yml b/docker-compose.yml
index 4160ba5..e253cc0 100644
--- a/docker-compose.yml
```

2. Disable Nginx in the docker-compose:

```diff
diff --git a/docker-compose.yml b/docker-compose.yml
index 4160ba5..e253cc0 100644
--- a/docker-compose.yml
+++ b/docker-compose.yml
@@ -29,6 +29,8 @@ services:
   nginx:
     image: nginx:1.17
     restart: always
+    deploy:
+      replicas: 0
     volumes:
       - ./conf/nginx/http.conf:/etc/nginx/nginx.conf
     ports:
```

Don't forget to run `up` to refresh the compose:

```sh
cd ~/CTFd
sudo docker-compose up -d
```

3. Install and configure Nginx:

```sh
sudo apt install nginx

sudo ln -vsf ~/CTFd/conf/nginx/http.conf /etc/nginx/sites-available/ctfd.conf

sudo unlink /etc/nginx/sites-enabled/default
sudo ln -vsf /etc/nginx/sites-available/ctfd.conf /etc/nginx/sites-enabled/ctfd.conf

sudo nginx -s reload
```

## Certificate (Let's Encrypt)

Requires native Nginx, from the section above.

```sh
# Update snap
sudo snap install core
sudo snap refresh core

# Install certbot
sudo snap install --classic certbot

# Prepare certbot
sudo ln -sv /snap/bin/certbot /usr/bin/certbot

# Run certbot. It's interactive
sudo certbot --nginx
```

Certbot will automatically update the `~/CTFd/conf/nginx/http.conf` file as it
was linked in the Nginx section above, adding something like:

```diff
diff --git a/conf/nginx/http.conf b/conf/nginx/http.conf
index ff40613..c6580a4 100644
--- a/conf/nginx/http.conf
+++ b/conf/nginx/http.conf
@@ -45,5 +35,26 @@ http {
       proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
       proxy_set_header X-Forwarded-Host $server_name;
     }
-  }
+
+    listen 443 ssl; # managed by Certbot
+    ssl_certificate /etc/letsencrypt/live/2021.santahack.xyz/fullchain.pem; # managed by Certbot
+    ssl_certificate_key /etc/letsencrypt/live/2021.santahack.xyz/privkey.pem; # managed by Certbot
+    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
+    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot
+
 }
+
+
+  server {
+    if ($host = 2021.santahack.xyz) {
+        return 301 https://$host$request_uri;
+    } # managed by Certbot
+
+
+
+    listen 80;
+    server_name 2021.santahack.xyz;
+    return 404; # managed by Certbot
+
+
+}
\ No newline at end of file
```
