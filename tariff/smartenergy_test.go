package tariff

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSmartenergy(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Vienna")
	time.Local = loc
	at, err := NewSmartengeryFromConfig(map[string]interface{}{})
	require.NoError(t, err)
	tf := at.(*Smartengery)
	assert.NotNil(t, tf.data)
	val, _ := tf.data.Get()
	assert.NotNil(t, val)
	assert.GreaterOrEqual(t, len(val), 96, "expected 96 entries 24*4 (15 minute interval)")
	firstRate := val[0]

	assert.NotNil(t, firstRate.Start)
	fmt.Printf("%v\n", firstRate.Start)
	assert.Equal(t, firstRate.Start.Day(), time.Now().Day())
	assert.Equal(t, firstRate.Start.Month(), time.Now().Month())
	assert.Equal(t, firstRate.Start.Year(), time.Now().Year())
	assert.Equal(t, firstRate.Start.Hour(), 0)
	assert.Equal(t, firstRate.Start.Minute(), 0)
	assert.Equal(t, firstRate.Start.Second(), 0)

	assert.NotNil(t, firstRate.End)
	fmt.Printf("%v\n", firstRate.End)
	assert.Equal(t, firstRate.End.Day(), time.Now().Day())
	assert.Equal(t, firstRate.End.Month(), time.Now().Month())
	assert.Equal(t, firstRate.End.Year(), time.Now().Year())
	assert.Equal(t, firstRate.End.Hour(), 0)
	assert.Equal(t, firstRate.End.Minute(), 14)
	assert.Equal(t, firstRate.End.Second(), 59)
	assert.NotNil(t, firstRate.Price)
	assert.GreaterOrEqual(t, firstRate.Price, 0.0, "price is not greater or equal to 0.0")
}
