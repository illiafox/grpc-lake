package item

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"server/internal/domain/entity"
	"server/internal/domain/usecase/item/mocks"
)

const TestID = "test"

var TestItem = entity.Item{
	Name:        "name",
	Data:        []byte("data"),
	Created:     time.Now(),
	Description: "description",
}

// //

type ItemTestSuite struct {
	suite.Suite
	mockItem    *mocks.MockItemService
	mockSender  *mocks.MockEventService
	itemUsecase ItemUsecase
}

func TestItemTestSuite(t *testing.T) {
	suite.Run(t, new(ItemTestSuite))
}

func (suite *ItemTestSuite) SetupSuite() {
	ctrl := gomock.NewController(suite.T())

	suite.mockItem = mocks.NewMockItemService(ctrl)
	suite.mockSender = mocks.NewMockEventService(ctrl)
	suite.itemUsecase = NewItemUsecase(suite.mockItem, suite.mockSender)
}

func (suite *ItemTestSuite) TestGetItem() {
	t := suite.T()

	t.Run("No Error", func(t *testing.T) {
		i := entity.Item{Name: "1"}

		suite.mockItem.EXPECT().
			GetItem(gomock.Any(), TestID).
			Return(i, nil)

		g, err := suite.itemUsecase.GetItem(context.Background(), TestID)
		require.NoError(t, err)
		require.Equal(t, i, g)
	})

	t.Run("Error", func(t *testing.T) {
		ErrGet := errors.New("get error")

		suite.mockItem.EXPECT().
			GetItem(gomock.Any(), TestID).
			Return(entity.Item{}, ErrGet)

		g, err := suite.itemUsecase.GetItem(context.Background(), TestID)
		require.ErrorIs(t, ErrGet, err)
		require.Zero(t, g)
	})
}

func (suite *ItemTestSuite) TestCreateItem() {
	t := suite.T()

	t.Run("No Error", func(t *testing.T) {

		suite.mockItem.EXPECT().
			CreateItem(gomock.Any(), TestItem.Name, TestItem.Data, TestItem.Description).
			Return(TestID, nil)

		suite.mockSender.EXPECT().
			SendItemEvent(gomock.Any(), TestID, entity.CreateEvent).
			Return(nil)

		cTestID, err := suite.itemUsecase.CreateItem(context.Background(),
			TestItem.Name, TestItem.Data, TestItem.Description,
		)
		require.NoError(t, err)
		require.Equal(t, TestID, cTestID)
	})

	t.Run("Error", func(t *testing.T) {
		ErrCreate := errors.New("create error")

		suite.mockItem.EXPECT().
			CreateItem(gomock.Any(), TestItem.Name, TestItem.Data, TestItem.Description).
			Return("", ErrCreate)

		cTestID, err := suite.itemUsecase.CreateItem(context.Background(),
			TestItem.Name, TestItem.Data, TestItem.Description,
		)
		require.ErrorIs(t, ErrCreate, err)
		require.Zero(t, cTestID)
	})
}

func (suite *ItemTestSuite) TestDeleteItem() {

	t := suite.T()
	t.Run("No Error", func(t *testing.T) {
		suite.mockItem.EXPECT().
			DeleteItem(gomock.Any(), TestID).
			Return( /*deleted*/ true, nil)

		suite.mockSender.EXPECT().
			SendItemEvent(gomock.Any(), TestID, entity.DeleteEvent).
			Return(nil)

		deleted, err := suite.itemUsecase.DeleteItem(context.Background(), TestID)
		require.NoError(t, err)
		require.True(t, deleted)
	})

	t.Run("Error", func(t *testing.T) {
		ErrDelete := errors.New("delete error")

		suite.mockItem.EXPECT().
			DeleteItem(gomock.Any(), TestID).
			Return( /*deleted*/ false, ErrDelete)

		deleted, err := suite.itemUsecase.DeleteItem(context.Background(), TestID)
		require.ErrorIs(t, ErrDelete, err)
		require.False(t, deleted)
	})
}
