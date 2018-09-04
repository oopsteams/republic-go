package grpc

import (
	"log"
	"math/rand"
	"time"

	"github.com/republicprotocol/republic-go/shamir"

	"github.com/republicprotocol/republic-go/identity"
	"github.com/republicprotocol/republic-go/smpc"
	"github.com/republicprotocol/republic-go/testutils"
)

// RedNodeBehaviour indicates the malicious behaviours the
// red node will exhibit.
type RedNodeBehaviour int

// Values for a RedNodeBehaviour
const (
	InvalidRequests RedNodeBehaviour = iota
	InvalidNonce
	InvalidSignature
	InvalidBlindings
	InvalidJoins
	DropMessages
	DropMultiAddresses
	DropSignatures
)

// String returns a human-readable representation of RedNodeTypes.
func (behaviours RedNodeBehaviour) String() string {
	switch behaviours {
	case InvalidRequests:
		return "invalid requests"
	case InvalidNonce:
		return "invalid nonce"
	case InvalidSignature:
		return "invalid multi-address signature"
	case DropMultiAddresses:
		return "drop multi-addresses"
	case DropSignatures:
		return "drop multi-address signatures"
	default:
		return "unexpected behaviour"
	}
}

// RedNodeSwarmerTypes contains an array of all possible malicious swarming
// behaviours.
var RedNodeSwarmerTypes = []RedNodeBehaviour{
	InvalidRequests,
	InvalidNonce,
	InvalidSignature,
	DropMultiAddresses,
	DropSignatures,
}

// RedNodeStreamerTypes contains an array of all possible malicious streaming
// behaviours.
var RedNodeStreamerTypes = []RedNodeBehaviour{
	InvalidRequests,
	InvalidBlindings,
	InvalidJoins,
	DropMessages,
}

func getTamperedMultiAddress(multiAddr identity.MultiAddress) MultiAddress {
	redNodeType := RedNodeSwarmerTypes[rand.Intn(len(RedNodeSwarmerTypes))]

	rand.Seed(time.Now().UnixNano())
	multiAddress := MultiAddress{
		Signature:         multiAddr.Signature,
		MultiAddress:      multiAddr.String(),
		MultiAddressNonce: multiAddr.Nonce,
	}

	switch redNodeType {
	case InvalidRequests:
		multiAddress.Signature = tamperSignature(multiAddr)
		multiAddress.MultiAddressNonce = tamperNonce(multiAddr)
		multiAddress.MultiAddress = tamperMultiAddress(multiAddr)
	case InvalidNonce:
		multiAddress.MultiAddressNonce = tamperNonce(multiAddr)
	case InvalidSignature:
		multiAddress.Signature = tamperSignature(multiAddr)
	case DropMultiAddresses:
		multiAddress.MultiAddress = ""
	case DropSignatures:
		multiAddress.Signature = []byte{}
	default:
	}

	log.Printf("Red-node swarmer will exhibit behaviour: %v\n", redNodeType)
	log.Printf("Red-node tampered multi-address %v to look like %v", multiAddr, multiAddress)
	return multiAddress
}

func getTamperedMessage(message smpc.Message) smpc.Message {
	rand.Seed(time.Now().UnixNano())

	redNodeType := RedNodeStreamerTypes[rand.Intn(len(RedNodeStreamerTypes))]

	switch redNodeType {
	case InvalidRequests:
		message = tamperMessage(message)
	case InvalidBlindings:

	case InvalidJoins:

	case DropMessages:

	default:
	}

	log.Printf("Red-node streamer will exhibit behaviour: %v\n", redNodeType)
	log.Printf("Red-node tampered with smpc message %v to look like %v", message, message.MessageJoin)

	return message
}

