package types

import "fmt"

type Weight int32

func (w Weight) String() string {
	return fmt.Sprintf("%.1f", float64(w/10))
}
