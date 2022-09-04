package log

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateLogger(t *testing.T) {
	logger, err := New(io.Discard)

	require.NoError(t, err)
	require.NotNil(t, logger)
}

func TestLoggerJsonOutput(t *testing.T) {

	type cases = []struct {
		Name string
		Func func(msg string, fields ...Field)
	}

	generate := func(logger Logger) cases {
		return cases{
			{
				Name: "Debug",
				Func: logger.Debug,
			},
			{
				Name: "Info",
				Func: logger.Info,
			},
			{
				Name: "Warn",
				Func: logger.Warn,
			},
			{
				Name: "Error",
				Func: logger.Error,
			},
		}
	}

	var buf bytes.Buffer
	logger, err := New(io.Discard, &buf)
	require.NoError(t, err)
	require.NotNil(t, logger)

	t.Run("JSON", func(t *testing.T) {
		defer buf.Reset()

		l := logger
		for _, c := range generate(l) {
			t.Run(c.Name, func(t *testing.T) {
				buf.Reset()

				c.Func("test")
				require.NoError(t, logger.Sync())
				require.NoError(t, json.NewDecoder(&buf).Decode(&struct{}{}))
			})
		}
	})

	t.Run("WithField", func(t *testing.T) {
		defer buf.Reset()

		const value = "test"

		var actual = struct {
			With string `json:"with"`
		}{}
		l := logger.With(String("with", value))

		for _, c := range generate(l) {
			t.Run(c.Name, func(t *testing.T) {
				buf.Reset()

				c.Func("test")
				require.NoError(t, logger.Sync())
				require.NoError(t, json.NewDecoder(&buf).Decode(&actual))
			})
		}

	})

}
