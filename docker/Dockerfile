
FROM golang:1.10-alpine3.8


LABEL AUTHOR=garfieldkwong<garfieldkwong@gmail.com>


WORKDIR /go/src/app

#run apk add proxychains-ng

#RUN echo "socks5 192.168.1.25 1080" >> /etc/proxychains.conf

run apk add curl wget

#run wget https://github.com/cyfdecyf/cow/releases/download/0.9.8/cow-linux64-0.9.8.gz

run  curl -L git.io/cow | sh

run apk add git


#run gunzip cow-linux64-0.9.8.gz
run echo "proxy = socks5://192.168.1.25:1080" >> ~/.cow/rc

run apk add bash

env http_proxy=http://127.0.0.1:7777
env https_proxy=http://127.0.0.1:7777

env SOCKS5_PROXY=192.168.1.25:1080

env COOKIES=Please_insert_cookies

#run nohup /go/src/app/cow

run echo fs.inotify.max_user_watches=1048576 | tee -a /etc/sysctl.conf && sysctl -p

copy . .
#run go get github.com/muyouming/gphotosuploader
run nohup bash -c "./cow &" &&  go get github.com/garfieldkwong/gphotosuploader

env IGNORE_PATTERN='^.*/\.@__thumb/.*$'

CMD ["/bin/bash","run.sh"]
