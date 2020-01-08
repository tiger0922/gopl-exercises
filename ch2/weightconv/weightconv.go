package weightconv

import "fmt"

type Pound float64
type Kilogram float64

func (lb Pound) String() string  { return fmt.Sprintf("%glb", lb) }
func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }
func LbToKg(lb Pound) Kilogram { return Kilogram(lb / 2.2046) }
func KgToLb(kg Kilogram) Pound { return Pound(kg * 2.2046) }
