all:
	go build -o build/arch src/arch/arch.go
	go build -o build/cat src/cat/cat.go
	go build -o build/chmod src/chmod/chmod.go
	go build -o build/date src/date/date.go
	go build -o build/echo src/echo/echo.go
	go build -o build/exit src/exit/exit.go
	go build -o build/false src/false/false.go
	go build -o build/groups src/groups/groups.go
	go build -o build/ls src/ls/ls.go
	go build -o build/md5sum src/md5sum/md5sum.go
	go build -o build/mkdir src/mkdir/mkdir.go
	go build -o build/mv src/mv/mv.go
	go build -o build/pwd src/pwd/pwd.go
	go build -o build/rm src/rm/rm.go
	go build -o build/sha1sum src/sha1sum/sha1sum.go
	go build -o build/sha256sum src/sha256sum/sha256sum.go
	go build -o build/sha512sum src/sha512sum/sha512sum.go
	go build -o build/sleep src/sleep/sleep.go
	go build -o build/touch src/touch/touch.go
	go build -o build/true src/true/true.go
	go build -o build/whoami src/whoami/whoami.go
	go build -o build/yes src/yes/yes.go
	go build -o build/kill src/kill/kill.go
test:
	go test src/ls/ls_test.go
	go test src/arch/arch_test.go
	go test src/echo/echo_test.go
	go test src/pwd/pwd_test.go
	go test src/whoami/whoami_test.go
	go test src/cat/cat_test.go
	go test src/false/false_test.go
	go test src/true/true_test.go
	go test src/md5sum/md5sum_test.go
	go test src/sha1sum/sha1sum_test.go
	go test src/sha256sum/sha256sum_test.go
	go test src/sha512sum/sha512sum_test.go
