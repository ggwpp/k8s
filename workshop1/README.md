# Workshop1
Create html content and deploy to Tomcat.

To deploy html. We have to create directory inside `webapps` and then put html file inside there.

## Create container


1. Start tomcat container
```sh
$ docker run -p <hostPort>:<containerPort|8080> tomcat
```
2. Check container is running or not ?
```sh
$ docker ps -a
``` 
  Shoud see `CONTAINER ID`  and its `STATUS` : UP


## Create html file


  Now you have 2 options to deploy html to our tomcat server
  
  [1]. Create html file inside content.

  [2]. Create html in your machine and copy that file to container.

  Feel free to do any option or both if you can or dont do anything and buy some food for us thanks. :pizza::hamburger::beers:

### Option1: Create html file inside container
  1. Get inside container.
  ```sh
  $ docker exec -it <container id> /bin/bash
  ``` 


  Your shell must have change to container na. :ok_hand:


2.   Create dir inside `webapps` 
     ```sh 
     $ mkdir -p webapps/<dirName>
      ```

1. Create simple html file 
```sh
$ echo "<html><h1>Your content</h1></html>" > webapps/<dir name>/<fileName>.html
```

### Option2: Create html in your machine and copy that file to container.

1. Create html content with your favorite editor.
2. Create directory inside `webapps`
```sh
$ docker exec <container ID> mkdir -p webapps/<dir name> 
```
3. Copy your file to container
```sh
$ docker cp <path/to/file> <contianer ID>:/usr/local/tomcat/webapps/<dirName>
```

## Check your deployment 
Just check our deployment work or not. Go to browser and enter `localhost:<hostPort>/<dirName>/<fileName>`


