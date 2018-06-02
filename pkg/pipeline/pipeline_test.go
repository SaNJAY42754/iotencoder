package pipeline_test

import (
	"context"
	"encoding/base64"
	"errors"
	"testing"

	kitlog "github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	datastore "github.com/thingful/twirp-datastore-go"

	"github.com/thingful/iotencoder/pkg/mocks"
	"github.com/thingful/iotencoder/pkg/pipeline"
	"github.com/thingful/iotencoder/pkg/postgres"
)

func TestProcess(t *testing.T) {
	logger := kitlog.NewNopLogger()
	ds := mocks.Datastore{}

	payload := []byte("hello")
	encodedPayload := []byte(base64.StdEncoding.EncodeToString(payload))

	// set up two mock responses
	ds.On(
		"WriteData",
		context.Background(),
		&datastore.WriteRequest{
			PublicKey: "key1",
			UserUid:   "bob",
			Data:      encodedPayload,
		},
	).Return(
		&datastore.WriteResponse{},
		nil,
	)

	ds.On(
		"WriteData",
		context.Background(),
		&datastore.WriteRequest{
			PublicKey: "key2",
			UserUid:   "bob",
			Data:      encodedPayload,
		},
	).Return(
		&datastore.WriteResponse{},
		nil,
	)

	processor := pipeline.NewProcessor(datastore.Datastore(&ds), true, logger)

	device := &postgres.Device{
		UserUID: "bob",
		Streams: []*postgres.Stream{
			{
				PublicKey: "key1",
			},
			{
				PublicKey: "key2",
			},
		},
	}

	err := processor.Process(device, payload)
	assert.Nil(t, err)

	ds.AssertExpectations(t)
}

func TestProcessWithError(t *testing.T) {
	logger := kitlog.NewNopLogger()
	ds := mocks.Datastore{}

	payload := []byte("hello")
	encodedPayload := []byte(base64.StdEncoding.EncodeToString(payload))

	ds.On(
		"WriteData",
		context.Background(),
		&datastore.WriteRequest{
			PublicKey: "key1",
			UserUid:   "bob",
			Data:      encodedPayload,
		},
	).Return(
		&datastore.WriteResponse{},
		errors.New("error"),
	)

	processor := pipeline.NewProcessor(datastore.Datastore(&ds), true, logger)
	device := &postgres.Device{
		UserUID: "bob",
		Streams: []*postgres.Stream{
			{
				PublicKey: "key1",
			},
		},
	}

	err := processor.Process(device, payload)
	assert.NotNil(t, err)
	assert.Equal(t, "error", err.Error())

	ds.AssertExpectations(t)
}