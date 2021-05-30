/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package filter

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	overlay "github.com/docbull/inlab-fabric-modules/inlab-sdn/protos/overlay"
	"github.com/hyperledger/fabric/gossip/comm"
	"github.com/hyperledger/fabric/gossip/discovery"
	"github.com/hyperledger/fabric/gossip/util"
	"google.golang.org/grpc"
)

func init() { // do we really need this?
	rand.Seed(int64(util.RandomUInt64()))
}

// RoutingFilter defines a predicate on a NetworkMember
// It is used to assert whether a given NetworkMember should be
// selected for be given a message
type RoutingFilter func(discovery.NetworkMember) bool

// SelectNonePolicy selects an empty set of members
var SelectNonePolicy = func(discovery.NetworkMember) bool {
	return false
}

// SelectAllPolicy selects all members given
var SelectAllPolicy = func(discovery.NetworkMember) bool {
	return true
}

// CombineRoutingFilters returns the logical AND of given routing filters
func CombineRoutingFilters(filters ...RoutingFilter) RoutingFilter {
	return func(member discovery.NetworkMember) bool {
		for _, filter := range filters {
			if !filter(member) {
				return false
			}
		}
		return true
	}
}

// SelectPeers returns a slice of at most k peers randomly chosen from peerPool that match routingFilter filter.
func SelectPeers(k int, peerPool []discovery.NetworkMember, filter RoutingFilter) []*comm.RemotePeer {
	var res []*comm.RemotePeer
	// Iterate over the possible candidates in random order
	for _, index := range rand.Perm(len(peerPool)) {
		// If we collected K peers, we can stop the iteration.
		if len(res) == k {
			break
		}
		peer := peerPool[index]
		// For each one, check if it is a worthy candidate to be selected
		if filter(peer) {
			p := &comm.RemotePeer{PKIID: peer.PKIid, Endpoint: peer.PreferredEndpoint()}
			res = append(res, p)
		}
	}
	return res
}

// SelectPeers2 returns peer information that is included its subnode.
func SelectPeers2(peerName string, peerPool []discovery.NetworkMember, filter RoutingFilter) []*comm.RemotePeer {
	var res []*comm.RemotePeer

	for index := 0; index < len(peerPool); index++ {
		peer := peerPool[index]

		if peerName == "peer0.org1.example.com:7051" {
			if peer.Endpoint == "peer1.org1.example.com:7051" {
				if filter(peer) {
					p := &comm.RemotePeer{PKIID: peer.PKIid, Endpoint: peer.PreferredEndpoint()}
					res = append(res, p)
				}
			}
		}
	}

	fmt.Println("-------Selected Peers-------")
	fmt.Println(time.Now())
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i].Endpoint)
	}
	fmt.Println("----------------------------")

	return res
}

// GetOverlay connects to SDN Client to get Overlay Structure
func GetOverlay(peerName string) *overlay.OverlayStructure {
	hostIP := os.Getenv("IP_ADDRESS")
	conn, err := grpc.Dial(hostIP+":8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	peerClient := overlay.NewOverlayStructureServiceClient(conn)

	endpoint := overlay.PeerEndpoint{Endpoint: peerName}
	peerOverlayStructure, err := peerClient.SeeOverlayStructure(context.Background(), &endpoint)
	if err != nil {
		fmt.Println(err)
	}

	return peerOverlayStructure
}

// SelectPeers2 returns peer information that is included its subnode.
func SelectPeersWithinOverlayStructure(peerName string, peerPool []discovery.NetworkMember, filter RoutingFilter) ([]*comm.RemotePeer, bool) {
	var res []*comm.RemotePeer
	var isD2D bool

	mySubPeer := GetOverlay(peerName)
	isD2D = false

	for index := 0; index < len(peerPool); index++ {
		peer := peerPool[index]

		for sidx := 0; sidx < len(mySubPeer.SubPeers); sidx++ {
			if strings.Contains(mySubPeer.SubPeers[sidx], peer.Endpoint) {
				if filter(peer) {
					p := &comm.RemotePeer{PKIID: peer.PKIid, Endpoint: peer.PreferredEndpoint()}
					res = append(res, p)
				}
			}
		}
	}

	if len(res) > 0 && peerName != "peer0.org1.example.com:7051" {
		isD2D = true
	}

	fmt.Println("-------Selected Peers-------")
	fmt.Println(time.Now())
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i].Endpoint)
	}
	fmt.Println("----------------------------")

	return res, isD2D
}

// First returns the first peer that matches the given filter
func First(peerPool []discovery.NetworkMember, filter RoutingFilter) *comm.RemotePeer {
	for _, p := range peerPool {
		if filter(p) {
			return &comm.RemotePeer{PKIID: p.PKIid, Endpoint: p.PreferredEndpoint()}
		}
	}
	return nil
}

// AnyMatch filters out peers that don't match any of the given filters
func AnyMatch(peerPool []discovery.NetworkMember, filters ...RoutingFilter) []discovery.NetworkMember {
	var res []discovery.NetworkMember
	for _, peer := range peerPool {
		for _, matches := range filters {
			if matches(peer) {
				res = append(res, peer)
				break
			}
		}
	}
	return res
}
