// Handle communication between clients and the pit
package bus

import "google.golang.org/cloud/pubsub"
import "golang.org/x/net/context"
import "golang.org/x/oauth2/google"
import "google.golang.org/cloud"
import "google.golang.org/cloud/storage"
import "os"
import "fmt"

func Authenticate(key string) string {
  context, err := cloudContext("wtf")
  if err != nil {
    fmt.Println("Error creating context", err)
    os.Exit(1)
  }

  pubsub.CreateTopic(context, "wtftopic")
  return key
}

func cloudContext(projectId string) (context.Context, error) {
  ctx := context.Background()
	httpClient, err := google.DefaultClient(ctx, storage.ScopeFullControl, pubsub.ScopePubSub)
	if err != nil {
		return nil, err
	}
	return cloud.WithContext(ctx, projectId, httpClient), nil
}
