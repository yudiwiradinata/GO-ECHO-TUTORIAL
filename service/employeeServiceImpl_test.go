package service_test

import (
	"GO-ECHO-TUTORIAL/mock"
	"GO-ECHO-TUTORIAL/model"
	"GO-ECHO-TUTORIAL/service"
	"testing"

	"github.com/pkg/errors"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetEmployeeByName(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	empReturn := model.Employee{
		Name:   "yudi",
		Id:     "1",
		Salary: "1000",
		Age:    "10",
	}

	em := []model.Employee{empReturn}
	empReturns := &model.Employees{
		Employees: em,
	}

	mockRepo := mock.NewMockEmployeeRepositoryInterface(mockCtrl)
	//	mockEmployeeService := mock.NewMockEmployeeServiceInterface(mockCtrl)
	empServiceMock := service.NewEmployeeServiceInterface(mockRepo)

	t.Run("GIVEN GetEmployeeByName", func(t *testing.T) {
		t.Run("WHEN error", func(t *testing.T) {
			mockRepo.EXPECT().GetEmployeeByName(gomock.Any()).Return(nil, errors.New("test"))
			_, err := empServiceMock.GetEmployeeByName("yudi")
			require.EqualError(t, err, "test")
		})

		t.Run("WHEN success", func(t *testing.T) {
			mockRepo.EXPECT().GetEmployeeByName(gomock.Any()).Return(empReturns, nil)
			res, _ := empServiceMock.GetEmployeeByName("yudi")
			require.Equal(t, empReturns, res)
		})
	})

	t.Run("GIVEN GetEmployee", func(t *testing.T) {
		t.Run("WHEN error", func(t *testing.T) {
			mockRepo.EXPECT().GetEmployee().Return(nil, errors.New("test"))
			_, err := empServiceMock.GetEmp()
			require.EqualError(t, err, "test")
		})

		t.Run("WHEN success", func(t *testing.T) {
			mockRepo.EXPECT().GetEmployee().Return(empReturns, nil)
			res, _ := empServiceMock.GetEmp()
			require.Equal(t, empReturns, res)
		})
	})

}
