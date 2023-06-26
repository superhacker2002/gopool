run00:
	go build -o readDB cmd/ex00/main.go
	./readDB -f stolen_database.json
	rm -rf readDB

run01:
	go build -o compareDB cmd/ex01/main.go
	./compareDB --old ../original_database.xml --new ../stolen_database.json
	rm -rf compareDB