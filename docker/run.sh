#! /bin/bash -eu

echo fs.inotify.max_user_watches=1048576 | tee -a /etc/sysctl.conf && sysctl -p
echo "${COOKIES}" >> /go/src/app/auth.json
ignore_arg=''
for val in ${IGNORE_PATTERN}
do
    ignore_arg="$ignore_arg --ignore $val"
done
echo "Run gphotosUploader $ignore_arg"
/bin/bash -c "/go/bin/gphotosuploader --watch /photo --maxConcurrent 4 $ignore_arg" 
