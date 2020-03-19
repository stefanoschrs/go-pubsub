package pubsub

import (
	"bytes"
	"errors"
	"testing"
)

func TestCreate(t *testing.T) {
	id := "123"

	t.Run("valid", func(t *testing.T) {
		s, err := Create(id)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		s.Close()
	})

	t.Run("exists", func(t *testing.T) {
		s, err := Create(id)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		_, err = Create(id)
		if !errors.Is(err, ErrExists) {
			t.Error("something went wrong")
			t.FailNow()
		}

		s.Close()
	})
}

func TestGet(t *testing.T) {
	id := "123"
	defaultData := []byte("Hello World")

	t.Run("not found", func(t *testing.T) {
		_, err := Get(id)
		if !errors.Is(err, ErrNotFound) {
			t.Error("something went wrong")
			t.FailNow()
		}
	})

	t.Run("valid", func(t *testing.T) {
		s, err := Create(id)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		s2, err := Get(id)
		if err != nil {
			t.Error(err)
			t.Fail()
		}

		if s.channel != s2.channel {
			t.Error("mismatch")
			t.Fail()
		}

		s.Close()
	})

	t.Run("publish", func(t *testing.T) {
		if _, err := Create(id); err != nil {
			t.Error(err)
			t.FailNow()
		}

		s, err := Get(id)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		go func() {
			data := s.Sub()

			if bytes.Compare(defaultData, data) != 0 {
				t.Error("bytes mismatch")
				t.FailNow()
			}
		}()

		s.Pub(defaultData)
		s.Close()
	})
}
