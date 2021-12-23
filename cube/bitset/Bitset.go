package bitset

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// Bitset stores bits
type Bitset struct {
	// The set stored how many bits.
	size int

	// Storage
	set []byte
}

// New returns an initialised Bitset.
func New(bools ...bool) (bits *Bitset) {
	bits = &Bitset{size: 0, set: make([]byte, 0)}
	bits.AppendBools(bools...)

	return bits
}

// NewWithSize return an initialised Bitset.
func NewWithSize(size int) (bits *Bitset) {
	offset := 0
	if size%8 != 0 {
		offset = 1
	}
	return &Bitset{size: size, set: make([]byte, size/8+offset)}
}

// NewFrom01String constructs and returns a Bitset from 01-string, e.g. "1010 0101".
func NewFrom01String(str01 string) (bits *Bitset) {
	bits = &Bitset{size: 0, set: make([]byte, 0)}

	for _, c := range str01 {
		switch c {
		case '1':
			bits.AppendBools(true)
		case '0':
			bits.AppendBools(false)
		case ' ':
		default:
			log.Panicf("Invalid char %c in NewFrom01String", c)
		}
	}

	return bits
}

// Clone returns a copy.
func (bits *Bitset) Clone() (newBits *Bitset) {
	size := bits.size / 8
	if bits.size%8 != 0 {
		size++
	}
	newSet := make([]byte, size)
	copy(newSet, bits.set)

	return &Bitset{size: bits.size, set: newSet}
}

func (bits *Bitset) Set(v bool, index int) {
	if bits.At(index) != v {
		if v {
			bits.set[index/8] |= 0x80 >> uint(index%8)
		} else {
			n := 1
			for i := 0; i < 7-index%8; i++ {
				n *= 2
			}
			bits.set[index/8] = byte(int(bits.set[index/8]) - n)
		}
	}
}

// GetUint64 returns an integer of the from-end bits
func (bits *Bitset) GetUint64(from, end int) (res uint64) {
	sub := bits.Sub(from, end)
	for i, n := sub.Len()-1, 0; i >= 0; i-- {
		res += uint64(ConditionToInt(sub.At(i), int(math.Pow(2, float64(n))), 0))
		n++
	}
	return res
}

// Append a bitset of other-copied, new size = bits.Len() + other.Len()
func (bits *Bitset) Append(other *Bitset) *Bitset {
	bits.ensureCapacity(other.size)

	for i := 0; i < other.size; i++ {
		if other.At(i) {
			bits.set[bits.size/8] |= 0x80 >> uint(bits.size%8)
		}
		bits.size++
	}
	return bits
}

// AppendByte appends the size of trunc from bt.
func (bits *Bitset) AppendByte(bt byte, trunc int) *Bitset {
	bits.ensureCapacity(trunc)

	if trunc > 8 {
		log.Panicf("Param `trunc` %d out of range 0-8", trunc)
	}

	for i := trunc - 1; i >= 0; i-- {
		if bt&(1<<uint(i)) != 0 {
			bits.set[bits.size/8] |= 0x80 >> uint(bits.size%8)
		}

		bits.size++
	}
	return bits
}

// AppendBytes appends a list of whole bytes.
func (bits *Bitset) AppendBytes(bs []byte) *Bitset {
	for _, d := range bs {
		bits.AppendByte(d, 8)
	}
	return bits
}

// AppendUint32 appends the size least significant set from value.
func (bits *Bitset) AppendUint32(n uint32, trunc int) *Bitset {
	bits.ensureCapacity(trunc)

	if trunc > 32 {
		log.Panicf("Param `trunc` %d out of range 0-32", trunc)
	}

	for i := trunc - 1; i >= 0; i-- {
		if n&(1<<uint(i)) != 0 {
			bits.set[bits.size/8] |= 0x80 >> uint(bits.size%8)
		}

		bits.size++
	}
	return bits
}

// AppendBools appends a bitset of bools
func (bits *Bitset) AppendBools(bools ...bool) *Bitset {
	bits.ensureCapacity(len(bools))

	for _, v := range bools {
		if v {
			bits.set[bits.size/8] |= 0x80 >> uint(bits.size%8)
		}
		bits.size++
	}
	return bits
}

// AppendNBools appends b n-times repeated
func (bits *Bitset) AppendNBools(n int, b bool) *Bitset {
	for i := 0; i < n; i++ {
		bits.AppendBools(b)
	}
	return bits
}

// At returns the false|0 or true|1 at index.
func (bits *Bitset) At(index int) bool {
	if index >= bits.size {
		log.Panicf("Index %d out of range", index)
	}

	return (bits.set[index/8] & (0x80 >> byte(index%8))) != 0
}

