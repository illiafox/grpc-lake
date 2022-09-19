package event

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"server/internal/domain/entity"
	"server/internal/domain/service/event/mocks"
	"server/internal/domain/service/event/model"
)

func TestEventService(t *testing.T) {
	ctrl := gomock.NewController(t)

	message := mocks.NewMockMessageStorage(ctrl)

	service := NewEventService(message)

	t.Run("SendItemEvent", func(t *testing.T) {
		const id = "test"

		data, err := model.NewMessage(id, entity.CreateEvent)
		require.NoError(t, err)

		message.EXPECT().
			SendMessageJSON(gomock.Any(), data).
			Return(nil).Times(1)

		err = service.SendItemEvent(context.Background(), "test", entity.CreateEvent)
		require.NoError(t, err)
	})
}
