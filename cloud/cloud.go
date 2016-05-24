package cloud

import "google.golang.org/cloud"
import "golang.org/x/net/context"
import "golang.org/x/oauth2/google"
import "google.golang.org/cloud/pubsub"
import "google.golang.org/cloud/storage"

func CloudContext(projectId string) (context.Context, error) {
  ctx := context.Background()
	httpClient, err := google.DefaultClient(
    ctx, storage.ScopeFullControl, pubsub.ScopePubSub)

	if err != nil {
		return nil, err
	}

	return cloud.WithContext(ctx, projectId, httpClient), nil
}