func tamperMessage(message smpc.Message) smpc.Message {

	r := rand.Intn(100)

	switch message.MessageType {
	case smpc.MessageTypeJoin:
		if r < 50 {
			message.MessageType = smpc.MessageTypeJoinResponse
		}

		message.MessageJoin.NetworkID = tamperNetworkID(message.MessageJoin.NetworkID)
		message.MessageJoin.Join = tamperJoin(message.MessageJoin.Join)
	case smpc.MessageTypeJoinResponse:
		if r < 50 {
			message.MessageType = smpc.MessageTypeJoin
		}
		message.MessageJoinResponse.NetworkID = tamperNetworkID(message.MessageJoinResponse.NetworkID)
		message.MessageJoinResponse.Join = tamperJoin(message.MessageJoinResponse.Join)
	default:
		message.MessageType = smpc.MessageType(15)
	}
	if r < 80 && r >= 50 {
		message.MessageType = smpc.MessageType(0)
	}

	return message
}

func tamperJoin(join smpc.Join) smpc.Join {
	r := rand.Intn(100)
	// Return an empty smpc.Join.
	if r < 10 {
		return smpc.Join{}
	}
	// Return an updated smpc.Join.
	if r < 90 {
		join.ID = tamperJoinID(join.ID)
		join.Index = tamperJoinIndex(join.Index)
		join.Shares = tamperShares(join.Shares)
	}
	return join
}

func tamperJoinID(joinID smpc.JoinID) smpc.JoinID {
	r := rand.Intn(100)
	// Return an empty spmc.JoinID.
	if r < 10 {
		return smpc.JoinID{}
	}
	// Return a random [33]byte array as JoinID.
	if r < 50 {
		return smpc.JoinID(testutils.Random33Bytes())
	}
	// Modify the joinID slightly.
	if r < 90 {
		index := rand.Intn(33)
		joinID[index] = byte(index)
		return joinID
	}
	return joinID
}

func tamperJoinIndex(joinIndex smpc.JoinIndex) smpc.JoinIndex {
	r := rand.Intn(100)
	// Return an 0.
	if r < 10 {
		return 0
	}
	// Return a random uint64 as JoinIndex.
	if r < 50 {
		return smpc.JoinIndex(rand.Intn(200))
	}
	// Modify the joinIndex slightly.
	if r < 70 {
		return joinIndex + 1
	}
	if r < 90 {
		return joinIndex - 1
	}
	return joinIndex
}

func tamperShares(shares shamir.Shares) shamir.Shares {
	// r := rand.Intn(100)
	// // Return an empty spmc.JoinID.
	// if r < 10 {
	// 	return smpc.JoinID{}
	// }
	// // Return a random [33]byte array as JoinID.
	// if r < 50 {
	// 	return smpc.JoinID(testutils.Random33Bytes())
	// }
	// // Modify the joinID slightly.
	// if r < 90 {
	// 	index := rand.Intn(33)
	// 	joinID[index] = byte(index)
	// 	return joinID
	// }
	// return joinID
	return shamir.Shares{}
}

func tamperNetworkID(networkID smpc.NetworkID) smpc.NetworkID {
	return networkID
}

func tamperSignature(multiAddr identity.MultiAddress) []byte {
	r := rand.Intn(100)
	if r < 50 {
		randBytes := testutils.Random64Bytes()
		return randBytes[:]
	}
	multiAddr.Signature[rand.Intn(64)] = byte(rand.Intn(100))
	return multiAddr.Signature
}

func tamperMultiAddress(multiAddr identity.MultiAddress) string {
	r := rand.Intn(100)
	if r < 75 {
		multiAddr, _ := testutils.RandomMultiAddress()
		return multiAddr.String()
	}
	return multiAddr.String()
}

func tamperNonce(multiAddr identity.MultiAddress) uint64 {
	r := rand.Intn(100)
	if r < 33 {
		return multiAddr.Nonce + uint64(r)
	}
	if r < 66 {
		return multiAddr.Nonce - uint64(r)
	}
	if r < 90 {
		return 0
	}
	return multiAddr.Nonce
}

func tamperMessageJoin(join smpc.Join) smpc.Join {
	r := rand.Intn(100)
	if r < 33 {

	}
	if r < 66 {
	}
	if r < 99 {
	}
	return smpc.Join{}
}
