module github.com/kraser/logger

go 1.17

require (
	github.com/kraser/errorshandler v0.0.0-20181012014344-40a6026a0d12
	github.com/sirupsen/logrus v1.8.1
)

require golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect

replace github.com/kraser/errorshandler => ../errorshandler
