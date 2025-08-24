hello:
	echo "Hello"

build-web:
	pnpm -r build & cp -R web/dashboard/dist/ deploy/dashboard/

build-app:
	go build -o cmd/junily cmd/junily/main.go  

build: build-web build-app
	echo "Hello!"

clean:
	rm -Rf deploy/dashboard