package gee

import (
	"log"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(c *Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
	*Group
	groups []*Group
}

type Group struct {
	Prefix string
	Middleware []HandlerFunc
	Parent *Group		  // Get what the parent's group is
	engine *Engine        // Makes us access the router
}

// New is the constructor of gee.Engine
func NewEngine() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.Group = &Group{engine: engine}
	engine.groups = []*Group{engine.Group}
	return engine
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (g *Group) NewGroup(prefix string) *Group {
	engine := g.engine
	group := &Group{
		Prefix:     prefix,
		Middleware: nil,
		Parent:     g,
		engine:     engine,
	}
	return group
}

func (g *Group) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := g.Prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	g.engine.router.addRoute(method,pattern,handler)
}

func (g *Group) Get(pattern string, handler HandlerFunc) {
	g.addRoute("GET", pattern, handler)
}

func (g *Group) Post(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}
// GET defines the method to add GET request
func (engine *Engine) Get(pattern string, handler HandlerFunc) {
	engine.router.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) Post(pattern string, handler HandlerFunc) {
	engine.router.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	engine.router.handle(c)
}