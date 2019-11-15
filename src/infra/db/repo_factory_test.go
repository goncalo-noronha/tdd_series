package db

import (
	"github.com/goncalo-noronha/tdd_series/src/app"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppointmentRepoFactory(t *testing.T)  {
	
	t.Run("Test create appointment", func(t *testing.T) {
		assert.IsType(t, app.AppointmentRepository{}, AppointmentRepoFactory{}.Make(new(MockDynamoConnection)))
	})
}