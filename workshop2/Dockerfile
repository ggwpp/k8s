FROM tomcat
#You can edit `maintainer` value
LABEL maintainer="Patipat <p.ponyanan@gmail.com>"


WORKDIR /usr/local/tomcat

#You can edit dir 
RUN mkdir -p webapps/grean
COPY index.html webapps/grean

EXPOSE 8080