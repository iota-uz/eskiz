package eskizapi_test

import (
	"context"
	"testing"

	eskizapiclient "github.com/iota-uz/eskiz/eskizapi"
	"github.com/stretchr/testify/require"
)

func NewTestApiClient(t *testing.T) (*eskizapiclient.APIClient, *EnvConfiguration) {
	t.Helper()
	env := LoadEnvConfiguration()

	cfg := eskizapiclient.NewConfiguration()
	cfg.Servers = eskizapiclient.ServerConfigurations{{URL: env.URL, Description: "env server"}}

	client := eskizapiclient.NewAPIClient(cfg)

	return client, env
}

func GetAccessedContext(t *testing.T, client *eskizapiclient.APIClient, env *EnvConfiguration) context.Context {
	t.Helper()

	result, resp, err := client.DefaultApi.
		Login(context.Background()).
		Email(env.Email).
		Password(env.Password).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)

	loginData := result.GetData()
	token := loginData.GetToken()

	ctx := context.WithValue(context.Background(), eskizapiclient.ContextAccessToken, token)

	return ctx
}
