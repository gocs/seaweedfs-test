# seaweedfs-test

status: `fixed`

This a test where it manifests the error: `Upload err:multipart: NextPart: EOF`

Prerequisite: docker, docker-compose, golang 1.16

1. build and run the service

        make run

2. go to `localhost:8000`
3. upload an image
4. check the docker logs for the app service i.e. `seaweedfs-test-app-1`
