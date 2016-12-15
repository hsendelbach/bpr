package main

import (
	"fmt"
	//"github.com/fatih/color"
	"log"
	"os"
	//"strings"
)

var (
	didInit bool = false // establish first peak/valley

	bitBool  bool         // allows byte conversion to boolean
	prevBit  bool         // previous bit in stream of binary
	bitSlice []string     // array of all the binary that has passed through the system
	BitIndex uint     = 0 // bit # in the net total of bits ~~~~ check whether needs to be global

	PSeed  *Pattern
	VSeed  *Pattern
	P      *Pattern
	V      *Pattern
	PDepth uint
	VDepth uint
)

type Pattern struct {
	counter  uint     // # of times a given pattern has been collected
	positive *Pattern // pos = 1 (bit value) | Pointer to next bit in pattern
	negative *Pattern // neg = 0 (bit value) | Pointer to next bit in pattern
	occurs   []uint   // array of where the pattern(s) start
}

func main() {
	PSeed = new(Pattern)
	VSeed = new(Pattern)
	findPatternsIn("test.md")
}

func findPatternsIn(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err) // shut down the program if something goes wrong with reading the file
	}
	defer file.Close()            // force graceful file close regardless of what happens
	singleByte := make([]byte, 1) // # of bytes to read per itteration (do not change unless you modify the bitmasking below)
	for {
		readByte, err := file.Read(singleByte) //read one byte from the file
		if readByte <= 0 {                     //EOF?
			break // end of file, close down the loop
		}
		if err != nil {
			log.Fatal(err) // shut down the program if something goes wrong with reading the byte
		}

		for i := uint(8); i > 0; i-- {
			mask := i - 1                              //creates eight different bitmasks to read each bit from a byte individualy
			bit := singleByte[0] & (1 << mask) >> mask //binary opperations on bitmask
			if bit == 0 {
				bitSlice = append(bitSlice, "0") //add bit to "history"
				bitBool = false                  // converts bit (type byte) to type bool (simplifies comparision)
				if !didInit {
					P = PSeed
					V = VSeed
					didInit = true // foces prevBit to be set before comparing it
				} else {
					if bitBool != prevBit {
						//PEAK
						P.occurs = append(P.occurs, BitIndex)
						P.counter++
						fmt.Println("P ", P)
						P = PSeed
						PDepth = 0
						//perhaps flip this
					}
				}
				prevBit = false

			} else if bit == 1 {
				bitSlice = append(bitSlice, "1") //add bit to "history"
				bitBool = true                   // converts bit (type byte) to type bool (simplifies comparision)
				if !didInit {
					P = PSeed
					V = VSeed
					didInit = true // foces prevBit to be set before comparing it
				} else {
					if bitBool != prevBit {
						//VALLEY
						V.occurs = append(V.occurs, BitIndex)
						V.counter++
						fmt.Println(bitSlice[:BitIndex], bitSlice[BitIndex:])
						V = VSeed
						VDepth = 0
						//perhaps flip this
					}
				}
				prevBit = true
			}
			P = P.NextBitPointer(bitBool)
			V = V.NextBitPointer(bitBool)
			PDepth++
			VDepth++
			BitIndex++
		}
	}
}

func (p *Pattern) NextBitPointer(b bool) *Pattern {
	if b {
		if p.positive == nil {
			p.positive = new(Pattern)
		}
		return p.positive
	} else {
		if p.negative == nil {
			p.negative = new(Pattern)
		}
		return p.negative
	}
}
