// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package unmarshalertest

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/plog"
)

func TestNewNopLogs(t *testing.T) {
	unmarshaler := NewNopLogs()
	got, err := unmarshaler.UnmarshalLogs(nil)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, typeStr, unmarshaler.Type())
}

func TestNewWithLogs(t *testing.T) {
	logs := plog.NewLogs()
	logs.ResourceLogs().AppendEmpty()
	unmarshaler := NewWithLogs(logs)
	got, err := unmarshaler.UnmarshalLogs(nil)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, logs, got)
	require.Equal(t, typeStr, unmarshaler.Type())
}

func TestNewErrLogs(t *testing.T) {
	wantErr := errors.New("test error")
	unmarshaler := NewErrLogs(wantErr)
	got, err := unmarshaler.UnmarshalLogs(nil)
	require.Error(t, err)
	require.Equal(t, wantErr, err)
	require.NotNil(t, got)
	require.Equal(t, typeStr, unmarshaler.Type())
}
