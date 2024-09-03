module example.com/main

go 1.23.0

replace example.com/cmdLineParser => ../../pkg/cmdLineParser

replace example.com/mixer => ../../pkg/mixer

replace example.com/demixer => ../../pkg/demixer

replace example.com/dataBytesFileManager => ../../pkg/dataBytesFileManager

replace example.com/dataBytesManager => ../../pkg/dataBytesManager

replace example.com/dataBytesDumper => ../../pkg/dataBytesDumper

replace example.com/metadataManager => ../../pkg/metadataManager

replace example.com/dataBytesManagerMock => ../../pkg/dataBytesManagerMock

require (
	example.com/cmdLineParser v0.0.0-00010101000000-000000000000
	example.com/mixer v0.0.0-00010101000000-000000000000
)

require (
	example.com/dataBytesDumper v0.0.0-00010101000000-000000000000
	example.com/dataBytesFileManager v0.0.0-00010101000000-000000000000
	example.com/dataBytesManager v0.0.0-00010101000000-000000000000
	example.com/demixer v0.0.0-00010101000000-000000000000
	example.com/metadataManager v0.0.0-00010101000000-000000000000
)
