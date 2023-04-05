package alert

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParsePayload(t *testing.T) {
	data := []byte(`{"id":2480431,"last_value":"25","updated_at":"2023-04-05 05:37:31 UTC","key":"locker1-temperature","data":{"created_at":"2023-04-05T05:37:31.607Z","value":"25","location":null,"id":"0F94FK7HK7EXXE91Q1AF6PT1PJ"}}`)

	_, err := ParsePayload(data)

	require.NoError(t, err)

}
