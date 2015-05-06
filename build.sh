cmd_prefix=please-

for pkg in cmd/*; do
    pkg=$(basename $pkg)

    go build -o $cmd_prefix$pkg cmd/$pkg/main.go
done
