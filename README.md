## Usage:

1. create base container (during the build Go and Blender will be installed)

`docker build --tag go_and_blender base/`

2. create final container (during the build necessary files will be copied and go binary builded)

`docker build --tag render .`

3. run container

`docker run -p 8080:8080 render`

4. go to `localhost:8080`
5. upload \*.blend file
6. wait for the result
