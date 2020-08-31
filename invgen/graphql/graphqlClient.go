package graphql

import (
	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
	"golang.org/x/net/context"
)

type GraphqlClient struct {
	client *graphql.Client
}

func NewClient(accessToken string, serverUrl string) *GraphqlClient {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	return &GraphqlClient{graphql.NewClient(serverUrl, httpClient)}
}

func (c GraphqlClient) Query(req interface{}, vars map[string]interface{}) error{
	return c.client.Query(context.Background(), &req, vars)
}

func (c GraphqlClient) Mutate(req interface{}, vars map[string]interface{}) error{
	return c.client.Mutate(context.Background(), &req, vars)
}

