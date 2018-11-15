# Workshop2
If we have to take care vary and many container, it will be difficult task and can make you hothead.:rage: 

So let's create `dockerfile` for better life :sunglasses:

## Edit dockerfile
Edit your `dockerfile` in the way that you want or just use mine for see how things work.

## Create docker image
  Now we have got template already, it time to create something cool.

  Before run make you are in the same directory `./workshop2` For create docker image just run. 
  ```sh
  $ docker build -t <dockerimageName> .
  ```
  View our docker image
  ```sh
  $ docker images
  ```

  Should see your docker image

## Let's what we build.
Run your docker image to see the result.
```sh
$ docker run -p <hostPath>:<containerPath|8080> <dockerimageName>
```

## We can create tag to versioning.
For versioning the image, we can tag image. Just put `:<some string>` after `<dockerimage>`. Default tag is `latest`

Docker image format is  `<dockerimage>:<tag>`

Tag image when building.
```sh
$ docker build -t <dockerimage>:<tag> .
``` 

Tag image from other image.
```sh
$ docker tag <imageID | imageName:tag original> <imageName:tag NewOne>
```