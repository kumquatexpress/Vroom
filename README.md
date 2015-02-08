# Vroom
Trying my hand at a lightweight static site generator based on the default [Go templating engine](http://golang.org/pkg/text/template/). Maybe something in the style of [Jekyll](http://jekyllrb.com).

#### VroomOptions
Configuration options are currently given in JSON form and read from a file. As of right now there are only two useful options: 

| Option Name   | Description   | Default |
| ------------- | ------------- | ------- |
| LayoutDirectory      | Location of the template/snippets used | templates/ |
| BuildDirectory    | Destination of the rendered files     | build/ | 
| Pages Directory | Location of the pages to be rendered | _pages/ | 
| Metadata | This is where all top-level metadata will live, right now this is unused but later can be used for titles, authors, url, names of plugins, and variables accessible from the templates.      | {} |

#### Metadata
Vroom supports multiple layers of metadata in its templates. The more specific the level, the more priority the data will have when rendering. Each data is a key-value pair in a JSON file of the format *.vroom.json. Vroom will parse the metadata by directory. For instance, if the structure is

-	rootdir/
	-	\_pages/
	- data.vroom.json
		- blogposts/
			- blogdata.vroom.json
-	vroom_options.json

and the key _title_ exists in all of the json files, Vroom will create metadata maps specific to each directory --

    rootdir/_pages/blogposts/blog1.vroom.html 

will render the value of _{{ .title }}_ as the value in _blogdata.vroom.json_ since it's the most specific to that directory, whereas 

    rootdir/_pages/page1.vroom.html
    
will render _{{ .title }}_ as the value in _data.vroom.json_.

Top-level metadata keys contained in the     VroomOptions will be overridden by any *.vroom.json located in the pages directory.