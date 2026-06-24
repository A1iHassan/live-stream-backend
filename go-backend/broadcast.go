package main

import "github.com/pion/webrtc/v3"

func (r *Room) BroadcastTrack(remoteTrack *webrtc.TrackRemote, receiver *Participant) {
	// Create a local track that behaves as the master source for subscribers, I need to look at what each of these passed attributes actually means
	localTrack, err := webrtc.NewTrackLocalStaticRTP(
		remoteTrack.Codec().RTPCodecCapability,
		remoteTrack.ID(),
		remoteTrack.StreamID(),
	)

	if err != nil {
		return
	}
	// Adding the track to the Room
	r.Lock()
	r.Tracks[remoteTrack.ID()] = localTrack
	r.Unlock()
	// Push the new track to existing participants
	r.RLock()
	for _, p := range r.Participants {
		if p.I
}
