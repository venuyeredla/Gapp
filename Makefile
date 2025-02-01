Hello :
	echo "Hello Gapp"

Run:
	go run Application.go

Build:
	go bulid

BuildIMG:
	docker build . -t goapp:latest

DocRun:
	docker run --name=GoApp -d -p 2024:2024 goapp:latest