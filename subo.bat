:: Comandos Linux
git add . 
git commit -m "Next commit"
git push
set GOOS=linux
set GOARCH=amd64
go build -tags lambda.norpc -o bootstrap main.go
del main.zip
tar.exe -a -cf main.zip bootstrap