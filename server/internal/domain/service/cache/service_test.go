package cache

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"server/internal/domain/entity"
	"server/internal/domain/service/cache/mocks"
)

const TestID = "test"

var TestItem = entity.Item{
	Name:        "1",
	Data:        []byte("data"),
	Description: "test description",
}

// //

type CacheTestSuite struct {
	suite.Suite
	cacheMock    *mocks.MockCacheStorage
	itemMock     *mocks.MockItemStorage
	cacheService CacheService
}

func TestCacheTestSuite(t *testing.T) {
	suite.Run(t, new(CacheTestSuite))
}

func (suite *CacheTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())

	suite.cacheMock = mocks.NewMockCacheStorage(ctrl)
	suite.itemMock = mocks.NewMockItemStorage(ctrl)
	suite.cacheService = NewCacheService(suite.itemMock, suite.cacheMock)
}

func (suite *CacheTestSuite) TestGetItem() {
	t := suite.T()

	t.Run("Without Cache", func(t *testing.T) {
		suite.cacheMock.EXPECT().
			GetItem(gomock.Any(), TestID).
			Return(entity.Item{}, entity.ErrItemNotFound)

		suite.itemMock.EXPECT().
			GetItem(gomock.Any(), TestID).
			Return(TestItem, nil)

		suite.cacheMock.EXPECT().
			SetItem(gomock.Any(), TestID, TestItem)

		g, err := suite.cacheService.GetItem(context.Background(), TestID)
		require.NoError(t, err)
		require.Equal(t, TestItem, g)
	})

	t.Run("With Cache", func(t *testing.T) {
		suite.cacheMock.EXPECT().
			GetItem(gomock.Any(), TestID).
			Return(TestItem, nil)

		g, err := suite.cacheService.GetItem(context.Background(), TestID)
		require.NoError(t, err)
		require.Equal(t, TestItem, g)
	})

	t.Run("With Error", func(t *testing.T) {
		suite.cacheMock.EXPECT().
			GetItem(gomock.Any(), TestID).
			Return(entity.Item{}, entity.ErrItemNotFound)

		suite.itemMock.EXPECT().
			GetItem(gomock.Any(), TestID).
			Return(entity.Item{}, entity.ErrItemNotFound)

		cTestID, err := suite.cacheService.GetItem(context.Background(), TestID)
		require.ErrorIs(t, entity.ErrItemNotFound, err)
		require.Zero(t, cTestID)
	})
}

func (suite *CacheTestSuite) TestDeleteItem() {
	t := suite.T()

	t.Run("Deleted", func(t *testing.T) {

		suite.cacheMock.EXPECT().
			DeleteItem(gomock.Any(), TestID).
			Return(nil)

		suite.itemMock.EXPECT().
			DeleteItem(gomock.Any(), TestID).
			Return( /*deleted*/ true, nil)

		deleted, err := suite.cacheService.DeleteItem(context.Background(), TestID)
		require.NoError(t, err)
		require.True(t, deleted)
	})

	t.Run("Not Deleted", func(t *testing.T) {
		suite.cacheMock.EXPECT().
			DeleteItem(gomock.Any(), TestID).
			Return(nil)

		suite.itemMock.EXPECT().
			DeleteItem(gomock.Any(), TestID).
			Return( /*deleted*/ false, nil)

		deleted, err := suite.cacheService.DeleteItem(context.Background(), TestID)
		require.NoError(t, err)
		require.False(t, deleted)
	})
}

func (suite *CacheTestSuite) TestCreateItem() {
	t := suite.T()

	t.Run("No Error", func(t *testing.T) {

		suite.cacheMock.EXPECT().
			DeleteItem(gomock.Any(), TestID).
			Return(nil)

		suite.itemMock.EXPECT().
			CreateItem(gomock.Any(), TestItem.Name, TestItem.Data, TestItem.Description).
			Return(TestID, nil)

		cTestID, err := suite.cacheService.CreateItem(context.Background(), TestItem.Name, TestItem.Data, TestItem.Description)
		require.NoError(t, err)
		require.Equal(t, TestID, cTestID)
	})

	t.Run("With Error", func(t *testing.T) {
		ErrCreate := errors.New("can't create")

		suite.itemMock.EXPECT().
			CreateItem(gomock.Any(), TestItem.Name, TestItem.Data, TestItem.Description).
			Return("", ErrCreate)

		cTestID, err := suite.cacheService.CreateItem(context.Background(), TestItem.Name, TestItem.Data, TestItem.Description)
		require.ErrorIs(t, ErrCreate, err)
		require.Zero(t, cTestID)
	})
}

func (suite *CacheTestSuite) TestUpdateItem() {
	t := suite.T()

	t.Run("Not Updated", func(t *testing.T) {

		suite.cacheMock.EXPECT().
			DeleteItem(gomock.Any(), TestID).
			Return(nil)

		suite.itemMock.EXPECT().
			UpdateItem(gomock.Any(), TestID, TestItem).
			Return( /*updated*/ false, nil)

		updated, err := suite.cacheService.UpdateItem(context.Background(), TestID, TestItem)
		require.NoError(t, err)
		require.False(t, updated)
	})

	t.Run("Updated", func(t *testing.T) {

		suite.cacheMock.EXPECT().
			DeleteItem(gomock.Any(), TestID).
			Return(nil)

		suite.itemMock.EXPECT().
			UpdateItem(gomock.Any(), TestID, TestItem).
			Return( /*updated*/ true, nil)

		updated, err := suite.cacheService.UpdateItem(context.Background(), TestID, TestItem)
		require.NoError(t, err)
		require.True(t, updated)
	})

	t.Run("With Error", func(t *testing.T) {
		ErrUpdate := errors.New("can't update")

		suite.cacheMock.EXPECT().
			DeleteItem(gomock.Any(), TestID).
			Return(nil)

		suite.itemMock.EXPECT().
			UpdateItem(gomock.Any(), TestID, TestItem).
			Return( /*updated*/ false, ErrUpdate)

		_, err := suite.cacheService.UpdateItem(context.Background(), TestID, TestItem)
		require.ErrorIs(t, ErrUpdate, err)
	})
}