// Sub returns a sub of the bitset from start to end
func (bits *Bitset) Sub(start int, end int) (subBits *Bitset) {
	if start > end || end > bits.size {
		log.Panicf("Out of range start=%d end=%d size=%d", start, end, bits.size)
	}

	subBits = New()
	subBits.ensureCapacity(end - start)

	for i := start; i < end; i++ {
		if bits.At(i) {
			subBits.set[subBits.size/8] |= 0x80 >> uint(subBits.size%8)
		}
		subBits.size++
	}

	return subBits
}

func (bits *Bitset) HasSub(ss ...string) bool {
	for _, s := range ss {
		t := NewFrom01String(s)
		for i := 0; i < bits.Len()-len(s); {
			if bits.Sub(i, i+len(s)).Equals(t) {
				return true
			}
			i++
		}
	}
	return false
}

func (bits *Bitset) IndexOfSub(s string, from int) (index int) {
	t := NewFrom01String(s)
	for i := from; i < bits.Len()-len(s); {
		if bits.Sub(i, i+len(s)).Equals(t) {
			return i
		}
		i++
	}
	return -1
}

// ByteAt returns a byte of the bitset from index to index+8
func (bits *Bitset) ByteAt(index int) byte {
	if index < 0 || index >= bits.size {
		log.Panicf("Index %d out of range", index)
	}

	var result byte

	for i := index; i < index+8 && i < bits.size; i++ {
		result <<= 1
		if bits.At(i) {
			result |= 1
		}
	}

	return result
}

// Len returns the size of the Bitset.
func (bits *Bitset) Len() int {
	return bits.size
}

// Bools returns the bools of the Bitset.
func (bits *Bitset) Bools() []bool {
	result := make([]bool, bits.size)

	var i int
	for i = 0; i < bits.size; i++ {
		result[i] = (bits.set[i/8] & (0x80 >> byte(i%8))) != 0
	}

	return result
}

// Bytes returns the Bytes of the Bitset.
func (bits *Bitset) Bytes() []byte {
	if bits.size >= 8 {
		size := bits.size / 8
		if (bits.size % 8) != 0 {
			size++
		}
		return bits.set[:size]
	} else {
		return []byte{bits.set[0] >> uint8(8-bits.size%8)}
	}
}

// Equals returns `if the Bitset equals other`.
func (bits *Bitset) Equals(other *Bitset) bool {
	if bits.size != other.size {
		return false
	}

	if !bytes.Equal(bits.set[0:bits.size/8], other.set[0:bits.size/8]) {
		return false
	}

	for i := 8 * (bits.size / 8); i < bits.size; i++ {
		a := bits.set[i/8] & (0x80 >> byte(i%8))
		b := other.set[i/8] & (0x80 >> byte(i%8))

		if a != b {
			return false
		}
	}

	return true
}

// ensureCapacity ensures the Bitset has enough size to store additional bits.
func (bits *Bitset) ensureCapacity(size int) {
	size += bits.size

	btSize := size / 8
	if size%8 != 0 {
		btSize++
	}

	if len(bits.set) >= btSize {
		return
	}

	bits.set = append(bits.set, make([]byte, btSize+2*len(bits.set))...)
}

// Print output "[01 ]*" to console, cause `String()` maybe too long to console(just can't )
func (bits *Bitset) Print() {
	for i := 0; i < bits.size; i++ {
		if (i % 8) == 0 {
			fmt.Print(" ")
		}

		if (bits.set[i/8] & (0x80 >> byte(i%8))) != 0 {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
}

// separator `String()` separator, step is the interval of separator
var separator = " "
var step = 8

func (bits *Bitset) InitStringFunc(_separator string, _step int) *Bitset {
	separator = _separator
	step = _step
	if step < 1 {
		fmt.Printf("\n[WARNING] InitStringFunc(): Invalid Param `_step`(must >= 1), Now is %v. Automatically switches to 8\n\n", _step)
		step = 8
	}
	return bits
}

// String returns "[01]+8[separator]" string of the Bitset
func (bits *Bitset) String() string {
	var sb strings.Builder
	for i := 0; i < bits.size; i++ {
		if (i%step) == 0 && i != 0 {
			sb.WriteString(separator)
		}
		if (bits.set[i/8] & (0x80 >> byte(i%8))) != 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
	}
	//return fmt.Sprintf("size=%d, set=%s", bits.size, sb.String())
	return sb.String()
}

func (bits *Bitset) WriteString(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < bits.size; i++ {
		if (i % step) == 0 {
			_, _ = file.Write([]byte(separator))
		}
		if (bits.set[i/8] & (0x80 >> byte(i%8))) != 0 {
			_, _ = file.Write([]byte("1"))
		} else {
			_, _ = file.Write([]byte("0"))
		}
	}
}

func ConditionToInt(condition bool, Y, N int) int {
	if condition {
		return Y
	} else {
		return N
	}
}
