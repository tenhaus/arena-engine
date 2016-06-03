package bus

import (
  "fmt"
  "testing"
)

func TestCreateRoutingTopicForHandle(t *testing.T) {
  handle := "testytesterson1134"
  topic, err := CreateRoutingTopicForHandle(handle)
  expectedTopic := fmt.Sprintf("%s-routing", handle)

  if err != nil {
    t.Error(err)
  }

  if topic != expectedTopic {
    t.Errorf("Created topic |%s| doesn't match |%s|", topic, expectedTopic)
  }

  if err := DeleteTopic(topic); err != nil {
    t.Error(err)
  }
}
