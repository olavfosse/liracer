# liracer
The free and open source type racing game.

## Rewrite
liracer is in the process of being rewritten from scratch. The playable version at https://play.liracer.org uses the previous codebase written in node.js and React.

## Build frontend
The frontend is written in browser runable code, so no building is required for it.

## Build backend
Simply run `go build`. This produces a executable `play.liracer.org`. The executable is entirely self contained, it is statically linked and embeds the frontend
