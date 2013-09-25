package mersenne

import (
	"errors"
	"flag"
	"fmt"
)

func uint64toUint32(i uint64) (uint32, error) {
	if i > 0xFFFFFFFF {
		return 0, errors.New("Input out of range")
	}
	return uint32(i), nil
}

var seedFlag = flag.Uint64("seed", 1, "Seeds the random number generator.")
var iterations = flag.Uint64("iteration", 0, "The iteration to resume generating random numbers from.")
var format = flag.String("format", "d", `
	Output format:
		b: binary
		o: octal
		d: decimal <default>
		x: lowercase hexadecimal
		X: uppercase hexadecimal
`)

func Main() {
	flag.Parse()
	seed, err := uint64toUint32(*seedFlag)
	if err != nil {
		panic("Seed must be an unsigned 32 bit integer.")
	}

	twister := New(seed)
	for i := uint64(0); i < *iterations; i += 1 {
		twister.Next()
	}

	switch *format {
	case "b":
		fmt.Printf("%b\n", twister.Get())
	case "o":
		fmt.Printf("%o\n", twister.Get())
	case "d":
		fmt.Printf("%d\n", twister.Get())
	case "x":
		fmt.Printf("%x\n", twister.Get())
	case "X":
		fmt.Printf("%X\n", twister.Get())
	default:
		fmt.Println("Invalid format, see help:")
		flag.PrintDefaults()
		panic("Invalid Format")
	}
}
