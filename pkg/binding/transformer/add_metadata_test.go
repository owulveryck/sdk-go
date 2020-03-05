package transformer

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cloudevents/sdk-go/pkg/binding"
	"github.com/cloudevents/sdk-go/pkg/binding/spec"
	"github.com/cloudevents/sdk-go/pkg/binding/test"
	"github.com/cloudevents/sdk-go/pkg/types"
)

func TestAddAttribute(t *testing.T) {
	e := test.MinEvent()
	e.Context = e.Context.AsV1()

	subject := "aaa"
	expectedEventWithSubject := test.CopyEventContext(e)
	require.NoError(t, expectedEventWithSubject.Context.SetSubject(subject))

	timestamp, err := types.ToTime(time.Now())
	require.NoError(t, err)
	expectedEventWithTime := test.CopyEventContext(e)
	require.NoError(t, expectedEventWithTime.Context.SetTime(timestamp))

	test.RunTransformerTests(t, context.Background(), []test.TransformerTestArgs{
		{
			Name:         "No change to id to Mock Structured message",
			InputMessage: test.MustCreateMockStructuredMessage(test.CopyEventContext(e)),
			WantEvent:    test.CopyEventContext(e),
			Transformers: binding.TransformerFactories{AddAttribute(spec.ID, "new-id")},
		},
		{
			Name:         "No change to id to Mock Binary message",
			InputMessage: test.MustCreateMockBinaryMessage(test.CopyEventContext(e)),
			WantEvent:    test.CopyEventContext(e),
			Transformers: binding.TransformerFactories{AddAttribute(spec.ID, "new-id")},
		},
		{
			Name:         "No change to id to Event message",
			InputMessage: binding.EventMessage(test.CopyEventContext(e)),
			WantEvent:    test.CopyEventContext(e),
			Transformers: binding.TransformerFactories{AddAttribute(spec.ID, "new-id")},
		},
		{
			Name:         "Add subject to Mock Structured message",
			InputMessage: test.MustCreateMockStructuredMessage(test.CopyEventContext(e)),
			WantEvent:    expectedEventWithSubject,
			Transformers: binding.TransformerFactories{AddAttribute(spec.Subject, subject)},
		},
		{
			Name:         "Add subject to Mock Binary message",
			InputMessage: test.MustCreateMockBinaryMessage(test.CopyEventContext(e)),
			WantEvent:    expectedEventWithSubject,
			Transformers: binding.TransformerFactories{AddAttribute(spec.Subject, subject)},
		},
		{
			Name:         "Add subject to Event message",
			InputMessage: binding.EventMessage(test.CopyEventContext(e)),
			WantEvent:    expectedEventWithSubject,
			Transformers: binding.TransformerFactories{AddAttribute(spec.Subject, subject)},
		},
		{
			Name:         "Add time to Mock Structured message",
			InputMessage: test.MustCreateMockStructuredMessage(test.CopyEventContext(e)),
			WantEvent:    expectedEventWithTime,
			Transformers: binding.TransformerFactories{AddAttribute(spec.Time, timestamp)},
		},
		{
			Name:         "Add time to Mock Binary message",
			InputMessage: test.MustCreateMockBinaryMessage(test.CopyEventContext(e)),
			WantEvent:    expectedEventWithTime,
			Transformers: binding.TransformerFactories{AddAttribute(spec.Time, timestamp)},
		},
		{
			Name:         "Add time to Event message",
			InputMessage: binding.EventMessage(test.CopyEventContext(e)),
			WantEvent:    expectedEventWithTime,
			Transformers: binding.TransformerFactories{AddAttribute(spec.Time, timestamp)},
		},
	})
}

func TestAddExtension(t *testing.T) {
	e := test.MinEvent()
	e.Context = e.Context.AsV1()

	extName := "aaa"
	extValue := "bbb"
	expectedEventWithExtension := test.CopyEventContext(e)
	require.NoError(t, expectedEventWithExtension.Context.SetExtension(extName, extValue))

	test.RunTransformerTests(t, context.Background(), []test.TransformerTestArgs{
		{
			Name:         "No change to extension 'aaa' to Mock Structured message",
			InputMessage: test.MustCreateMockStructuredMessage(test.CopyEventContext(expectedEventWithExtension)),
			WantEvent:    test.CopyEventContext(expectedEventWithExtension),
			Transformers: binding.TransformerFactories{AddExtension(extName, extValue)},
		},
		{
			Name:         "No change to extension 'aaa' to Mock Binary message",
			InputMessage: test.MustCreateMockBinaryMessage(test.CopyEventContext(expectedEventWithExtension)),
			WantEvent:    test.CopyEventContext(expectedEventWithExtension),
			Transformers: binding.TransformerFactories{AddExtension(extName, extValue)},
		},
		{
			Name:         "No change to extension 'aaa' to Event message",
			InputMessage: binding.EventMessage(test.CopyEventContext(expectedEventWithExtension)),
			WantEvent:    test.CopyEventContext(expectedEventWithExtension),
			Transformers: binding.TransformerFactories{AddExtension(extName, extValue)},
		},
		{
			Name:         "Add extension 'aaa' to Mock Structured message",
			InputMessage: test.MustCreateMockStructuredMessage(test.CopyEventContext(e)),
			WantEvent:    test.CopyEventContext(expectedEventWithExtension),
			Transformers: binding.TransformerFactories{AddExtension(extName, extValue)},
		},
		{
			Name:         "Add extension 'aaa' to Mock Binary message",
			InputMessage: test.MustCreateMockBinaryMessage(test.CopyEventContext(e)),
			WantEvent:    test.CopyEventContext(expectedEventWithExtension),
			Transformers: binding.TransformerFactories{AddExtension(extName, extValue)},
		},
		{
			Name:         "Add extension 'aaa' to Event message",
			InputMessage: binding.EventMessage(test.CopyEventContext(e)),
			WantEvent:    test.CopyEventContext(expectedEventWithExtension),
			Transformers: binding.TransformerFactories{AddExtension(extName, extValue)},
		},
	})
}