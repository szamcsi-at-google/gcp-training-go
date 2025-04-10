# gcp-training-go

A toy application for very specific demo purposes.

## Local Prerequisites:

- Go installed.
- Docker installed.
- Google Cloud SDK (gcloud) installed and configured (logged into your Google Cloud account).

## Remote Prerequisites:

- A Google Cloud Project with the Cloud Run API and Artifact Registry API enabled.
- An Artifact Registry Docker repository created.

## Development:

### Go Run

If you want to run the local go application you can run:
```
go run main.go
```
Once it's running you can curl against localhost:8080 and see the response

### Go Tests
To run the unit tests:
```
go test
```

### Go Linting
We like to lint with golangci-lint, install it and then run:
```
golangci-lint run
```

### Docker

To build the docker image locally run:
```
docker build -t gcp-training-go .
```

To test the local image run:
```
docker run --rm -p 9090:8080 --name local-gcp-training-go gcp-training-go
```
Once it's running you can curl against localhost:9090 and see the response

### GCP

Uploading to Artifact Registry (replace the capitals vars):
```
# authenticate
gcloud auth configure-docker YOUR_REGION-docker.pkg.dev

# build your image
docker buildx build --platform linux/amd64 -t YOUR_REGION-docker.pkg.dev/GCP_PROJECT_NAME/YOUR_REPO_NAME/gcp-training-go .

# push image
docker push YOUR_REGION-docker.pkg.dev/GCP_PROJECT_NAME/YOUR_REPO_NAME/gcp-training-go
```

Deploy to Cloud Run:
```
gcloud run deploy gcp-training-go \
  --image YOUR_REGION-docker.pkg.dev/YOUR_PROJECT_ID/YOUR_REPO_NAME/gcp-training-go:latest \
  --region YOUR_REGION \
  --allow-unauthenticated
```

Once deployed, gcloud will output the Service URL. You can access your running application at this URL.
