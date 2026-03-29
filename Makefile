gen:
	cd generators/apigen && go run .
	cd generators/examplegen && go run .

run:
	npm run start

build:
	npm run build

new:
	hugo new content content/blog/collect.md

publish: build
	gcloud storage rsync --recursive --checksums-only public gs://melrose.ernestmicklei.com