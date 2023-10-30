bin = webapp

test:
	go test -v ./...

watch:
	@go install github.com/cosmtrek/air@latest
	air \
	    --build.cmd "templ generate && go build -o ./dist/$(bin) ./cmd/$(bin)" --build.bin "./dist/$(bin)" \
	    --build.include_ext "go, templ" \
		--build.exclude_regex ".*_templ.go" \
	    --misc.clean_on_exit "true" & \
	npx tailwindcss -i ./assets/css/input.css -o ./assets/css/styles.css --watch
	# templ generate --watch

run: build
	./dist/$(bin)

build: tidy build-tailwind build-templ 
	go build -o ./dist/$(bin) ./cmd/$(bin)/

build-tailwind:
	npx tailwindcss -i ./assets/css/input.css -o ./assets/css/styles.css

build-templ:
	templ generate

audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -vet=off ./...

tidy:
	go mod tidy

migrate:
	go install -tags 'postgres,sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	migrate -path ./migrations -database "${DB_URL}" up

