#!/bin/sh

echo -n "Password: "
read line

echo "$line" | ruby AQRPC.rb > out.ruby
echo "$line" | AQRPC         > out.go

diff out.ruby out.go

rm out.ruby out.go
