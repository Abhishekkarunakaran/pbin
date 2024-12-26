package service

import (
	"context"
	"errors"
	"testing"

	"github.com/Abhishekkarunakaran/pbin/src/core/domain"
	"github.com/Abhishekkarunakaran/pbin/src/core/ports/mock_ports"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSaveContent(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_ports.NewMockRepository(ctrl)

	testCases := []struct {
		name        string
		payload     *domain.Payload
		mockErr     error
		expectedErr error
		wantErr     bool
	}{
		{
			name:        "#1 Save Content-Postive Test",
			payload:     &domain.Payload{},
			mockErr:     nil,
			expectedErr: nil,
			wantErr:     false,
		},
		{
			name: "#2 Save Content-Negative Test",
			payload: &domain.Payload{},
			mockErr:     errors.New("Some error"),
			expectedErr: ErrSaveData,
			wantErr:     true,
		},
	}

	pbinService := NewPbinService(mockRepo)

	ctx := context.Background()

	for _, testCase := range testCases {
		mockRepo.EXPECT().AddData(ctx, gomock.Any(), gomock.Any()).Return(testCase.mockErr)

		_, err := pbinService.SaveContent(ctx, testCase.payload)
		if testCase.wantErr {
			assert.ErrorIs(t, err, testCase.expectedErr)
		}

	}

}
