# Workshop1
We are going to use container !!!
In this workshop we'll create Tomcat container for serve html content.

While doing this workshop try to think how easier compare to virtual machine.

## Create container


1. Start tomcat container
```sh
docker run -p 6969:8080 tomcat
```
2. Check container is running or not ?
```sh
docker ps -a
```
  Shoud see `CONTAINER ID`  and its `STATUS` : UP


## Create html file


  Now you have 2 options to deploy html to our tomcat server

  [A]. Create html file inside container.

  [B]. Create html in your machine and copy that file to container.

  ***To deploy html. We have to create directory inside `webapps` and then put html file inside there.

  Feel free to do any option or both if you can or dont do anything and buy some food for us thanks. :pizza::hamburger::beers:

### OptionA: Create html file inside container
  0. Get container id.
  ```sh
  docker ps -a
  ```
  1. Get inside container.
  ```sh
  docker exec -it <container id> /bin/bash
  ```
  2.   Create directory inside `webapps`
  ```sh
  mkdir -p webapps/workshop1
   ```

  3. Create simple html file
```sh
echo "<html><h1>Hello Tomcat</h1></html>" > webapps/workshop1/hello.html
```

### OptionB: Create html in your machine and copy that file to container.
0. Get container id.
  ```sh
  docker ps -a
  ```
1. Create html content with your favorite text editor.
2. Create directory inside `webapps`
```sh
$ docker exec <container ID> mkdir -p webapps/workshop1
```
3. Copy your file to container
```sh
$ docker cp <path/to/file> <contianer ID>:/usr/local/tomcat/webapps/workshop1
```

## Check your deployment
Just check our deployment work or not. Go to browser and enter `localhost:6969/workshop1/hello.html`
or change `hello.html` to your filename for option B.