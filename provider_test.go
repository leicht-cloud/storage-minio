package main

import (
	"testing"

	"github.com/leicht-cloud/leicht-cloud/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	provider, err := storage.SetupTestEnv(&StorageProvider{}, t.TempDir())
	if err != nil {
		t.Skip(err)
	}
	defer func() {
		err = storage.TeardownTestEnv()
		assert.NoError(t, err)
	}()

	storage.TestStorageProvider(provider, t)
}

func BenchmarkProvider(b *testing.B) {
	provider, err := storage.SetupTestEnv(&StorageProvider{}, b.TempDir())
	if err != nil {
		b.Skip(err)
	}
	defer func() {
		err = storage.TeardownTestEnv()
		assert.NoError(b, err)
	}()

	storage.BenchmarkStorageProvider(provider, b)
}
