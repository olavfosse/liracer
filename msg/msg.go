// msg provides the types used to marshal/unmarshal JSON-encoded WebSocket
// messages.
//
// This serves as a WebSocket API interface definition.
//
// The package is split into two files incoming.go and outgoing.go. incoming.go
// defines the type IncomingMsg which is used to unmarshal messages going from
// the clients to the server. The other file, outgoing.go, defines the type
// OutgoingMsg which is used to marshal messages going from the server to the
// clients.
package msg
