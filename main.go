package main

import (
	"fmt"
	//"github.com/fatih/color"
	"log"
	"os"
	"strings"
)

var (
	didInit bool = false // establish first peak/valley

	bitBool  bool         // allows byte conversion to boolean
	prevBit  bool         // previous bit in stream of binary
	bitSlice []string     // array of all the binary that has passed through the system
	bitIndex uint     = 0 // bit # in the net total of bits

	PositionP   *Pattern
	PositionV   *Pattern
	PeakDepth   uint
	ValleyDepth uint
)

type Pattern struct {
	counter  uint     // # of times this struct has been accessed
	positive *Pattern // pos = 1 (bit value) | Pointer to next bit in pattern
	negative *Pattern // neg = 0 (bit value) | Pointer to next bit in pattern
	occures  []uint   // array of where the pattern(s) start
}

func main() {
	PeakSeed := new(Pattern)
	ValleySeed := new(Pattern)
	findPatternsIn("t2.txt")
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
			mask := i - 1                      //creates eight different bitmasks to read each bit from a byte individualy
			bit := b1[0] & (1 << mask) >> mask //binary opperations on bitmask
			if bit == 0 {
				bitSlice = append(bitSlice, "0") //add bit to "history"
				bitBool = false                  // converts bit (type byte) to type bool (simplifies comparision)
				if !didInit {
					didInit = true
				} else {
					if bitBool != prevBit {
						//PEAK
						PositionP = PeakSeed
						PeakDepth = 0
						//return to top of peak tree
					} else {
						//currentPosition := return peaks
						//keep traveling down peak tree
					}
				}
				prevBit = false

			} else if bit == 1 {
				bitSlice = append(bitSlice, "1") //add bit to "history"
				bitBool = true                   // converts bit (type byte) to type bool (simplifies comparision)
				if !didInit {
					didInit = true
				} else {
					if bitBool != prevBit {
						//VALLEY
						PositionV = ValleySeed
						PeakDepth = 0
						//return to top of valley tree
					} else {
						//PositionV =
						//keep traveling down valley tree
					}
				}
				prevBit = true
			}
			bitIndex++
		}
	}
}
