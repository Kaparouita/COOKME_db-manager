
# Define variables
IMAGE_NAME := kaparouita/db-manager
VERSION := latest

# Build the Docker image
build:
	docker build -t $(IMAGE_NAME):$(VERSION) .
    #docker build -t $(IMAGE_NAME):$(VERSION) .
# Push the Docker image
# push:
#     docker push $(IMAGE_NAME):$(VERSION)

# # Run the Docker container
# # run:
# #     docker run -p 4000:80 $(IMAGE_NAME):$(VERSION)

# # Example of a multi-command workflow
# deploy: build push

