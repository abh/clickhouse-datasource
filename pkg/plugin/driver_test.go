package plugin

import (
	"testing"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	t.Skip()
	clickhouse := Clickhouse{}
	t.Run("should not error when valid settings passed", func(t *testing.T) {
		_, err := clickhouse.Connect(backend.DataSourceInstanceSettings{JSONData: []byte(`{ "server": "localhost", "port": 8123 }`), DecryptedSecureJSONData: map[string]string{}})
		assert.Equal(t, nil, err)
	})
}

func TestConnectSecure(t *testing.T) {
	t.Skip()
	clickhouse := Clickhouse{}
	t.Run("should not error when valid settings passed", func(t *testing.T) {
		json := `{ "server": "server", "port": 9440, "username": "foo", "secure": true }`
		secure := map[string]string{}
		_, err := clickhouse.Connect(backend.DataSourceInstanceSettings{JSONData: []byte(json), DecryptedSecureJSONData: secure})
		assert.Equal(t, nil, err)
	})
}
