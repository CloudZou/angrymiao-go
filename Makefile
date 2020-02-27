BAZEL           = bazel
KCTL            = kubectl
REDIS_POD       = $(shell $(KCTL) get pods -l app=redis -o jsonpath='{.items[0].metadata.name}')
REDIS_CLI_EXEC  = $(KCTL) exec -it $(REDIS_POD) -- redis-cli

clean:
	$(BAZEL) clean --expunge

.PHONY: run
run: gazelle
	$(BAZEL) run //cmd -- -conf ${config

.PHONY: build
build: gazelle
	$(BAZEL) build //...

dep-ensure:
	go mod download
	go mod vendor

#gazelle-repos:
#	$(BAZEL) run //:gazelle -- update-repos -from_file=go.mod

gazelle:
	$(BAZEL) run //:gazelle

go-setup: dep-ensure gazelle

docker-registry:
	docker run -d -p 5000:5000 --restart=always --name registry registry:2

minikube-start:
	minikube start --memory 5120 --cpus 4 --insecure-registry localhost:5000

minikube-setup:
	minikube addons enable ingress

# docker-local-push: go-setup
# 	$(BAZEL) run //:colossus-web -- --norun
# 	$(BAZEL) run //:colossus-auth -- --norun
# 	$(BAZEL) run //:colossus-data -- --norun
# 	$(BAZEL) run //:colossus-userinfo -- --norun

k8s-app-init-deploy:
	$(KCTL) apply -f k8s/service.yaml
	$(KCTL) apply -f k8s/ingress.yaml
	$(KCTL) apply -f k8s/deployment-public.yaml

k8s-app-deploy:
	$(KCTL) apply -f k8s/deployment-public.yaml

redis-set-password:
	$(REDIS_CLI_EXEC) SET password tonydanza

redis-get-password:
	$(REDIS_CLI_EXEC) GET password

deploy: docker-local-push k8s-colossus-deploy

# restart-colossus:
# 	$(KCTL) delete -f k8s/colossus.yaml
# 	$(KCTL) apply -f k8s/colossus.yaml

teardown:
	$(KCTL) delete svc,deployment,ing --all
	$(KCTL) delete po/busybox

restart: teardown deploy

busybox-run:
	$(KCTL) run curl --image=radial/busyboxplus:curl -i --tty
