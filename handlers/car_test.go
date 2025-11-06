package handlers

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hse-trpo-taxi/backend/errors"
	"github.com/hse-trpo-taxi/backend/models"
	mock_cars "github.com/hse-trpo-taxi/backend/usecases/cars/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestCarHandler_GetCars(t *testing.T) {
	type mockBehaviour func(ctx context.Context, s *mock_cars.MockCarUseCase)

	currTime := time.Now().Truncate(time.Second)

	table := []struct {
		name          string
		mockBehaviour mockBehaviour
		expectedCode  int
		expectedBody  string
	}{
		{
			name: "success empty",
			mockBehaviour: func(ctx context.Context, s *mock_cars.MockCarUseCase) {
				s.EXPECT().GetCars().Return([]*models.Car{}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: `[]`,
		},
		{
			name: "success valid car",
			mockBehaviour: func(ctx context.Context, s *mock_cars.MockCarUseCase) {
				s.EXPECT().GetCars().Return([]*models.Car{
					{
						ID:           1,
						DriverID:     2,
						Brand:        "Cytroen",
						Model:        "C10000",
						Year:         2456,
						LicensePlate: "1ABCD",
						Color:        "Black",
						CreatedAt:    currTime,
						UpdatedAt:    currTime,
					},
				}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: fmt.Sprintf(
				`[{"id":1,"driver_id":2,"brand":"Cytroen","model":"C10000","year":2456,"license_plate":"1ABCD","color":"Black","created_at":"%s","updated_at":"%s"}]`,
				currTime.Format(time.RFC3339),
				currTime.Format(time.RFC3339),
			),
		},
		{
			name: "Error",
			mockBehaviour: func(ctx context.Context, s *mock_cars.MockCarUseCase) {
				s.EXPECT().GetCars().Return(nil, &errors.ValidationError{})
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"errors":"internal server error"}`,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCar := mock_cars.NewMockCarUseCase(ctrl)
			tt.mockBehaviour(ctx, mockCar)

			lgr := slog.New(slog.NewJSONHandler(os.Stderr, nil))

			handler := NewCarHandler(mockCar, lgr)
			router := mux.NewRouter()
			router.HandleFunc("/api/cars", handler.GetCars).Methods(http.MethodGet)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/api/cars", nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Equal(t, tt.expectedBody, w.Body.String())
		})
	}
}

func TestCarHandler_GetCarById(t *testing.T) {
	type mockBehaviour func(ctx context.Context, s *mock_cars.MockCarUseCase)

	currTime := time.Now().Truncate(time.Second)

	table := []struct {
		name          string
		mockBehaviour mockBehaviour
		expectedCode  int
		expectedBody  string
	}{
		{
			name: "Error",
			mockBehaviour: func(ctx context.Context, s *mock_cars.MockCarUseCase) {
				s.EXPECT().GetCarById(uint32(1)).Return(nil, &errors.ValidationError{})
			},
			expectedCode: http.StatusNotFound,
			expectedBody: `{"errors":"internal server error"}`,
		},
		{
			name: "Success",
			mockBehaviour: func(ctx context.Context, s *mock_cars.MockCarUseCase) {
				s.EXPECT().GetCarById(uint32(1)).Return(&models.Car{
					ID:           1,
					DriverID:     2,
					Brand:        "Cytroen",
					Model:        "C10000",
					Year:         2456,
					LicensePlate: "1ABCD",
					Color:        "Black",
					CreatedAt:    currTime,
					UpdatedAt:    currTime,
				}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: fmt.Sprintf(
				`{"id":1,"driver_id":2,"brand":"Cytroen","model":"C10000","year":2456,"license_plate":"1ABCD","color":"Black","created_at":"%s","updated_at":"%s"}`,
				currTime.Format(time.RFC3339),
				currTime.Format(time.RFC3339),
			),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCar := mock_cars.NewMockCarUseCase(ctrl)
			tt.mockBehaviour(ctx, mockCar)

			lgr := slog.New(slog.NewJSONHandler(os.Stderr, nil))
			handler := NewCarHandler(mockCar, lgr)
			w := httptest.NewRecorder()

			req := httptest.NewRequest(http.MethodGet, "/api/cars/1", nil)
			router := mux.NewRouter()

			router.HandleFunc("/api/cars/{id}", handler.GetCarById).Methods(http.MethodGet)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Equal(t, tt.expectedBody, w.Body.String())
		})
	}
}

func TestCarHandler_CreateCar(t *testing.T) {
	type mockBehaviour func(ctx context.Context, s *mock_cars.MockCarUseCase, request *models.CreateCarModel)

	currTime := time.Now().Truncate(time.Second)

	table := []struct {
		name          string
		inputBody     string
		parseBody     *models.CreateCarModel
		mockBehaviour mockBehaviour
		expectedCode  int
		expectedBody  string
	}{
		{
			name:      "Error empty body",
			inputBody: "",
			parseBody: &models.CreateCarModel{},
			mockBehaviour: func(ctx context.Context, s *mock_cars.MockCarUseCase, request *models.CreateCarModel) {
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"errors":"bad request"}`,
		},
		{
			name:      "Success",
			inputBody: `{"driver_id":2,"brand":"Cytroen","model":"C10000","year":2456,"license_plate":"1ABCD","color":"Black"}`,
			parseBody: &models.CreateCarModel{
				DriverID:     2,
				Brand:        "Cytroen",
				Model:        "C10000",
				Year:         2456,
				LicensePlate: "1ABCD",
				Color:        "Black",
			},
			mockBehaviour: func(ctx context.Context, s *mock_cars.MockCarUseCase, request *models.CreateCarModel) {
				s.EXPECT().CreateCar(request).Return(
					&models.Car{
						ID:           1,
						DriverID:     2,
						Brand:        "Cytroen",
						Model:        "C10000",
						Year:         2456,
						LicensePlate: "1ABCD",
						Color:        "Black",
						CreatedAt:    currTime,
						UpdatedAt:    currTime,
					}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: fmt.Sprintf(
				`{"id":1,"driver_id":2,"brand":"Cytroen","model":"C10000","year":2456,"license_plate":"1ABCD","color":"Black","created_at":"%s","updated_at":"%s"}`,
				currTime.Format(time.RFC3339),
				currTime.Format(time.RFC3339),
			),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCar := mock_cars.NewMockCarUseCase(ctrl)
			tt.mockBehaviour(ctx, mockCar, tt.parseBody)

			lgr := slog.New(slog.NewJSONHandler(os.Stderr, nil))
			handler := NewCarHandler(mockCar, lgr)
			w := httptest.NewRecorder()

			req := httptest.NewRequest(http.MethodPost, "/api/cars", bytes.NewBufferString(tt.inputBody))
			router := mux.NewRouter()

			router.HandleFunc("/api/cars", handler.CreateCar).Methods(http.MethodPost)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Equal(t, tt.expectedBody, w.Body.String())
		})
	}
}

func TestCarHandler_UpdateCar(t *testing.T) {
	type mockBehaviour func(ctx context.Context, s *mock_cars.MockCarUseCase, request *models.UpdateCarModel)

	currTime := time.Now().Truncate(time.Second)

	table := []struct {
		name          string
		inputBody     string
		parseBody     *models.UpdateCarModel
		mockBehaviour mockBehaviour
		expectedCode  int
		expectedBody  string
	}{
		{
			name:      "Error empty body",
			inputBody: "",
			parseBody: &models.UpdateCarModel{},
			mockBehaviour: func(ctx context.Context, s *mock_cars.MockCarUseCase, request *models.UpdateCarModel) {
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"errors":"bad request"}`,
		},
		{
			name:      "Success",
			inputBody: `{"driver_id":3,"color":"Pink"}`,
			parseBody: &models.UpdateCarModel{
				DriverID: 3,
				Color:    "Pink",
			},
			mockBehaviour: func(ctx context.Context, s *mock_cars.MockCarUseCase, request *models.UpdateCarModel) {
				s.EXPECT().UpdateCar(uint32(1), request).Return(
					&models.Car{
						ID:           1,
						DriverID:     3,
						Brand:        "Cytroen",
						Model:        "C10000",
						Year:         2456,
						LicensePlate: "1ABCD",
						Color:        "Pink",
						CreatedAt:    currTime,
						UpdatedAt:    currTime,
					}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: fmt.Sprintf(
				`{"id":1,"driver_id":3,"brand":"Cytroen","model":"C10000","year":2456,"license_plate":"1ABCD","color":"Pink","created_at":"%s","updated_at":"%s"}`,
				currTime.Format(time.RFC3339),
				currTime.Format(time.RFC3339),
			),
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCar := mock_cars.NewMockCarUseCase(ctrl)
			tt.mockBehaviour(ctx, mockCar, tt.parseBody)

			lgr := slog.New(slog.NewJSONHandler(os.Stderr, nil))
			handler := NewCarHandler(mockCar, lgr)
			w := httptest.NewRecorder()

			req := httptest.NewRequest(http.MethodPut, "/api/cars/1", bytes.NewBufferString(tt.inputBody))
			router := mux.NewRouter()

			router.HandleFunc("/api/cars/{id}", handler.UpdateCar).Methods(http.MethodPut)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Equal(t, tt.expectedBody, w.Body.String())
		})
	}
}

func TestCarHandler_DeleteCar(t *testing.T) {
	type mockBehaviour func(ctx context.Context, s *mock_cars.MockCarUseCase)

	table := []struct {
		name          string
		mockBehaviour mockBehaviour
		expectedCode  int
		expectedBody  string
	}{
		{
			name: "Error not found",
			mockBehaviour: func(ctx context.Context, s *mock_cars.MockCarUseCase) {
				s.EXPECT().DeleteCar(uint32(1)).Return(&errors.ValidationError{})
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"errors":"internal server error"}`,
		},
		{
			name: "Success",
			mockBehaviour: func(ctx context.Context, s *mock_cars.MockCarUseCase) {
				s.EXPECT().DeleteCar(uint32(1)).Return(nil)
			},
			expectedCode: http.StatusNoContent,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCar := mock_cars.NewMockCarUseCase(ctrl)
			tt.mockBehaviour(ctx, mockCar)

			lgr := slog.New(slog.NewJSONHandler(os.Stderr, nil))
			handler := NewCarHandler(mockCar, lgr)
			w := httptest.NewRecorder()

			req := httptest.NewRequest(http.MethodDelete, "/api/cars/1", nil)
			router := mux.NewRouter()

			router.HandleFunc("/api/cars/{id}", handler.DeleteCar).Methods(http.MethodDelete)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Equal(t, tt.expectedBody, w.Body.String())
		})
	}
}
