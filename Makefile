dev:
	air	

build:
	rm -rf build;
	go build -C ./cmd/swaystats -o ../../build/swaystats;
