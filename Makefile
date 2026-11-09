gen:
	cd generators/apigen && go run .
	cd generators/libgen && go run .

run:
	npm run start

build:
	npm run build

publish: build
	gcloud storage rsync --recursive --checksums-only public gs://xn--melrse-egb.org