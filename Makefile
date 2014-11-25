default:
	go build -o anagram-generator
	./anagram-generator "hello world"

build:
	go build -o anagram-generator examples/anagram-generator/anagram-generator.go
	go build -o anagram-service examples/anagram-service/anagram-service.go
	go build -o parallel-generator examples/parallel-generator/parallel-generator.go

generator:
	go build -o anagram-generator examples/anagram-generator/anagram-generator.go

service:
	go build -o anagram-service examples/anagram-service/anagram-service.go

parallel:
	go build -o parallel-generator examples/parallel-generator/parallel-generator.go

