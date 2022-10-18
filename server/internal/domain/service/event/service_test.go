package event

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"server/internal/domain/entity"
	"server/internal/domain/service/event/mocks"
)

const TestID = "test"

func TestEventService(t *testing.T) {
	ctrl := gomock.NewController(t)

	messageMock := mocks.NewMockMessageStorage(ctrl)
	eventService := NewEventService(messageMock)

	t.Run("SendItemEvent", func(t *testing.T) {
		actions := []entity.Action{entity.CreateEvent, entity.DeleteEvent, entity.UpdateEvent}

		for _, action := range actions {
			t.Run(string(action), func(t *testing.T) {
				msg := entity.NewMessage(TestID, entity.CreateEvent)

				messageMock.EXPECT().
					SendMessageJSON(gomock.Any(), msg).
					Return(nil)

				err := eventService.SendItemEvent(context.Background(), TestID, entity.CreateEvent)
				require.NoError(t, err)
			})
		}
	})
}
