FROM tomcat
#You can edit `maintainer` value
LABEL maintainer="Patipat <p.ponyanan@gmail.com>"


WORKDIR /usr/local/tomcat

#You can change directory name or change to index.html to your file.
RUN mkdir -p webapps/workshop2
COPY index.html webapps/workshop2

EXPOSE 8080