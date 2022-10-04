package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

// List of all sequences (8, 16, 32, ...)
type NumArray struct {
	nums []Num
}

// All required data on certain sequence
type Num struct {
	SeqNum       int64
	KeysAmnt     big.Int
	randKey      big.Int
	BruteElapsed time.Duration
}

func FloatToBigInt(bigval big.Float) *big.Int {
	// library magic
	bigInt := big.NewInt(0)
	bigInt, _ = bigval.Int(nil)
	return bigInt
}

func (n *Num) GenKeysAmnt() {
	// Amount of keys that can be generated from this sequence equal 2 (a bit)
	// in the power of amount of bits in a sequence
	n.KeysAmnt.Exp(big.NewInt(2), big.NewInt(n.SeqNum), nil)
}

func (n Num) Print() {
	fmt.Printf("Byte sequence: %d\n", n.SeqNum)
	fmt.Printf("All possible keys: %s\n", n.KeysAmnt.String())
	fmt.Printf("Random key: %s\n", n.randKey.String())
	fmt.Printf("Time for brute-force: %d ms\n\n", n.BruteElapsed.Milliseconds())
}

func (arr *NumArray) fillSequence() {

	// Used smallest sequences for tests since time for bruteforce is enormous
	// seq := []int64{8, 16, 32, 64, 128, 256, 512, 1024, 2048}
	seq := []int64{8, 16}

	// Creating Num structure for every sequence
	for i := 0; i < len(seq); i++ {
		arr.nums = append(arr.nums, Num{})
	}

	// Setting a sequence for every Num
	for i := 0; i < len(arr.nums); i++ {
		arr.nums[i].SeqNum = seq[i]
	}
}

// General print function
func (arr NumArray) Print() {
	for i := 0; i < len(arr.nums); i++ {
		arr.nums[i].Print()
	}
}

// Calculating total amount of keys, generating a random key and bruteforcing
func (arr NumArray) Work() {
	for i := 0; i < len(arr.nums); i++ {
		// Calculating total amount of keys
		arr.nums[i].GenKeysAmnt()

		// Generating a random key
		rand.Seed(time.Now().Unix())
		rnd := big.NewFloat(rand.Float64())
		var tempFloatKey big.Float
		tempFloatKey.Mul(new(big.Float).SetInt(&arr.nums[i].KeysAmnt), rnd)
		arr.nums[i].randKey = *FloatToBigInt(tempFloatKey)

		// Bruteforcing
		var key big.Int
		key.SetInt64(0)
		start := time.Now()
		for true {
			if key.Cmp(&arr.nums[i].randKey) == 0 {
				fmt.Println("GOTCHA")
				break
			} else {
				key = *key.Add(&key, big.NewInt(1))
			}
		}
		arr.nums[i].BruteElapsed = time.Since(start)
	}
}

func main() {
	var arrN NumArray
	arrN.fillSequence()
	arrN.Work()
	arrN.Print()
}
