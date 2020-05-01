#! /bin/bash -eu

echo fs.inotify.max_user_watches=1048576 | tee -a /etc/sysctl.conf && sysctl -p
echo "listen = http://127.0.0.1:7777" > ~/.cow/rc
echo "proxy = socks5://${SOCKS5_PROXY}" >> ~/.cow/rc
echo "${COOKIES}" >> /go/src/app/auth.json
ignore_patterns=(${IGNORE_PATTERN})
ignore_arg=''
for val in "${ignore_patterns}"
do
    ignore_arg="$ignore_arg --ignore $val"
done
echo "Run gphotosUploader $ignore_arg"
/bin/bash -c "nohup ./cow & "
/bin/bash -c "sleep 4; /go/bin/gphotosuploader --watch /photo --maxConcurrent 4 $ignore_arg" 
