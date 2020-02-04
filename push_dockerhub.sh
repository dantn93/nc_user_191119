# Build the docker image
docker build -t tranngocdan/tranngocdan-nc-user .

# Tag the image

# Login to docker with your docker Id
# $ docker login
# Login with your Docker ID to push and pull images from Docker Hub. If you don\'t have a Docker ID, head over to https://hub.docker.com to create one.
# Username (callicoder): callicoder
# Password:
# Login Succeeded

# Push the image to docker hub
docker push tranngocdan/tranngocdan-nc-user:latest