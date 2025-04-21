templ:
	templ generate --watch --open-browser=false -v

run-air:
	air -c .air.toml

tailwind-clean:
	tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --clean

tailwind-watch:
	tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --watch

dev:
	@echo "Starting..."
	@make -j3 tailwind-watch templ run-air
