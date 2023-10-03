bin = webapp

audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...

test:
	go test -v -race -buildvcs ./...

watch:
	go run github.com/cosmtrek/air@v1.43.0 \
	    --build.cmd "make build-without-tailwind" --build.bin "./dist/$(bin)" --build.delay "100" \
	    --build.exclude_regex "(.*?)_templ.go" \
	    --build.include_ext "go, tpl, templ, html, css, js, sql, svg" \
	    --misc.clean_on_exit "true" & \
	npx tailwindcss -i ./assets/css/input.css -o ./assets/css/styles.css --watch

run: build
	./dist/$(bin)

build: build-templ build-tailwind
	go build -o ./dist/$(bin) ./cmd/$(bin)/

build-without-tailwind: build-templ
	go build -o ./dist/$(bin) ./cmd/$(bin)/

build-tailwind:
	npx tailwindcss -i ./assets/css/input.css -o ./assets/css/styles.css

build-templ:
	templ generate
