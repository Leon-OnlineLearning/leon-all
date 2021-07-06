FROM canyan/janus-gateway:latest

RUN apt-get -y update
RUN apt-get install -y nginx

COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 41234
CMD [ "nginx" ]