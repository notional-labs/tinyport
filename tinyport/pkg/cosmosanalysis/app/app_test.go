package app_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/notional-labs/tinyport/tinyport/pkg/cosmosanalysis/app"
)

var (
	AppFile = []byte(`
package foo

type Foo struct {
	FooKeeper foo.keeper
}

func (f Foo) RegisterAPIRoutes() {}
func (f Foo) RegisterTxService() {}
func (f Foo) RegisterTendermintService() {}
`)

	NoAppFile = []byte(`
package foo

type Bar struct {
	FooKeeper foo.keeper
}
`)

	TwoAppFile = []byte(`
package foo

type Foo struct {
	FooKeeper foo.keeper
}

func (f Foo) RegisterAPIRoutes() {}
func (f Foo) RegisterTxService() {}
func (f Foo) RegisterTendermintService() {}

type Bar struct {
	FooKeeper foo.keeper
}

func (f Bar) RegisterAPIRoutes() {}
func (f Bar) RegisterTxService() {}
func (f Bar) RegisterTendermintService() {}
`)
)

func TestCheckKeeper(t *testing.T) {
	tmpDir := t.TempDir()

	// Test with a source file containing an app
	tmpFile := filepath.Join(tmpDir, "app.go")
	err := os.WriteFile(tmpFile, AppFile, 0644)
	require.NoError(t, err)

	err = app.CheckKeeper(tmpDir, "FooKeeper")
	require.NoError(t, err)
	err = app.CheckKeeper(tmpDir, "BarKeeper")
	require.Error(t, err)

	// No app in source must return an error
	tmpDirNoApp := t.TempDir()
	tmpFileNoApp := filepath.Join(tmpDirNoApp, "app.go")
	err = os.WriteFile(tmpFileNoApp, NoAppFile, 0644)
	require.NoError(t, err)
	err = app.CheckKeeper(tmpDirNoApp, "FooKeeper")
	require.Error(t, err)

	// More than one app must return an error
	tmpDirTwoApp := t.TempDir()
	tmpFileTwoApp := filepath.Join(tmpDirTwoApp, "app.go")
	err = os.WriteFile(tmpFileTwoApp, TwoAppFile, 0644)
	require.NoError(t, err)
	err = app.CheckKeeper(tmpDirTwoApp, "FooKeeper")
	require.Error(t, err)
}
