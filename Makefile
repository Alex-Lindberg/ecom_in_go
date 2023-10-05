build:
	go build -o goecom cmd/main.go

run: build
	./myapp

clean:
	rm myapp
