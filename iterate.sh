(go test -v && git add . && git commit -m "iterate `date +%Y%m%d%H%M%S`") || git checkout .
git push -q &
