module main

go 1.23.0

replace github.com/tsarquis88/file_mixer/pkg/midem => ../../pkg/midem

replace github.com/tsarquis88/file_mixer/pkg/cmdLineParser => ../../pkg/cmdLineParser

require (
	github.com/tsarquis88/file_mixer/pkg/cmdLineParser v0.0.0-00010101000000-000000000000
	github.com/tsarquis88/file_mixer/pkg/midem v0.0.0-00010101000000-000000000000
)
