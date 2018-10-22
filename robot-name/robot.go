package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Robot may not injure a human being or, through inaction, allow a human being to come to harm
type Robot struct {
	name string
}

// Name returns the name of a robot
func (r *Robot) Name() string {
	if r.name == "" {
		r.name = generateName()
	}
	return r.name
}

// Reset wipes out robot's name
func (r *Robot) Reset() {
	*r = Robot{}
}

func generateName() string {
	rand.Seed(time.Now().UTC().UnixNano())
	return fmt.Sprintf("%c%c%03d", 'A'+byte(rand.Intn(26)), 'A'+byte(rand.Intn(26)), rand.Intn(1000))
}
