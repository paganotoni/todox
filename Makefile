# Setup task:
# - Downloads Tailwind CSS standalone CLI
# - Downloads Air for live reloading
setup:
	@echo "ℹ️  Downloading Air Binary. Please wait..."
	@curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
	@echo "✅ Air binary downloaded."

	@echo "ℹ️  Downloading Tailwind CSS binary. Please wait..."
	@curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64

	@echo "✅ TailwindCSS binary downloaded."
	@chmod +x tailwindcss-macos-arm64
	@mv tailwindcss-macos-arm64 bin/tailwindcss


	@go run ./cmd/tools/ migrate

# Run the application in development mode watching for changes in the
# html and go files.
run:
	@./bin/air --build.cmd "go build -o bin/app ./cmd/cach/main.go" --build.bin "./bin/app" -log.main_only true -build.exclude_regex "\\.html" &\
	./bin/tailwindcss -i application.css --content "./*.html,./**/*.html" -o public/application.css -w &\
	wait

# Build the css and then the app, a few notes:
# - This assumes the build command is being run on a container using linux and x64 architecture.
# - The order matters here because the app uses the css file.
build:
	@echo "ℹ️  Downloading Tailwind CSS binary. Please wait..."
	@curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
	@chmod +x tailwindcss-linux-x64
	@mv tailwindcss-linux-x64 tailwindcss
	@echo "✅ Tailwind CSS binary downloaded."

	./tailwindcss -i application.css --content "./*.html,./**/*.html" -o public/application.css
	go build -ldflags '-s -w -extldflags "-static"' -tags osusergo,netgo,sqlite_omit_load_extension -o bin/app ./cmd/app/main.go
	go build -ldflags '-s -w -extldflags "-static"' -tags osusergo,netgo,sqlite_omit_load_extension -o bin/tools ./cmd/tools/main.go