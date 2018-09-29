
FROM centos
MAINTAINER shadow.zhoou "shadowzhoutianlei@gmail.com"
ADD ./binary /var/www/

ADD ./runtime.config.json /var/www/

ADD ./src/static /var/www/src/static
WORKDIR /var/www
EXPOSE 80
ENTRYPOINT ["./binary"]