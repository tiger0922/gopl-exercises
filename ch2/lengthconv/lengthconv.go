package lengthconv

import "fmt"

type Meter float64
type Feet float64

func (m Meter) String() string  { return fmt.Sprintf("%gm", m) }
func (ft Feet) String() string { return fmt.Sprintf("%gft", ft) }
func MToFt(m Meter) Feet { return Feet(m * 3.2808) }
func FtToM(ft Feet) Meter { return Meter(ft / 3.2808) }
