package cli_test

import (
	"fmt"
	"go-hexagonal/adapters/cli"
	mock_application "go-hexagonal/application/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 25.99
	productStatus := "enabled"
	productId := "abc"

	producMock := mock_application.NewMockProductInterface(ctrl)
	producMock.EXPECT().GetId().Return(productId).AnyTimes()
	producMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	producMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	producMock.EXPECT().GetName().Return(productName).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(producMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(producMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(producMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(producMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s", productId, productName, productPrice, productStatus)
	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID %s with the name %s has been enabled", productId, productName)
	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID %s with the name %s has been disabled", productId, productName)
	result, err = cli.Run(service, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s", productId, productName, productPrice, productStatus)
	result, err = cli.Run(service, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}