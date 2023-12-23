package data

type Post struct {
	Title       string
	Date        string
	Slug        string
	Description string
}

var Posts = []Post{
	{
		Title: "My approach to Nest.js",
		Date:  "23rd of December 2023",
		Slug:  "my-approach-to-nest-js",
		Description: "As I've been using Nest.js for a few years now I noticed I introduce several simple patterns in " +
			"all of my projects from the get go. Documenting them here, maybe you'll find them of interest.",
	},
	{
		Title: "My approach to Nest.js",
		Date:  "23rd of December 2023",
		Slug:  "my-approach-to-nest-js",
		Description: "As I've been using Nest.js for a few years now I noticed I introduce several simple patterns in " +
			"all of my projects from the get go. Documenting them here, maybe you'll find them of interest.",
	},
	{
		Title: "My approach to Nest.js",
		Date:  "23rd of December 2023",
		Slug:  "my-approach-to-nest-js",
		Description: "As I've been using Nest.js for a few years now I noticed I introduce several simple patterns in " +
			"all of my projects from the get go. Documenting them here, maybe you'll find them of interest.",
	},
}
