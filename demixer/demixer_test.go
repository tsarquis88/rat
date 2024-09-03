package demixer

import (
	"testing"
	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	// "example.com/dataBytesManagerMock"
)

type DemixerTestSuite struct {
    suite.Suite
}

// func (suite *DemixerTestSuite) TestGenerateOneFile() {

// 	managerMock := dataBytesManagerMock.NewDataBytesManagerMock([]byte{25})
// 	Demix(managerMock)
// }

func TestDemixerTestSuite(t *testing.T) {
    suite.Run(t, new(DemixerTestSuite))
}
