package db

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type DAOMock struct {
	DAO
	mock.Mock
}

func (em *DAOMock) Read(item map[string]interface{}) (interface{}, error) {
	args := em.Called(item)
	if args.Error(1) != nil {
		return 0, args.Error(1)
	}
	return args.Get(0), args.Error(1)
}

func TestAppointRepository(t *testing.T) {

	t.Run("Test find one", func(t *testing.T) {
		daoMock := new(DAOMock)

		uuid := "uuid"
		input := map[string]interface{}{
			"id": uuid,
		}

		daoMock.On("Read", input).Return(Appointment{ID: uuid, Specialty: "Bucanons"}, nil)

		sut := AppointmentRepository{daoMock}

		appointment, _ := sut.FindOneByID("uuid")

		assert.Equal(t, Appointment{ID: "uuid", Specialty: "Bucanons"}, appointment)
	})
}
