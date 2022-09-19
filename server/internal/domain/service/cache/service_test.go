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

type CacheTestSuite struct {
	suite.Suite
	cache   *mocks.MockCacheStorage
	item    *mocks.MockItemStorage
	service CacheService
}

func TestCacheTestSuite(t *testing.T) {
	suite.Run(t, new(CacheTestSuite))
}

func (suite *CacheTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())

	suite.cache = mocks.NewMockCacheStorage(ctrl)
	suite.item = mocks.NewMockItemStorage(ctrl)
	suite.service = NewCacheService(suite.item, suite.cache)
}

func (suite *CacheTestSuite) TestGetItem() {
	t := suite.T()
	const id = "test"

	t.Run("Without Cache", func(t *testing.T) {
		suite.cache.EXPECT().
			GetItem(gomock.Any(), id).
			Return(entity.Item{}, entity.ErrItemNotFound).
			Times(1)

		i := entity.Item{Name: "1"}

		suite.item.EXPECT().
			GetItem(gomock.Any(), id).
			Return(i, nil).
			Times(1)

		suite.cache.EXPECT().
			SetItem(gomock.Any(), id, i).
			Times(1)

		g, err := suite.service.GetItem(context.Background(), id)
		require.NoError(t, err)
		require.Equal(t, i, g)
	})

	t.Run("With Cache", func(t *testing.T) {
		i := entity.Item{Name: "1"}

		suite.cache.EXPECT().
			GetItem(gomock.Any(), id).
			Return(i, nil).
			Times(1)

		g, err := suite.service.GetItem(context.Background(), id)
		require.NoError(t, err)
		require.Equal(t, i, g)
	})

	t.Run("With Error", func(t *testing.T) {
		suite.cache.EXPECT().
			GetItem(gomock.Any(), id).
			Return(entity.Item{}, entity.ErrItemNotFound).
			Times(1)

		suite.item.EXPECT().
			GetItem(gomock.Any(), id).
			Return(entity.Item{}, entity.ErrItemNotFound).
			Times(1)

		i, err := suite.service.GetItem(context.Background(), id)
		require.ErrorIs(t, entity.ErrItemNotFound, err)
		require.Zero(t, i)
	})
}

func (suite *CacheTestSuite) TestDeleteItem() {
	t := suite.T()
	const id = "test"

	t.Run("Deleted", func(t *testing.T) {

		suite.cache.EXPECT().
			DeleteItem(gomock.Any(), id).
			Return(nil).
			Times(1)

		suite.item.EXPECT().
			DeleteItem(gomock.Any(), id).
			Return( /*deleted*/ true, nil).
			Times(1)

		deleted, err := suite.service.DeleteItem(context.Background(), id)
		require.NoError(t, err)
		require.True(t, deleted)
	})

	t.Run("Not Deleted", func(t *testing.T) {
		suite.cache.EXPECT().
			DeleteItem(gomock.Any(), id).
			Return(nil).
			Times(1)

		suite.item.EXPECT().
			DeleteItem(gomock.Any(), id).
			Return( /*deleted*/ false, nil).
			Times(1)

		deleted, err := suite.service.DeleteItem(context.Background(), id)
		require.NoError(t, err)
		require.False(t, deleted)
	})
}

func (suite *CacheTestSuite) TestCreateItem() {
	t := suite.T()
	const id = "test"

	t.Run("No Error", func(t *testing.T) {

		i := entity.Item{
			Name:        "1",
			Data:        []byte("data"),
			Description: "test description",
		}

		suite.cache.EXPECT().
			DeleteItem(gomock.Any(), id).
			Return(nil).
			Times(1)

		suite.item.EXPECT().
			CreateItem(gomock.Any(), i.Name, i.Data, i.Description).
			Return(id, nil).
			Times(1)

		cid, err := suite.service.CreateItem(context.Background(), i.Name, i.Data, i.Description)
		require.NoError(t, err)
		require.Equal(t, id, cid)
	})

	t.Run("With Error", func(t *testing.T) {
		ErrCreate := errors.New("can't create")
		i := entity.Item{
			Name:        "1",
			Data:        []byte("data"),
			Description: "test description",
		}

		suite.item.EXPECT().
			CreateItem(gomock.Any(), i.Name, i.Data, i.Description).
			Return("", ErrCreate).
			Times(1)

		cid, err := suite.service.CreateItem(context.Background(), i.Name, i.Data, i.Description)
		require.ErrorIs(t, ErrCreate, err)
		require.Zero(t, cid)
	})
}

func (suite *CacheTestSuite) TestUpdateItem() {
	t := suite.T()
	const id = "test"

	t.Run("Not Updated", func(t *testing.T) {

		i := entity.Item{
			Name:        "1",
			Data:        []byte("data"),
			Description: "test description",
		}

		suite.cache.EXPECT().
			DeleteItem(gomock.Any(), id).
			Return(nil).
			Times(1)

		suite.item.EXPECT().
			UpdateItem(gomock.Any(), id, i).
			Return( /*updated*/ false, nil).
			Times(1)

		updated, err := suite.service.UpdateItem(context.Background(), id, i)
		require.NoError(t, err)
		require.False(t, updated)
	})

	t.Run("Updated", func(t *testing.T) {

		i := entity.Item{
			Name:        "1",
			Data:        []byte("data"),
			Description: "test description",
		}

		suite.cache.EXPECT().
			DeleteItem(gomock.Any(), id).
			Return(nil).
			Times(1)

		suite.item.EXPECT().
			UpdateItem(gomock.Any(), id, i).
			Return( /*updated*/ true, nil).
			Times(1)

		updated, err := suite.service.UpdateItem(context.Background(), id, i)
		require.NoError(t, err)
		require.True(t, updated)
	})

	t.Run("With Error", func(t *testing.T) {
		ErrUpdate := errors.New("can't update")

		suite.cache.EXPECT().
			DeleteItem(gomock.Any(), id).
			Return(nil).
			Times(1)

		suite.item.EXPECT().
			UpdateItem(gomock.Any(), id, entity.Item{}).
			Return( /*updated*/ false, ErrUpdate).
			Times(1)

		_, err := suite.service.UpdateItem(context.Background(), id, entity.Item{})
		require.ErrorIs(t, ErrUpdate, err)
	})
}
