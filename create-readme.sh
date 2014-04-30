cat <<EOT
go-crypto
=========

Learning cryptography using Go. I am trying to learn cryptography -- the math,
the security and the programming aspects -- using Go language. Coincidentally
I also want to learn the Go language. Nice dovetailing huh!

These pieces of code are probably insecure and wrong. So use it only as a
starting point. I will make corrections based on my learning and based on
feedback.

How to run: Just go into a numbered directory and do \`go run src.go\`.

EOT

for i in `ls -d [0-9]*`
do
	echo "- [$i]($i)"
	echo -n " - "
	cat $i/README.md
done
