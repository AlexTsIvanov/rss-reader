package main

func main() {

	s := NewRssService()

	s.POST("/", s.handleRssFeed)

	s.Logger.Fatal(s.Start(":6060"))

}
