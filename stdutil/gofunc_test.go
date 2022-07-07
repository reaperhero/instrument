package stdutil_test

import (
	"testing"

	"github.com/reaperhero/instrument/dump"
	"github.com/reaperhero/instrument/stdutil"
	"github.com/stretchr/testify/assert"
)

func TestFuncName(t *testing.T) {
	name := stdutil.FuncName(stdutil.PkgName)
	assert.Equal(t, "github.com/reaperhero/instrument/stdutil.PkgName", name)

	name = stdutil.FuncName(stdutil.PanicIfErr)
	assert.Equal(t, "github.com/reaperhero/instrument/stdutil.PanicIfErr", name)
}

func TestPkgName(t *testing.T) {
	name := stdutil.PkgName(stdutil.FuncName(stdutil.PanicIfErr))
	assert.Equal(t, "github.com/reaperhero/instrument/stdutil", name)
}

func TestFullFcName_Parse(t *testing.T) {
	fullName := stdutil.FuncName(stdutil.PanicIfErr)

	ffn := stdutil.FullFcName{FullName: fullName}
	ffn.Parse()
	assert.Equal(t, fullName, ffn.String())
	assert.Equal(t, "stdutil", ffn.PkgName())
	assert.Equal(t, "PanicIfErr", ffn.FuncName())
	assert.Equal(t, "github.com/reaperhero/instrument/stdutil", ffn.PkgPath())
	dump.P(ffn)

	st := stdutil.FullFcName{}
	fullName = stdutil.FuncName(st.FuncName)

	ffn = stdutil.FullFcName{FullName: fullName}
	ffn.Parse()
	assert.Equal(t, "(*FullFcName).FuncName-fm", ffn.FuncName())
	dump.P(ffn)
}

func TestCutFuncName(t *testing.T) {
	fullName := stdutil.FuncName(stdutil.PanicIfErr)

	pkgPath, funcName := stdutil.CutFuncName(fullName)
	assert.Equal(t, "PanicIfErr", funcName)
	assert.Equal(t, "github.com/reaperhero/instrument/stdutil", pkgPath)
}
