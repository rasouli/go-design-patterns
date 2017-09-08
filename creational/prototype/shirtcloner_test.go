package prototype

import "testing"

func TestClone(t *testing.T){

	cloner := GetShirtCloner()

	if cloner == nil {
		t.Fatal("received empty cache")
	}

	item1, err := cloner.GetClone(White)

	if err != nil {
		t.Errorf("white tshirt failed on cloning: %v", err)
	}

	if item1 ==  whitePrototype {
		t.Errorf("the cloned shirt must be diffrent from prototype")
	}

	shirt1 , ok := item1.(*Shirt)

	if !ok {
		t.Fatal("type assertation for item 1 cannot be done")
	}

	shirt1.SKU = "shirt1"

	item2, err := cloner.GetClone(White)

	if err != nil {
		t.Errorf("could not make second clone of white shirt %v", err)
	}

	if item2 == whitePrototype {
		t.Errorf("the second clone of white shirt is the same")
	}

	shirt2, ok := item2.(*Shirt)

	if !ok {
		t.Fatal("could not type assert the second item to *Shirt")
	}

	if shirt1.SKU == shirt2.SKU {
		t.Errorf("the sku of shirts must be diffrent %s == %s", shirt1.SKU, shirt2.SKU)
	}

	if shirt1 == shirt2 {
		t.Errorf("two cloned objects must be different object")
	}

	t.Logf("LOG: %s", shirt1.GetInfo())
	t.Logf("LOG: %s", shirt2.GetInfo())

	t.Logf("LOG: the memory address are different: %p != %p", &shirt1, &shirt2)

}