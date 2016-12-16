Following instruction explain how to create a container with a web application with **Blender** as a backend to render 3D models.

### Technology stack

* Docker
* Blender
* Go
* HTML + [Tacit CSS framework](http://yegor256.github.io/tacit/)

## Usage:

* create final container (during the build necessary files will be copied and go binary builded)

`docker build --tag render .`

* run container

`docker run -p 8080:8080 render`

* go to `localhost:8080`
* upload \*.blend file
* wait for the result

## Google Cloud

Alternatively, container can be run on Container engine such as Google Cloud.
* [How to push image to Google Container Registry](https://cloud.google.com/container-registry/docs/pushing)
* [How to run container in Google Container Engine](https://cloud.google.com/container-engine/docs/quickstart)
