# Vroom
Trying my hand at a lightweight static site generator based on the default [Go templating engine](http://golang.org/pkg/text/template/). Maybe something in the style of [Jekyll](http://jekyllrb.com).

#### VroomOptions
Configuration options are currently given in JSON form and read from a file. As of right now there are only two useful options: 

| Option Name   | Description   |
| ------------- | ------------- |
| TemplateDirectory      | Location of the templates to be rendered |
| BuildDirectory    | Destination of the rendered files      |
| Metadata stripes | This is where all metadata will live, right now this is unused but later can be used for titles, authors, url, names of plugins, and variables accessible from the templates.      |
