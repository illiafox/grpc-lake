package item

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	entity2 "server/internal/domain/entity"
	"server/internal/domain/usecase/item/mocks"
)

type ItemTestSuite struct {
	suite.Suite
	item    *mocks.MockItemService
	sender  *mocks.MockEventService
	usecase ItemUsecase
}

func TestItemTestSuite(t *testing.T) {
	suite.Run(t, new(ItemTestSuite))
}

func (suite *ItemTestSuite) SetupSuite() {
	ctrl := gomock.NewController(suite.T())

	suite.item = mocks.NewMockItemService(ctrl)
	suite.sender = mocks.NewMockEventService(ctrl)
	suite.usecase = NewItemUsecase(suite.item, suite.sender)
}

func (suite *ItemTestSuite) TestGetItem() {
	t := suite.T()
	const id = "test"

	t.Run("No Error", func(t *testing.T) {
		i := entity2.Item{Name: "1"}

		suite.item.EXPECT().
			GetItem(gomock.Any(), id).
			Return(i, nil).
			Times(1)

		g, err := suite.usecase.GetItem(context.Background(), id)
		require.NoError(t, err)
		require.Equal(t, i, g)
	})

	t.Run("Error", func(t *testing.T) {
		ErrGet := errors.New("get error")

		suite.item.EXPECT().
			GetItem(gomock.Any(), id).
			Return(entity2.Item{}, ErrGet).
			Times(1)

		g, err := suite.usecase.GetItem(context.Background(), id)
		require.ErrorIs(t, ErrGet, err)
		require.Zero(t, g)
	})
}

func (suite *ItemTestSuite) TestCreateItem() {
	t := suite.T()

	const (
		id          = "1"
		name        = "test"
		description = "description"
	)
	var data = []byte("data")

	t.Run("No Error", func(t *testing.T) {

		suite.item.EXPECT().
			CreateItem(gomock.Any(), name, data, description).
			Return(id, nil).Times(1)

		suite.sender.EXPECT().
			SendItemEvent(gomock.Any(), id, entity2.CreateEvent).
			Return(nil).Times(1)

		cid, err := suite.usecase.CreateItem(context.Background(), name, data, description)
		require.NoError(t, err)
		require.Equal(t, id, cid)
	})

	t.Run("Error", func(t *testing.T) {
		ErrCreate := errors.New("create error")

		suite.item.EXPECT().
			CreateItem(gomock.Any(), name, data, description).
			Return("", ErrCreate).Times(1)

		cid, err := suite.usecase.CreateItem(context.Background(), name, data, description)
		require.ErrorIs(t, ErrCreate, err)
		require.Zero(t, cid)
	})
}

func (suite *ItemTestSuite) TestDeleteItem() {

	t := suite.T()
	const id = "test"

	t.Run("No Error", func(t *testing.T) {
		suite.item.EXPECT().
			DeleteItem(gomock.Any(), id).
			Return( /*deleted*/ true, nil).Times(1)

		suite.sender.EXPECT().
			SendItemEvent(gomock.Any(), id, entity2.DeleteEvent).
			Return(nil).Times(1)

		deleted, err := suite.usecase.DeleteItem(context.Background(), id)
		require.NoError(t, err)
		require.True(t, deleted)
	})

	t.Run("Error", func(t *testing.T) {
		ErrDelete := errors.New("delete error")

		suite.item.EXPECT().
			DeleteItem(gomock.Any(), id).
			Return( /*deleted*/ false, ErrDelete).Times(1)

		deleted, err := suite.usecase.DeleteItem(context.Background(), id)
		require.ErrorIs(t, ErrDelete, err)
		require.False(t, deleted)
	})
}
