gen:
	cd generators/apigen && go run .
	cd generators/examplegen && go run .

run:
	npm run start

build:
	npm run build

new:
	hugo new content content/blog/collect.md