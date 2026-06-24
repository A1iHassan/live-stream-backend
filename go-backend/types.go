package main

import (
	"sync"
	"github.com/pion/webrtc/v3"
)


type Room struct {
	sync.RWMutex
	Name string
	Participants map[string]*Participants
	Tracks map[string]*webrtc.TrackLocalStaticRTP
}

type Participants struct {
	ID string
	PublisherConn *webrtc.PeerConnection
	SubscriberConn *webrtc.PeerConnection
	SubscriberTrackRefs map[string]*webrtc.RTPSender
}


