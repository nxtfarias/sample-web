module example.com/hello

go 1.16

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	golang.org/x/text v0.3.6 // indirect
	rsc.io/quote v1.5.2
)

replace example.com/greetings => ../greetings
