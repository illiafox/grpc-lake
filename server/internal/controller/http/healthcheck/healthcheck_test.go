package healthcheck

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"server/internal/adapters/api/mocks"
	"testing"
)

type TestHealthCheckSuite struct {
	suite.Suite

	itemMock *mocks.MockItemUsecaseMockRecorder
	handler  http.HandlerFunc
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestHealthCheckSuite))
}

func (s *TestHealthCheckSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())
	itemUsecase := mocks.NewMockItemUsecase(ctrl)

	// //

	s.itemMock = itemUsecase.EXPECT()
	s.handler = NewServerHealthCheck(itemUsecase, zap.NewNop()).HealthCheck
}

func (s *TestHealthCheckSuite) TestOk() {
	t := s.T()
	const id = "test"

	// Mocks
	s.itemMock.CreateItem(gomock.Any(),
		ExpectedItem.Name, ExpectedItem.Data, ExpectedItem.Description,
	).Return(id, nil)
	s.itemMock.GetItem(gomock.Any(), id).Return(ExpectedItem, nil)
	s.itemMock.GetItem(gomock.Any(), id).Return(ExpectedItem, nil)
	s.itemMock.DeleteItem(gomock.Any(), id).Return( /*deleted*/ true, nil)

	// Call HealthCheck
	recorder := httptest.NewRecorder()
	s.handler(recorder,
		httptest.NewRequest(http.MethodGet, "/healthcheck", nil),
	)

	// Check Response
	var resp Response

	decoder := json.NewDecoder(recorder.Body)
	decoder.DisallowUnknownFields()
	require.NoError(t, decoder.Decode(&resp))

	require.Equal(t, http.StatusOK, recorder.Code)
	require.Equal(t, "ok", resp.Message)
}

// TODO: add more tests
