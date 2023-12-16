module github.com/cdvelop/gotools

go 1.20

require (
	github.com/cdvelop/input v0.0.73
	golang.org/x/text v0.14.0
)

require (
	github.com/cdvelop/model v0.0.102 // indirect
	github.com/cdvelop/strings v0.0.9 // indirect
	github.com/cdvelop/timetools v0.0.32 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/timetools => ../timetools

replace github.com/cdvelop/input => ../input
