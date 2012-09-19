primaupload:  primaupload.go
		go install

run:    primaupload
		./primaupload

all:    primaupload.go

fmt:
		gofmt -w primaupload.go

clean:
		rm -f static/uploads/*