package rssReader

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// Create a test HTTP server to return a predefined RSS feed
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rssFeed := `
<rss>
<channel>
  <title>W3Schools Home Page</title>
  <link>https://www.w3schools.com</link>
  <description>Free web building tutorials</description>
  <item>
    <title>RSS Tutorial</title>
    <link>https://www.w3schools.com/xml/xml_rss.asp</link>
    <description>New RSS tutorial on W3Schools</description>
	<source url="https://www.w3schools.com">W3Schools.com</source>
    <pubDate>Thu, 27 Apr 2006 09:00:00 +0000</pubDate>
  </item>
  <item>
	<title>Item 2</title>
	<link>https://example.com/item2</link>
  </item>
</channel>

</rss>
		`
		w.Write([]byte(rssFeed))
	}))
	defer ts.Close()

	// Call the Parse function with the test server URL
	items, err := Parse([]string{ts.URL})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// Check that the correct number of items were returned
	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}

	// Check the values of the first item
	item1 := items[0]

	assert.Equal(t, "RSS Tutorial", item1.Title)
	assert.Equal(t, "https://www.w3schools.com/xml/xml_rss.asp", item1.Link)
	assert.Equal(t, "W3Schools.com", item1.Source)
	assert.Equal(t, "https://www.w3schools.com", item1.SourceUrl)
	assert.Equal(t, "New RSS tutorial on W3Schools", item1.Description)

	pubTime, err := time.Parse("2006-01-02 15:04:05 -0700", "2006-04-27 09:00:00 +0000")
	assert.NoError(t, err)
	assert.Equal(t, pubTime, item1.PublishDate)

}
