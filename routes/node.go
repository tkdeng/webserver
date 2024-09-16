package routes

func init() {
	addCompiler("js", func(src, dist, path string, isDir bool) {
		//todo: create js route in dist
		// note: for node.js, keep a directory in dist
		// may handle files completely differently
	})
}
