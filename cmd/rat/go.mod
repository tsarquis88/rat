module main

go 1.23.0

replace github.com/tsarquis88/rat/pkg/rat => ../../pkg/rat

replace github.com/tsarquis88/rat/pkg/cmdLineParser => ../../pkg/cmdLineParser

require (
	github.com/tsarquis88/rat/pkg/cmdLineParser v0.0.0-00010101000000-000000000000
	github.com/tsarquis88/rat/pkg/rat v0.0.0-00010101000000-000000000000
)
