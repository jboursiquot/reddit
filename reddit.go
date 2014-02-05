// Package reddit implements a basic client for the Reddit API.
package reddit

import (
  "fmt"
  "net/http"
  "encoding/json"
  "errors"
)

// Item describes a Reddit item.
type Item struct {
  Title     string
  URL       string
  Comments  int `json:"num_comments"`
}

type response struct {
  Data struct {
    Children []struct {
      Data Item
    }
  }
}

func (i Item) String() string {
  com := ""

  switch i.Comments {
  case 0:
    // nothing
  case 1:
    com = " (1 comment)"
  default:
    com = fmt.Sprintf(" (%d comments)", i.Comments)
  }

  return fmt.Sprintf("%s%s\n%s", i.Title, com, i.URL)
}


// Get fetches the most recent items in the subreddit.
func Get(reddit string) ([]Item, error) {
  url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)

  resp, err := http.Get(url)

  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    return nil, errors.New(resp.Status)
  }

  r := new(response)
  err = json.NewDecoder(resp.Body).Decode(r)

  if err != nil {
    return nil, err
  }

  items := make([]Item, len(r.Data.Children))
  for i, child := range r.Data.Children {
    items[i] = child.Data
  }

  return items, nil
}


