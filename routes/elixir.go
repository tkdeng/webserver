package routes

func init() {
	addCompiler("exs", func(src, dist, path string, isDir bool) {
		//todo: create exs route in dist
	})
}
