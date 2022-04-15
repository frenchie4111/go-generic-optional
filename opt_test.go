package opt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSmoke(t *testing.T) {
	require.True(t, true)
}

func TestNew(t *testing.T) {
	anOptional := New[string]()
	_, ok := anOptional.Get()
	require.Equal(t, ok, false)
	require.False(t, anOptional.Exists())
	require.Equal(t, "something", anOptional.GetOrElse("something"))
}

func TestOf(t *testing.T) {
	anOptional := Of("hello")
	require.True(t, anOptional.Exists())
	value, ok := anOptional.Get()
	require.Equal(t, true, ok)
	require.Equal(t, "hello", value)
	require.Equal(t, "hello", anOptional.MustGet())
}

func TestMustGet(t *testing.T) {
	anOptional := New[string]()

	defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic")
        }
    }()

	anOptional.MustGet()
}

func TestIf(t *testing.T) {
	anOptional := New[string]()

	ranIf := false

	If(anOptional, func (value string) bool {
		ranIf = true
		return true
	})

	require.False(t, ranIf)

	anOptional = Of("hello")

	If(anOptional, func (value string) bool {
		ranIf = true
		return true
	})

	require.True(t, ranIf)
}

func TestMarshall(t *testing.T) {
	anOptional := New[string]()

	data, err := anOptional.MarshalJSON()

	require.Nil(t, err)
	require.Equal(t, `null`, string(data))

	anOptional = Of("hello")

	data, err = anOptional.MarshalJSON()

	require.Nil(t, err)
	require.Equal(t, `"hello"`, string(data))
}

func TestUnmarshall(t *testing.T) {
	anOptional := New[string]()

	err := anOptional.UnmarshalJSON([]byte(`null`))
	require.Nil(t, err)
	require.False(t, anOptional.Exists())

	anOptional = New[string]()
	err = anOptional.UnmarshalJSON([]byte(`"hello"`))
	require.Nil(t, err)
	require.True(t, anOptional.Exists())
	require.Equal(t, "hello", anOptional.MustGet())

	err = anOptional.UnmarshalJSON([]byte(`"asdjl:1k2l;j'""`))
	require.NotNil(t, err)
}
