# liracer
The free and open source type racing game.

## Rewrite
liracer is in the process of being rewritten from scratch. The playable version at https://play.liracer.org uses the previous codebase written in node.js and React.

NB: The TLS certificate of play.liracer.org is expired. A new certificate will be made when deploying the rewrite. 

## Build frontend
The frontend is written in browser runable code, so no building is required for it.

## Build backend
Simply run `go build`. This produces a executable `play.liracer.org`. The executable is entirely self contained, it is statically linked and embeds the frontend

## License
JetBrains Mono(Copyright The JetBrains Mono Project Authors) is licensed under the OFL1-1 license. The rest of the project, unless otherwise specified, is licensed under the [AGPLv3](https://www.gnu.org/licenses/agpl-3.0.html) license.
