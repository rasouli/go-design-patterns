package prototype

import (
	"errors"
	"fmt"
)

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

func GetShirtCloner() ShirtCloner {
	return ShirtCache{}
}

type ShirtCache struct{}

func (sc ShirtCache) GetClone(s int) (ItemInfoGetter, error) {
	switch s {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New(fmt.Sprintf("not supported color %s", s))
	}
}

type ItemInfoGetter interface {
	GetInfo() string
}

//////////////////////////////////////
type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("shirt with SKU %s  and color id %d that costs %f ", s.SKU,s.Color,s.Price)
}

func (i *Shirt) GetPrice() float32 {
	return i.Price
}

///////////////////////////////////////
var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}
