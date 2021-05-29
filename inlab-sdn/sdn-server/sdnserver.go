package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	info "github.com/docbull/inlab-fabric-modules/inlab-sdn/protos/peer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type OverlayStructure struct {
	peerMembership []*info.PeerConnection
	singleOverlay  []*info.OverlayStructure
}

// Start starts the SDN Server for constructing overlay structure
func Start() {
	overlay := OverlayStructure{
		peerMembership: nil,
		singleOverlay:  nil,
	}

	go overlay.StartSDNListening()
	for {
		// Overlay Structure will be constructed every 5 seconds
		overlay.ConstructOverlayStructure()
		time.Sleep(time.Second * 5)
		go overlay.PrintOverlay()
	}
}

// StartSDNListening starts to listen SDN Client connections
func (overlay *OverlayStructure) StartSDNListening() {
	lis, err := net.Listen("tcp", ":9000")
	checkError(err)

	fmt.Println("SDN Server ...")
	grpcServer := grpc.NewServer()
	info.RegisterSDNServiceServer(grpcServer, overlay)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

// StorePeerInformation receive Peer information and store it
func (overlay *OverlayStructure) StorePeerInformation(ctx context.Context, receivedPeer *info.PeerConnection) (*info.Status, error) {
	if receivedPeer.Mobile.Endpoint == "" {
		fmt.Println("None of Endpoint")
		return &info.Status{Code: info.StatusCode_Failed}, nil
	}

	for i := 0; i < len(overlay.peerMembership); i++ {
		if overlay.peerMembership[i].Mobile.Endpoint == receivedPeer.Mobile.Endpoint {
			overlay.peerMembership[i].Mobile.NetInfo = receivedPeer.Mobile.NetInfo

			return &info.Status{Code: info.StatusCode_Ok}, nil
		}
	}

	newPeer := info.PeerConnection{
		Mobile: receivedPeer.Mobile,
		D2D:    receivedPeer.D2D,
	}

	overlay.peerMembership = append(overlay.peerMembership, &newPeer)

	go overlay.ConstructOverlayStructure()

	return &info.Status{Code: info.StatusCode_Ok}, nil
}

// ConstructOverlayStructure constructs overlay structure by considering peers' network states
func (overlay *OverlayStructure) ConstructOverlayStructure() {
	overlay.singleOverlay = nil

	for i := 0; i < len(overlay.peerMembership); i++ {
		overlay.singleOverlay = append(overlay.singleOverlay, &info.OverlayStructure{Endpoint: overlay.peerMembership[i].Mobile.Endpoint, SubPeers: nil})

		if overlay.peerMembership[i].D2D.Endpoint != "" {
			for j := 0; j < len(overlay.peerMembership); j++ {

				if strings.Contains(overlay.peerMembership[j].Mobile.Endpoint, overlay.peerMembership[i].D2D.Endpoint) {
					// calculating network costs for block receiving
					delayFromMEC := overlay.peerMembership[i].Mobile.NetInfo.NetworkStrength
					delayFromD2D := overlay.peerMembership[j].Mobile.NetInfo.NetworkStrength + overlay.peerMembership[j].D2D.NetInfo.NetworkStrength

					// if is receiving block from D2D faster than mobile network,
					// it appends D2D partner as its super peer.
					if delayFromMEC > delayFromD2D {
						//overlay.singleOverlay[0].SubPeers = append(overlay.singleOverlay[0].SubPeers, overlay.peerMembership[i].Mobile.Endpoint)
						overlay.singleOverlay[i].SuperPeer = overlay.peerMembership[j].Mobile.Endpoint
						overlay.singleOverlay[i].SubPeers = append(overlay.singleOverlay[i].SubPeers, "")
						break
					} else {
						overlay.singleOverlay[0].SubPeers = append(overlay.singleOverlay[0].SubPeers, overlay.peerMembership[i].Mobile.Endpoint)
						overlay.singleOverlay[i].SuperPeer = overlay.peerMembership[0].Mobile.Endpoint

						// calculating D2D pair's network delays
						D2DPairMobileDelay := overlay.peerMembership[j].Mobile.NetInfo.NetworkStrength
						D2DPairD2DDelay := overlay.peerMembership[i].Mobile.NetInfo.NetworkStrength + overlay.peerMembership[i].D2D.NetInfo.NetworkStrength

						// if is D2D partner's mobile network slow than D2D network,
						// appends partner as sub peer.
						if D2DPairMobileDelay > D2DPairD2DDelay {
							overlay.singleOverlay[i].SubPeers = append(overlay.singleOverlay[i].SubPeers, overlay.peerMembership[j].Mobile.Endpoint)
						} else {
							overlay.singleOverlay[i].SubPeers = append(overlay.singleOverlay[i].SubPeers, "")
						}
						break
					}
				}
			}
		} else {
			overlay.singleOverlay[0].SubPeers = append(overlay.singleOverlay[0].SubPeers, overlay.peerMembership[i].Mobile.Endpoint)
			overlay.singleOverlay[i].SuperPeer = overlay.singleOverlay[0].Endpoint
			overlay.singleOverlay[i].SubPeers = append(overlay.singleOverlay[i].SubPeers, "")
		}
	}
}

func (overlay *OverlayStructure) UpdateOverlayStructure(ctx context.Context, reqPeer *info.PeerInfo) (*info.OverlayStructure, error) {
	var peerIdx int32

	for i := 0; i < len(overlay.singleOverlay); i++ {
		if strings.Contains(overlay.singleOverlay[i].Endpoint, reqPeer.Endpoint) {
			peerIdx = int32(i)
			break
		}
	}

	return overlay.singleOverlay[peerIdx], nil
}

func (overlay *OverlayStructure) PrintOverlay() {
	fmt.Println(time.Now())
	fmt.Println("-------------SDN Client Member-------------")
	for i := 0; i < len(overlay.singleOverlay); i++ {
		fmt.Println(overlay.singleOverlay[i].Endpoint)
		fmt.Println(overlay.singleOverlay[i].SubPeers)
	}
	fmt.Println("-------------------------------------------")
	fmt.Println()
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func main() {
	Start()
}
