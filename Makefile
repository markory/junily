build-web-install:
	pnpm -r install
build-web-dist: build-web-install
	pnpm -r build
build-web-clean:
	rm -Rf internal/assets/dashboard
build-web-cp: build-web-clean
	cp -R web/dashboard/dist/ internal/assets/dashboard/

build-web: build-web-dist build-web-clean build-web-cp

build-app:
	go build -o cmd/junily cmd/junily/main.go  

build: build-web build-app
