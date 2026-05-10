## Creating a Module
When starting a Go project, you need to initialize a module with a unique path. This acts as your project's identifier.
The convention is to use a URL you own (like a GitHub path) to keep it globally unique and avoid conflicts.
Example:
```go 
go mod init snippetbox.silviasaavedra.net
```

Run this in your project's root directory. It creates a go.mod file that tracks your module name and dependencies.

- use this throughout the rest of the book 

Setting up your project as a module has a number of advantages — including making
it much easier to manage third-party dependencies, avoid supply-chain attacks, and ensure
reproducible builds of your application in the future.


## Module Paths — Do I Need to Host It?
No. The module path is just a unique name for your project. It doesn't need to exist online.
Using github.com/your-username/project-name is just a naming convention. If you push it to GitHub later, it'll already match. If you don't, everything still works locally.
Only matters if you're publishing a reusable package/library for other people to import — then the path should match the actual download URL. For personal projects, just pick a unique name and move on.

`handlers` (controllers) - theyre responsible for executing your application logic and for writting HTTP response headers and bodies 


router (servemux) - this stores a mapping between the url patterns for your application and the coressponding handlers. 

Web Server: Go has a built-in web server — you don't need Nginx or Apache. Your app listens for requests on its own.


So a basic Go web app is just three things: 
- handlers (your logic)
- a router (which URL goes to which handler)
- a built-in server (listens for requests).


`http.ResponseWrite` paramater provides methods for assembling a HTTP response and sending it to the user 

`http.Request` parameter is a pointer to a struct which holds information about the current request (like the HTTP method and the URL being request)

## How a Request Flows
When your Go app runs, it starts a web server on a port (e.g. localhost:4000). When a request comes in:

Server receives the HTTP request
ServeMux (router) checks the URL path
Handler that matches the path runs and sends back a response

So: Request → Server → Router → Handler → Response


You (the browser) are the one sending the request. When you type localhost:4000 in your browser, that's the request.
Then:

Server is just sitting there listening, waiting for someone to visit
You visit a URL in your browser → that sends an HTTP request to the server
Router looks at the URL path and picks the right handler
Handler runs your code and sends back a response (like HTML) to the browser

So the handler doesn't trigger anything — it responds. Think of it like a restaurant: the server (web server) is the front door, the host (router) checks your reservation and seats you at the right table, and the chef (handler) makes your food and sends it out.



`servemux` treats the URL patter `/` like a catch-all.


## Using go run 
During development the go run command is a convenient way to try out your code. It’s
essentially a shortcut that compiles your code, creates an executable binary in your /tmp
directory, and then runs this binary in one step.

## Routing Request 

## Fixed Paths vs Subtree Paths
Go's router (servemux) has two types of URL patterns:
Fixed paths — don't end with /. Only match the exact URL.

- /snippet/view only matches /snippet/view
- /snippet/create only matches /snippet/create

Subtree paths — end with `/`. Match anything that starts with that path. Think of them like having a wildcard /** at the end.

/ matches everything (that's why it acts as a catch-all)
/static/ would match /static/css/main.css, /static/js/app.js, etc.

Quick rule: trailing slash = wildcard, no trailing slash = exact match.


## The DefaultServeMux

`http.Handle()`/`http.HandleFunc()`- these allow you to register routes without declaring a servemux

### What happens behind the scene? 
these functions register their routes with something called the
DefaultServeMux.

`DefaultServeMux` is a global variable, any package can access it and register a route - including any third-party package that the application imports. 

If compromised -> DefaultServeMux can be used to do malicious handler to the web

```go
var DefaultServeMux = NewServeMux
```

***NOT RECOMMENDED FOR PRODUCTION APPLICATIONS***

## Additional Information 
### Servemux features and quirks 
- In Go servemux, longer URL patterns have priority over shorter ones.
- Pattern registration order does not matter.
- URL paths are automatically cleaned:
  - Removes `.` and `..`
  - Removes repeated slashes
  - Redirects to clean URL with `301 Permanent Redirect`
- Subtree paths with missing trailing `/` are automatically redirected.
  - Example: `/foo` → `/foo/` if `/foo/` is registered.

## What about RESTful routing?
- go does not support routing based on request methods 
- does not support clean urls with variables
- does not support regexp-based patterns


## Customizing HTTP headers


| MEHTOD  |PATTERN | HANDLER| ACTION|
| -------- | -------- | -------- |-------- |
|ANY | / | home | Display the home page|
| ANY | /snippet/view| snippetView | Display a specific snippet |
|POST | /snippet/create| snippetCreate| Create a new snippet|





## URL query strings 
