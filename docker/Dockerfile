
FROM golang:1.10-alpine3.8


LABEL AUTHOR=garfieldkwong<garfieldkwong@gmail.com>


WORKDIR /go/src/app

run apk add curl wget gcc libc-dev

run apk add git

run apk add bash

env COOKIES=Please_insert_cookies

run echo fs.inotify.max_user_watches=1048576 | tee -a /etc/sysctl.conf && sysctl -p

copy . .
#run go get github.com/muyouming/gphotosuploader
run go get -u github.com/jinzhu/gorm && go get github.com/garfieldkwong/gphotosuploader

env IGNORE_PATTERN='^.*/\.@__thumb/.*$'

CMD ["/bin/bash","run.sh"]
