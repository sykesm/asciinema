package asciicast

import "time"

type Stream struct {
	Frames        []Frame
	elapsedTime   time.Duration
	lastWriteTime time.Time
	maxWait       time.Duration
}

func NewStream(maxWait float64) *Stream {
	return &Stream{
		lastWriteTime: time.Now(),
		maxWait:       time.Duration(maxWait) * time.Second,
	}
}

func (s *Stream) Write(p []byte) (int, error) {
	frame := Frame{}
	frame.Delay = s.incrementElapsedTime().Seconds()
	frame.Data = make([]byte, len(p))
	copy(frame.Data, p)
	s.Frames = append(s.Frames, frame)

	return len(p), nil
}

func (s *Stream) Close() {
	s.incrementElapsedTime()

	if string(s.Frames[len(s.Frames)-1].Data) == "exit\r\n" {
		s.Frames = s.Frames[:len(s.Frames)-1]
	}
}

func (s *Stream) Duration() time.Duration {
	return s.elapsedTime
}

func (s *Stream) incrementElapsedTime() time.Duration {
	now := time.Now()
	d := now.Sub(s.lastWriteTime)

	// Time can move backwards
	if d < 0 {
		d = time.Millisecond
	}
	if s.maxWait > 0 && d > s.maxWait {
		d = s.maxWait
	}

	s.elapsedTime += d
	s.lastWriteTime = now

	return d
}
