module github.com/cdvelop/gotools

go 1.20

require (
	github.com/cdvelop/input v0.0.58
	golang.org/x/text v0.14.0
)

require (
	github.com/cdvelop/model v0.0.75 // indirect
	github.com/cdvelop/strings v0.0.7 // indirect
	github.com/cdvelop/timetools v0.0.24 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/timetools => ../timetools

replace github.com/cdvelop/input => ../input
