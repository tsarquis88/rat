module example.com/main

go 1.23.0

replace example.com/cmdLineParser => ../cmdLineParser

replace example.com/mixer => ../mixer

replace example.com/fileManager => ../fileManager

require (
	example.com/cmdLineParser v0.0.0-00010101000000-000000000000
	example.com/mixer v0.0.0-00010101000000-000000000000
)

require example.com/fileManager v0.0.0-00010101000000-000000000000 // indirect
