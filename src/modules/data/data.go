package data

type Post struct {
	Title       string
	Date        string
	Slug        string
	Description string
}

var Posts = []Post{
	{
		Title:       "Test post",
		Date:        "24rd of December 2023",
		Slug:        "test-post",
		Description: "I'll start to add posts soon. Here's a placeholder for now!",
	},
}
