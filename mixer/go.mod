module example.com/mixer

go 1.23.0

replace example.com/dataBytesManager => ../dataBytesManager

replace example.com/dataBytesManagerMock => ../dataBytesManagerMock

require (
	example.com/dataBytesManager v0.0.0-00010101000000-000000000000
	example.com/dataBytesManagerMock v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
