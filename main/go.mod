module example.com/main

go 1.23.0

replace example.com/cmdLineParser => ../cmdLineParser

replace example.com/mixer => ../mixer

replace example.com/dataBytesFileManager => ../dataBytesFileManager

replace example.com/dataBytesManager => ../dataBytesManager

replace example.com/dataBytesDumper => ../dataBytesDumper

replace example.com/metadataManager => ../metadataManager

require (
	example.com/cmdLineParser v0.0.0-00010101000000-000000000000
	example.com/mixer v0.0.0-00010101000000-000000000000
)

require (
	example.com/dataBytesDumper v0.0.0-00010101000000-000000000000
	example.com/dataBytesFileManager v0.0.0-00010101000000-000000000000
	example.com/dataBytesManager v0.0.0-00010101000000-000000000000
	example.com/metadataManager v0.0.0-00010101000000-000000000000
)
