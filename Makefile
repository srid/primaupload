primaupload:  primaupload.go
		go install

run:    primaupload
		./primaupload

foreman:    primaupload
		foreman start

# create using: heroku create --buildpack git://github.com/kr/heroku-buildpack-go.git
push:
		git push -f heroku master

all:    primaupload.go

fmt:
		gofmt -w primaupload.go

clean:
		rm -f static/uploads/*