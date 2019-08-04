# https://github.com/Miroka96/docker-makefile

NAME = miroka96/stromberg-telegram-bot
TAG = 1.0

LOCALPORT = 8080
CONTAINERPORT = 8080

# if you want a special image name, edit this
IMAGE = $(NAME):$(TAG)

# if you publish no ports, delete the right part
PORTPUBLISHING = -p $(LOCALPORT):$(CONTAINERPORT)

.PHONY: build test test-shell build-test deploy build-deploy undeploy redeploy build-redeploy clean-volume clean-container clean install-dependencies configure

build:
	docker build -t $(IMAGE) .

build-nocache:
	docker build -t $(IMAGE) --no-cache .

test:
	docker run $(PORTPUBLISHING) --rm $(IMAGE)

test-shell:
	docker run $(PORTPUBLISHING) -it --rm $(IMAGE) /bin/bash

build-test: build test

deploy:
	docker run --detach --restart always --name=$(NAME) $(PORTPUBLISHING) $(IMAGE)

build-deploy: build deploy

undeploy:
	-docker stop $(NAME)
	docker rm $(NAME)

redeploy: undeploy deploy

build-redeploy: build redeploy

clean-volume:
	-docker volume rm $(VOLUME)

clean-container:
	-docker rm $(NAME)

clean: clean-volume clean-container