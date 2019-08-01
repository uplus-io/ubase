module github.com/uplus-io/ubase

go 1.12

require (
	github.com/uplus-io/ucluster v0.0.0
	github.com/uplus-io/uengine v0.0.0
	github.com/uplus-io/ugo v0.0.0
)

replace (
	github.com/uplus-io/ucluster => ../ucluster
	github.com/uplus-io/uengine => ../uengine
	github.com/uplus-io/ugo => ../ugo
)
