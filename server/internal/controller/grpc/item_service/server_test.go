package item_service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	pb "github.com/illiafox/grpc-lake/gen/go/api/item_service/service/v1"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"server/internal/adapters/api/mocks"
	"server/internal/controller/grpc/interceptor/logger"
	"server/internal/controller/grpc/item_service/dto"
	"server/internal/domain/entity"
	app_errors "server/pkg/errors"
)

const TestID = "test"

var TestItem = entity.Item{
	Name:        "name",
	Data:        []byte("data"),
	Created:     time.Now().UTC(), // UTC() needed for correct mocks comparison (in TestUpdate)
	Description: "description",
}

// //

type TestGrpcServerSuite struct {
	suite.Suite

	itemMock   *mocks.MockItemUsecaseMockRecorder
	itemServer pb.ItemServiceServer
}

func TestGrpcServer(t *testing.T) {
	suite.Run(t, new(TestGrpcServerSuite))
}

func (suite *TestGrpcServerSuite) SetupSuite() {
	ctrl := gomock.NewController(suite.T())

	m := mocks.NewMockItemUsecase(ctrl)

	suite.itemMock = m.EXPECT()
	suite.itemServer = NewServer(m)
}

func (suite *TestGrpcServerSuite) TestGetItem() {
	t := suite.T()

	// context with logger
	ctx := logger.WithLogger(context.Background(), zap.NewNop())

	t.Run("No Error", func(t *testing.T) {
		suite.itemMock.GetItem(ctx, TestID).Return(TestItem, nil)

		resp, err := suite.itemServer.GetItem(ctx, &pb.GetItemRequest{Id: TestID})
		require.NoError(t, err)

		require.Equal(t, TestItem.Name, resp.Item.Name)
		require.Equal(t, TestItem.Data, resp.Item.Data)
		require.Equal(t, TestItem.Description, resp.Item.Description)
		require.Equal(t, TestItem.Created.Unix(), resp.Item.Created.Seconds)
	})

	t.Run("Error Not Found", func(t *testing.T) {
		suite.itemMock.GetItem(ctx, TestID).Return(entity.Item{}, entity.ErrItemNotFound)

		resp, err := suite.itemServer.GetItem(ctx, &pb.GetItemRequest{Id: TestID})
		require.Equal(t, status.Convert(err).Message(), entity.ErrItemNotFound.Error())
		require.Equal(t, status.Convert(err).Code(), codes.NotFound)

		require.Nil(t, resp)
	})

	t.Run("Internal Error", func(t *testing.T) {
		suite.itemMock.GetItem(ctx, TestID).Return(
			entity.Item{},
			// internal error
			app_errors.NewInternal("test", errors.New("test")),
		)

		resp, err := suite.itemServer.GetItem(ctx, &pb.GetItemRequest{Id: TestID})

		require.Equal(t, status.Convert(err).Code(), codes.Internal, "error code must be Internal")
		require.Nil(t, resp)
	})
}

func (suite *TestGrpcServerSuite) TestDeleteItem() {
	t := suite.T()

	// context with logger
	ctx := logger.WithLogger(context.Background(), zap.NewNop())

	t.Run("No Error", func(t *testing.T) {
		test := func(deleted bool) func(t *testing.T) {
			//
			return func(t *testing.T) {
				suite.itemMock.DeleteItem(ctx, TestID).Return( /*deleted*/ deleted, nil)

				resp, err := suite.itemServer.DeleteItem(ctx, &pb.DeleteItemRequest{Id: TestID})
				require.NoError(t, err)

				if deleted {

					require.True(t, resp.Deleted)
				} else {
					require.False(t, resp.Deleted)
				}
			}
			//
		}

		t.Run("Deleted", test(true))
		t.Run("Not Deleted", test(false))
	})

	t.Run("Internal Error", func(t *testing.T) {
		suite.itemMock.DeleteItem(ctx, TestID).Return( /*deleted*/ true, nil)

		resp, err := suite.itemServer.DeleteItem(ctx, &pb.DeleteItemRequest{Id: TestID})
		require.NoError(t, err)

		require.True(t, resp.Deleted)
	})
}

func (suite *TestGrpcServerSuite) TestCreateItem() {
	t := suite.T()

	// context with logger
	ctx := logger.WithLogger(context.Background(), zap.NewNop())

	t.Run("No Error", func(t *testing.T) {
		suite.itemMock.CreateItem(ctx,
			TestItem.Name, TestItem.Data, TestItem.Description,
		).Return(TestID, nil)

		resp, err := suite.itemServer.CreateItem(ctx, &pb.CreateItemRequest{
			Name:        TestItem.Name,
			Data:        TestItem.Data,
			Description: TestItem.Description,
		})

		require.NoError(t, err)
		require.Equal(t, TestID, resp.Id)
	})

	t.Run("Internal Error", func(t *testing.T) {
		suite.itemMock.CreateItem(ctx,
			TestItem.Name, TestItem.Data, TestItem.Description,
		).Return(TestID,
			app_errors.NewInternal("test", errors.New("test")),
		)

		resp, err := suite.itemServer.CreateItem(ctx, &pb.CreateItemRequest{
			Name:        TestItem.Name,
			Data:        TestItem.Data,
			Description: TestItem.Description,
		})

		require.Equal(t, status.Convert(err).Code(), codes.Internal, "error code must be Internal")
		require.Nil(t, resp)
	})
}

func (suite *TestGrpcServerSuite) TestUpdate() {
	t := suite.T()

	// context with logger
	ctx := logger.WithLogger(context.Background(), zap.NewNop())

	t.Run("No Error", func(t *testing.T) {

		test := func(created bool) func(t *testing.T) {
			//
			return func(t *testing.T) {
				suite.itemMock.UpdateItem(ctx, TestID, TestItem).Return( /*created*/ created, nil)

				resp, err := suite.itemServer.UpdateItem(ctx, &pb.UpdateItemRequest{
					Id:   TestID,
					Item: dto.Item(TestItem).ToProto(),
				})

				require.NoError(t, err)

				if created {
					require.True(t, resp.Created)
				} else {
					require.False(t, resp.Created)
				}
			}
			//
		}

		t.Run("Created", test(true))
		t.Run("Updated", test(false))
	})

	t.Run("Internal Error", func(t *testing.T) {
		suite.itemMock.UpdateItem(ctx, TestID, TestItem).Return( /*created*/ true,
			app_errors.NewInternal("test", errors.New("test")),
		)

		resp, err := suite.itemServer.UpdateItem(ctx, &pb.UpdateItemRequest{
			Id:   TestID,
			Item: dto.Item(TestItem).ToProto(),
		})

		require.Equal(t, status.Convert(err).Code(), codes.Internal, "error code must be Internal")
		require.Nil(t, resp)
	})
}
