package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type Example interface {
	Run(r int) error
}

type MamboCreation struct {
	example Example
}

type ExampleMock struct {
	mock.Mock
}

func (em *ExampleMock) Run(r int) (int, error) {
	args := em.Called(r)
	if args.Error(1) != nil {
		return 0, args.Error(1)
	}
	return args.Int(0), args.Error(1)
}

func TestExample(t *testing.T) {

	t.Run("A test", func(t *testing.T) {
		mock := new(ExampleMock)
		mock.On("Run", 5).Return(2, nil)

		fmt.Println(mock.Run(5))

		//varible := MamboCreation{mock}

		assert.Equal(t, 1 ,1)
	})
}
