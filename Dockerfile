FROM nginx

COPY index.html /usr/share/nginx/html

COPY test.conf /etc/nginx

EXPOSE 80 

CMD ["nginx","-g","daemon off;"]

