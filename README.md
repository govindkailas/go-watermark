[![Go](https://github.com/govindkailas/go-watermark/actions/workflows/go.yml/badge.svg)](https://github.com/govindkailas/go-watermark/actions/workflows/go.yml)
[![Docker](https://github.com/govindkailas/go-watermark/actions/workflows/docker-publish.yml/badge.svg?branch=main)](https://github.com/govindkailas/go-watermark/actions/workflows/docker-publish.yml)

# Example

- add the image on assets folder
- change the name on main.go line **20**
- change the watermark on line **29**
- run main.go

# the result should be something like this:

![Example](image-with-overlay.jpg?raw=true)

## Running watermarker as a microservice
Use the branch `add-gin-router` to see the example with Gin Router microservice. This would expose an endpoint to add watermark to an image. 
Post the image url and watermark text as a form data to the endpoint.

```
git clone -b add-gin-router https://github.com/govindkailas/go-watermark.git
```

Once the app is up, try running the following curl command to see the health
`curl http://127.0.0.1:8080/`

```
curl --request POST \
  --url http://127.0.0.1:8080/watermark \
  --header 'content-type: multipart/form-data' \
  --form url=https://www.pitara.com/media/mango.jpg \
  --form text=MyWaterMark@2024
```

A [Hoppscotch](https://hoppscotch.io) screenshot is attached here, 
![post form](watermark-form-data.jpg?raw=true)

## Running watermarker as a microservice in Kubernetes
To deploy the watermarker as a microservice in Kubernetes, follow the steps below:
`kubectl apply -f watermark-k8s-deploy`
