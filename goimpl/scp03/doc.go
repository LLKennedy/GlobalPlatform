// Package scp03 implements the Secure Channel Protocol 03
package scp03

/*

Mutual Authentication

1. Application sends INITIALIZE UPDATE to the card with a random 'host' challenge
2. Card generates 'card' challenge
3. Card challenge + card cryptogram + SCP identifier (see config.go) is transmitted to application
4. Application re-derives card cryptogram from available information to confirm card identity
5. Application does the same basic thing to create host cryptogram, to pass back to the card on EXTERNAL AUTHENTICATE
6. Card can now re-derive host cryptogram to confirm application identity

*/ //
