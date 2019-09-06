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

	empRequest := model.EmployeeRequest{
		Name:   "yudi",
		Salary: "1000",
		Age:    "10",
	}
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
	mockCache := mock.NewMockCacheServiceInterface(mockCtrl)
	params := service.EmpServiceParams{
		Repository: mockRepo,
		Cache:      mockCache,
	}

	//	mockEmployeeService := mock.NewMockEmployeeServiceInterface(mockCtrl)
	empServiceMock := service.NewEmployeeServiceInterface(params)

	t.Run("GIVEN GetEmployeeByName", func(t *testing.T) {
		t.Run("WHEN error", func(t *testing.T) {
			mockCache.EXPECT().FindCache(gomock.Any(), gomock.Any()).Return(errors.New("empty"))
			mockRepo.EXPECT().GetEmployeeByName(gomock.Any()).Return(nil, errors.New("test"))
			_, err := empServiceMock.GetEmployeeByName("yudi")
			require.EqualError(t, err, "test")
		})
		t.Run("WHEN success cache nil", func(t *testing.T) {
			mockRepo.EXPECT().GetEmployeeByName(gomock.Any()).Return(empReturns, nil)
			mockCache.EXPECT().FindCache(gomock.Any(), gomock.Any()).Return(errors.New("empty"))
			mockCache.EXPECT().CreateCache(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			res, _ := empServiceMock.GetEmployeeByName("yudi")
			require.Equal(t, empReturns, res)
		})
		t.Run("WHEN success cache exist", func(t *testing.T) {
			mockCache.EXPECT().FindCache(gomock.Any(), gomock.Any()).Return(nil)
			res, _ := empServiceMock.GetEmployeeByName("yudi")
			require.Equal(t, &model.Employees{Employees: []model.Employee{model.Employee{}}}, res)
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

	t.Run("GIVEN CreateEmp ", func(t *testing.T) {
		t.Run("When error", func(t *testing.T) {
			mockRepo.EXPECT().InsertEmployee(empRequest).Return(nil, errors.New("error"))
			err := empServiceMock.CreateEmp(empRequest)
			require.EqualError(t, err, "error")
		})

		t.Run("When success", func(t *testing.T) {
			mockRepo.EXPECT().InsertEmployee(empRequest).Return(&empReturn, nil)
			mockCache.EXPECT().CreateCache(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			err := empServiceMock.CreateEmp(empRequest)
			require.Equal(t, nil, err)
		})
	})

	t.Run("GIVEN DeleteEmp ", func(t *testing.T) {
		t.Run("When error", func(t *testing.T) {
			mockCache.EXPECT().DeleteCache(gomock.Any()).Return(nil)
			mockRepo.EXPECT().DeleteEmployee(empRequest.Name).Return(false, errors.New("error"))
			_, err := empServiceMock.DeleteEmp(empRequest.Name)
			require.EqualError(t, err, "error")
		})

		t.Run("When success", func(t *testing.T) {
			mockCache.EXPECT().DeleteCache(gomock.Any()).Return(nil)
			mockRepo.EXPECT().DeleteEmployee(empRequest.Name).Return(true, nil)
			res, _ := empServiceMock.DeleteEmp(empRequest.Name)
			require.Equal(t, true, res)
		})
	})
}
