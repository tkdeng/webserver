package routes

func init() {
	addCompiler("exs", func(root, path string, isDir bool) {
		//todo: create exs route in dist

		dir, file, out := getPaths(root, path, isDir, "exs")
		if out == "" {
			return
		}

		// fmt.Println(dir, file, out)
		_, _ = dir, file
	})
}
