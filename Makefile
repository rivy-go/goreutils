COMPILE=cd build && go build ../src/
compile:
	-mkdir build
	$(COMPILE)arch/arch.go
	$(COMPILE)cat/cat.go
	$(COMPILE)chmod/chmod.go
	$(COMPILE)date/date.go
	$(COMPILE)echo/echo.go
	$(COMPILE)false/false.go
	$(COMPILE)groups/groups.go
	$(COMPILE)ls/ls.go
	$(COMPILE)md5sum/md5sum.go
	$(COMPILE)mkdir/mkdir.go
	$(COMPILE)mv/mv.go
	$(COMPILE)pwd/pwd.go
	$(COMPILE)rm/rm.go
	$(COMPILE)sha1sum/sha1sum.go
	$(COMPILE)sha256sum/sha256sum.go
	$(COMPILE)sha512sum/sha512sum.go
	$(COMPILE)sleep/sleep.go
	$(COMPILE)touch/touch.go
	$(COMPILE)true/true.go
	$(COMPILE)whoami/whoami.go
	$(COMPILE)yes/yes.go
	$(COMPILE)kill/kill.go
	$(COMPILE)truncate/truncate.go
	$(COMPILE)cp/cp.go
	$(COMPILE)ln/ln.go
