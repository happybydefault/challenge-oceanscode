package factory

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type factoryUnitTestSuite struct {
	suite.Suite
	adapter *Factory
}

func (s *factoryUnitTestSuite) SetupSuite() {
	s.adapter = New()
}

func TestFactoryUnitTestSuite(t *testing.T) {
	suite.Run(t, &factoryUnitTestSuite{})
}

func (s *factoryUnitTestSuite) TestSample() {
	// code here
	// Assert
	s.Assert().Equal(1, 1)
}

func (s *factoryUnitTestSuite) TestStartAssemblingProcessPositive() {
	amount := 13

	ch := s.adapter.StartAssemblingProcess(amount)

	for v := range ch {
		s.Require().NotEmpty(v)

		assembled := "Assembled"
		s.Require().Equal(assembled, v.Chassis)
		s.Require().Equal(assembled, v.Tires)
		s.Require().Equal(assembled, v.Engine)
		s.Require().Equal(assembled, v.Electronics)
		s.Require().Equal(assembled, v.Dash)
		s.Require().Equal(assembled, v.Sits)
		s.Require().Equal(assembled, v.Windows)
		s.Require().Equal(assembled, v.Windows)
	}
}
