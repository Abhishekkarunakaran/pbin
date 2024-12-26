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
		mockErr     error
		expectedErr error
		wantErr     bool
	}{
		{
			name:        "#1 Save Content-Postive Test",
			mockErr:     nil,
			expectedErr: nil,
			wantErr:     false,
		},
		{
			name:        "#2 Save Content-Negative Test",
			mockErr:     errors.New("Some error"),
			expectedErr: ErrSaveData,
			wantErr:     true,
		},
	}

	pbinService := NewPbinService(mockRepo)

	ctx := context.Background()

	for _, testCase := range testCases {
		mockRepo.EXPECT().AddData(ctx, gomock.Any(), gomock.Any()).Return(testCase.mockErr)

		_, err := pbinService.SaveContent(ctx, &domain.Payload{})

		if testCase.wantErr {
			assert.ErrorIs(t, err, testCase.expectedErr)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestGetContent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // Ensure all expectations are met
	mockRepo := mock_ports.NewMockRepository(ctrl)

	testCases := []struct {
		name        string
		payload     *domain.DataRequest
		mockData    *domain.Data
		mockErr     error
		expectedErr error
		wantErr     bool
	}{
		{
			name:        "#1 Get Content- Negative Test (Repo Error)",
			payload:     &domain.DataRequest{Password: "password"},
			mockData:    nil,
			mockErr:     errors.New("Some Error"),
			expectedErr: ErrGetData,
			wantErr:     true,
		},
		{
			name:        "#2 Get Content- Negative Test (Missing Data)",
			payload:     &domain.DataRequest{Password: "password"},
			mockData:    &domain.Data{Password: "", Content: "nonce" + "ciphertext"},
			mockErr:     nil,
			expectedErr: ErrGetDataAbsent,
			wantErr:     true,
		},
		{
			name:        "#3 Get Content- Negative Test (Incorrect Password)",
			payload:     &domain.DataRequest{Password: "wrong_password"},
			mockData:    &domain.Data{Password: "$2y$10$...hashed_password...", Content: "nonce" + "ciphertext"},
			mockErr:     nil,
			expectedErr: ErrIncorrectPassword,
			wantErr:     true,
		},
	}

	pbinService := NewPbinService(mockRepo)

	ctx := context.Background()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.mockData != nil {
				mockRepo.EXPECT().GetData(gomock.Any(), gomock.Any()).Return(testCase.mockData, testCase.mockErr)
			} else {
				mockRepo.EXPECT().GetData(gomock.Any(), gomock.Any()).Return(nil, testCase.mockErr)
			}

			content, err := pbinService.GetContent(ctx, testCase.payload)

			if testCase.wantErr {
				assert.ErrorIs(t, err, testCase.expectedErr)
				assert.Nil(t, content) 
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, content)
			}
		})
	}
}
