IMAGE_NAME = s2i-example

.PHONY: build
build:
	podman build -t $(IMAGE_NAME) .

.PHONY: test
test:
	podman build -t $(IMAGE_NAME)-candidate .
	IMAGE_NAME=$(IMAGE_NAME)-candidate test/run
