package rtc

type ParticipantParams struct {
}

type ParticipantImpl struct {
	publisher  *PCTransport
	subscriber *PCTransport
}

func NewParticipant(params ParticipantParams) (*ParticipantImpl, error) {
	p := &ParticipantImpl{}
}
