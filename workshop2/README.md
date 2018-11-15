# Workshop2
If we have to take care vary and many container, it will be difficult task and can make you hothead.:rage:

So let's create `Dockerfile` for better life :sunglasses:

## Edit dockerfile
Edit your `Dockerfile` in the way that you want or just use mine for see how things work.

## Create docker image
  Now we have got template already, it time to create something cool.

  Before run make you are in the same directory `./workshop2` For create docker image just run.
  ```sh
  $ docker build -t workshop2 .
  ```
  View our docker image
  ```sh
  $ docker images
  ```

  Should see your docker image

## Create container from our image
Run your docker image to see the result.
```sh
$ docker run -p 6868:8080 workshop2
```

## We can create tag to versioning.
For versioning the image, we can tag image. Just put `:<some string>` after `<dockerimage>`. If we do not specify tag when build docker image,default tag will be used. Default tag is `latest`

Docker image format is  `<dockerimage>:<tag>`

Tag image when building.
```sh
$ docker build -t workshop2:v1 .
```

Tag image from other image.
```sh
$ docker tag workshop2 workshop2:v1
```