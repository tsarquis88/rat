module main

go 1.23.0

replace github.com/tsarquis88/file_mixer/pkg/midem => ../../pkg/midem

require (
	github.com/stretchr/testify v1.9.0
	github.com/tsarquis88/file_mixer/pkg/midem v0.0.0-00010101000000-000000000000
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
