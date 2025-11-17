# Procfile is used by the dev tool to run the application
# in development mode, the `app` process will be restarted
# automatically when changes are made to the source code.
app: go run cmd/app/main.go
css: go tool tailo --watch always -i internal/system/assets/tailwind.css -o internal/system/assets/application.css
