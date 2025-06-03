# 🔥 FIRE 🔥

An experimental go project to allow real hot reloading in Go
Almost like [air](https://github.com/air-verse/air) but only restarts your application on serious
changes

## 👷🏿‍♂️🚧Still under construction 🚧👷🏿‍♂️ 

## Features it should support

- [ ] Hot reload on function body changing
- [ ] Restart if function signature changes
- [ ] Hot reload on method body changing
- [ ] Restart if method parameters changes
- [ ] Restart on struct type or parameters changing
- [ ] Restart on new method/function added that get's used
- [ ] Restart on method/function removed that was being used

## Basic rough sketch of how hot reloading should work

- Go is a very good language for backends, very simple to create microservices
that interact with each other, and what is a program most of the time than just
a monolith of what could easily be microservices?

- So why not split our codebases automatically into seperate executables,
whereby function call barriers are just an RPC call to another process, 
this would allow the other process to be updated (aka the function) while the actual main process
holding the state is still running so long as the main process doesn't change it's 
function contract with the other processes, eg: changing/adding/removing fields
from a struct or changing the function\method signatures, then the main process
won't need to restart, but only the sub processes that simulate a function and
method calls will be getting restarted when code changes.

- How do I plan to achieve this auto updating of the function/method processes,
if the actual user of this project for reloading codebase isn't separating their code
explicitly? Well this project would simply produce an executable, and this executables
work is to start your `main.go`, on the root with a simple `go run ./main.go`,
but first, we go through the project, split every project level function and method into it's own
individual mini-project, with their own `main.go` each, launch them, then launch your modified
`main.go`, and when you change anything on the codebase we update it on `shadow mirror project`
and restart the processes affected by your edit. These edits are only registered when a source
file is saved.

### Why not load .dll or .so, files like other C hotreloading libraries/frameworks do?
- For starters, Go has a runtime, and it bundles this runtime in every single executable or
shared object(library), loading more than 1 of these shared objects into a single process
already written in go is already troubling, due to possible collitions between the two runtimes
sharing the same process space, now imagine if we had 1000+ shared objects representing
individual functions, alot of very bad things can happen.

- But doesn't go have the plugin option when compiling? Yes it does, but the problem with this
option is a plugin cannot be unloaded once loaded, that means, if a function should change, then
we might end up loading 100000+ of the same function, which even if we find a way of mangling 
function names, we still stand the risk of bloating up memory, as alot of dead code would just be sitted
doing nothing. Ooh yeah one last thing, plugins aren't supported on Windows, atleast at the time of
this writing.

## Why are you doing this?
Go is already fast at compiling, but sometimes you don't want to restart your whole app
just to see a change as you'd loose current state, and sometimes that state can be important, that's basically it.

Plus also seems like a good challenge 😉
