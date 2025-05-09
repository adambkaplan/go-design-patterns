package decorator

import (
	"fmt"

	"github.com/adambkaplan/go-design-patterns/decorator/coffee"
)

func RunCoffeeShop() {
	var bev coffee.Beverage

	bev = &coffee.Espresso{}
	printBeverage(bev)

	// As noted in "Head First Design Patterns", this implementation is overly verbose. The
	// "Builder" and "Factory" patterns can be employed to make this code easier to comprehend.

	bev = &coffee.DarkRoast{}
	bev = &coffee.Mocha{
		Condiment: coffee.Condiment{
			Beverage: bev,
		},
	}
	bev = &coffee.Mocha{
		Condiment: coffee.Condiment{
			Beverage: bev,
		},
	}
	bev = &coffee.Whip{
		Condiment: coffee.Condiment{
			Beverage: bev,
		},
	}
	printBeverage(bev)

	bev = &coffee.HouseBlend{}
	bev = &coffee.Soy{
		Condiment: coffee.Condiment{
			Beverage: bev,
		},
	}
	bev = &coffee.Mocha{
		Condiment: coffee.Condiment{
			Beverage: bev,
		},
	}
	bev = &coffee.Whip{
		Condiment: coffee.Condiment{
			Beverage: bev,
		},
	}
	printBeverage(bev)

}

func printBeverage(bev coffee.Beverage) {
	fmt.Printf("%s $%.2f", bev.Description(), bev.Cost())
}
