package routes

//todo: add routes for: python, haskell, c#, assembly (maybe), and other functional programming languages

var RouteCompiler = map[string]func(src, dist, path string, isDir bool){}

func addCompiler(name string, cb func(src, dist, path string, isDir bool)) {
	RouteCompiler[name] = cb
}
