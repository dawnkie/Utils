package singularity

import (
	"bytes"
	"fmt"
	"main/bitset"
	"math"
	"math/rand"
	"testing"
)

func TestCompress(t *testing.T) {
	// 1. 测试读写文件
	test1 := ReadFile("..\\sources\\0-测试文件")
	WriteFile("..\\sources\\test1", test1)
	temp := ReadFile("..\\sources\\test1")
	if !bytes.Equal(test1, temp) {
		t.Error("比对 ReadFile 和 WriteFile 不通过!")
	}

	// 2. 测试压缩解压缩
	test2 := Compress(test1)
	WriteFile("..\\sources\\test2", test2)
	test2 = ReadFile("..\\sources\\test2")
	temp = Uncompress(test2)
	if !bytes.Equal(test1, temp) {
		t.Error("比对 压缩解压缩 不通过!")
	}

	// 3. 测试文件追加
	testString := "Test Append!"
	fbs1 := ReadFile("..\\sources\\append.txt")
	AppendFile("..\\sources\\append.txt", []byte(testString))
	fbs2 := ReadFile("..\\sources\\append.txt")
	if string(fbs1)+testString != string(fbs2) {
		t.Error("测试 文件追加 不通过!")
	}
}

func TestCompressCube(t *testing.T) {
	WriteFile("..\\sources\\1.普通压缩", Compress(ReadFile("..\\sources\\0-测试文件")))
	WriteFile("..\\sources\\1.普通解压", Uncompress(ReadFile("..\\sources\\1.普通压缩")))

	WriteFile("..\\sources\\2.cube压缩", CompressCube(ReadFile("..\\sources\\0-测试文件")))
	WriteFile("..\\sources\\2.cube解压", Uncompress(ReadFile("..\\sources\\2.cube压缩")))
}

func TestBitsetGetBytes(t *testing.T) {
	fbs := bytes.Repeat(GetTestBytes(), 10)
	//fmt.Println(bset)
	for i := 1; i <= 101; i += 10 {
		temp := bytes.Repeat(fbs, i)
		InitialSide(len(temp))
		bset := bitset.New()
		bset.AppendBytes(temp)
		planeZ := DimensionOfZ(bset)
		bs := planeZ.Bytes()
		fmt.Printf("bset.Bytes() len = %v\n", len(bs))
	}
	//if !bytes.Equal(fbs, bs) {
	//	t.Error("比对不通过!")
	//}
}

func TestSide(t *testing.T) {
	fmt.Printf("Side = %v\n", Side)

	InitialSide(4096)
	fmt.Printf("4096 Side = %v\n", Side)
	InitialSide(5349314)
	fmt.Printf("5349314 Side = %v\n", Side)
	InitialSide(5088448 * 8)
	fmt.Printf("5088448*8 Side = %v\n", Side)

	fmt.Printf("KB Side = %v => %v\n", math.Cbrt(float64(KB)), CalculateSide(KB))
	fmt.Printf("MB Side = %v => %v\n", math.Cbrt(float64(MB)), CalculateSide(MB))
	fmt.Printf("GB Side = %v => %v\n", math.Cbrt(float64(GB)), CalculateSide(GB))
	fmt.Printf("TB Side = %v => %v\n", math.Cbrt(float64(TB)), CalculateSide(TB))
	fmt.Printf("PB Side = %v => %v\n", math.Cbrt(float64(PB)), CalculateSide(PB))

	//InitialSide(511)
	//fmt.Printf("Side = %v\n",Side)
}

func TestFileBytes(t *testing.T) {
	ReadFile("..\\sources\\cube.txt")
}

func TestSplitCube(t *testing.T) {
	//cube, carryover := SplitCube(ReadFile("..\\sources\\cube.txt"))
	//cube.Print()
	cube, carryover := SplitCube(GetTestBytes())
	fmt.Println(cube)
	fmt.Println(carryover)
}

func TestDimensionOfX(t *testing.T) {
	testValue := "11111010000101001110100011100110111011101001110011111101001110100010101000101100111111111000010101101111100100011001111000110010110110110001010001001001001110001100010010001000111101010100101010001000101100101001100011101111110000000110010001110011110111100000011110010000010101000001111010111111101100010010010100101011011010001001101010101100111001101001001101100110001010101001100010100100100110010100100011011010101111111110011011110100010010001010111100010110011110001010010010110101010110100000101011101010"
	bits := BytesToBitset(GetTestBytes())
	if !DimensionOfX(bits).Equals(bitset.NewFrom01String(testValue)) {
		t.Errorf("%s", DimensionOfX(bits))
		t.Errorf("Wrong => %v, want %v", DimensionOfX(bits), testValue)
	}
}

func TestDimensionOfY(t *testing.T) {
	testValue := "10110011101001001100011110100000111101010000101111101001001010010001111000000000010100001011111101000110111000010001010100000010110101001110101111000101010110011111011101001100010000000110000011010111100101101011010100101010001110101101110110011110010100001011111111110000110010110000111111001010111010111100111001001111111010100001011100011110110010011010000110010110000001110100100011100010101100101011111011110010110001011110101001010101101110000001010100110011110010011101010010111111000100001111100100001000"
	bits := BytesToBitset(GetTestBytes())
	if !DimensionOfY(bits).Equals(bitset.NewFrom01String(testValue)) {
		t.Errorf("%s", DimensionOfY(bits))
		t.Errorf("Wrong => %v, want %v", DimensionOfY(bits), testValue)
	}
}

// 测试Cube对数据复杂度的影响，结论：Cube置换并不会改变复杂度
func TestStatistics(t *testing.T) {
	bits1 := bitset.New().AppendBytes(ReadFile("..\\sources\\0-测试文件"))
	Statistics("0-测试文件", bits1, true)
	Statistics("0-测试文件", bits1, false)

	//cube, carryover := SplitCube(ReadFile("..\\sources\\0-测试文件"))
	//Statistics("DimensionOfZ",DimensionOfZ(cube).AppendBytes(carryover),true)
	//Statistics("DimensionOfZ",DimensionOfZ(cube).AppendBytes(carryover),false)
	//
	//cube, carryover = SplitCube(ReadFile("..\\sources\\0-测试文件"))
	//Statistics("DimensionOfY",DimensionOfY(cube).AppendBytes(carryover),true)
	//Statistics("DimensionOfY",DimensionOfY(cube).AppendBytes(carryover),false)
	//
	//cube, carryover = SplitCube(ReadFile("..\\sources\\0-测试文件"))
	//Statistics("DimensionOfZ->DimensionOfZ",DimensionOfZ(DimensionOfZ(cube)).AppendBytes(carryover),true)
	//Statistics("DimensionOfZ->DimensionOfZ",DimensionOfZ(DimensionOfZ(cube)).AppendBytes(carryover),false)
	//
	//cube, carryover = SplitCube(ReadFile("..\\sources\\0-测试文件"))
	//Statistics("DimensionOfY->DimensionOfY",DimensionOfY(DimensionOfY(cube)).AppendBytes(carryover),true)
	//Statistics("DimensionOfY->DimensionOfY",DimensionOfY(DimensionOfY(cube)).AppendBytes(carryover),false)
	//
	//cube, carryover = SplitCube(ReadFile("..\\sources\\0-测试文件"))
	//Statistics("DimensionOfY->DimensionOfZ",DimensionOfZ(DimensionOfY(cube)).AppendBytes(carryover),true)
	//Statistics("DimensionOfY->DimensionOfZ",DimensionOfZ(DimensionOfY(cube)).AppendBytes(carryover),false)
	//
	//cube, carryover = SplitCube(ReadFile("..\\sources\\0-测试文件"))
	//Statistics("DimensionOfZ->DimensionOfY",DimensionOfY(DimensionOfZ(cube)).AppendBytes(carryover),true)
	//Statistics("DimensionOfZ->DimensionOfY",DimensionOfY(DimensionOfZ(cube)).AppendBytes(carryover),false)
}

func TestDimensionOfZ(t *testing.T) {
	testValue := "10111110101110101011101111000111101011110101111010011001000000100011011000101000111010010010011111101010011110101010101100111100100011101010101100010010110100101011010101001010100000011010001011111001000111110101011001100011101100010001010101010011000100100100110000100000000011110111110000011001101110101001100110001111011110011001010010110110010010011110001100110100010111100000100011011110001101111000111001011010011110011000111000011100010010001001100100100101101110010110110010100111110110001100011110001000"
	bits := BytesToBitset(GetTestBytes())
	if !DimensionOfZ(bits).Equals(bitset.NewFrom01String(testValue)) {
		t.Errorf("%s", DimensionOfZ(bits))
		t.Errorf("Wrong => %v, want %v", DimensionOfZ(bits), testValue)
	}
}

func TestPickXPillar(t *testing.T) {
	testIndex := []int{0, 5, 12, Side*Side - 1}
	testValue := []string{
		"11111010",
		"10011100",
		"01101111",
		"11101010",
	}
	bits := BytesToBitset(GetTestBytes())
	for i := 0; i < len(testIndex); i++ {
		if !PickXPillar(bits, testIndex[i]).Equals(bitset.NewFrom01String(testValue[i])) {
			t.Errorf("%s", PickXPillar(bits, i))
			t.Errorf("Wrong => %v, want %v", bitset.NewFrom01String(testValue[i]), PickXPillar(bits, i))
		}
	}
}

/*
* 1. Y柱和Z柱的有效位并不像X柱依次位于`第[01234567]位`
* 2. Y柱和Z柱的有效位位于 `bi%Side`
 */
func TestPickYPillar(t *testing.T) {
	testIndex := []int{0, 5, 12, Side*Side - 1}
	testValue := []string{
		"10110011",
		"00001011",
		"01000110",
		"00001000",
	}
	bits := BytesToBitset(GetTestBytes())
	for i := 0; i < len(testIndex); i++ {
		if !PickYPillar(bits, testIndex[i]).Equals(bitset.NewFrom01String(testValue[i])) {
			t.Errorf("%s", PickYPillar(bits, i))
			t.Errorf("Wrong => %v, want %v", bitset.NewFrom01String(testValue[i]), PickXPillar(bits, i))
		}
	}
}

/*
* 1. Y柱和Z柱的有效位并不像X柱依次位于`第[01234567]位`
* 2. Y柱和Z柱的有效位位于 `bi%Side`
 */
func TestPickZPillar(t *testing.T) {
	testIndex := []int{0, 5, 12, Side*Side - 1}
	testValue := []string{
		"10111110",
		"01011110",
		"11101010",
		"10001000",
	}
	bits := BytesToBitset(GetTestBytes())
	for i := 0; i < len(testIndex); i++ {
		if !PickZPillar(bits, testIndex[i]).Equals(bitset.NewFrom01String(testValue[i])) {
			t.Errorf("%s", PickZPillar(bits, i))
			t.Errorf("Wrong => %v, want %v", bitset.NewFrom01String(testValue[i]), PickXPillar(bits, i))
		}
	}
}

/*
	测试环境：const Side = 8
*/
func TestBytesToBitset(t *testing.T) {
	bs := GetTestBytes()
	bits := BytesToBitset(bs)
	fmt.Println(bits.String())
}

func TestShowBytes01(t *testing.T) {

}

func TestShowBytesHex(t *testing.T) {

}

func TestReverse(t *testing.T) {

}

/*
GetTestBytes():

	Layer[0] :
	1  1  1  1  1  0  1  0
	0  0  0  1  0  1  0  0
	1  1  1  0  1  0  0  0
	1  1  1  0  0  1  1  0
	1  1  1  0  1  1  1  0
	1  0  0  1  1  1  0  0
	1  1  1  1  1  1  0  1
	0  0  1  1  1  0  1  0

	Layer[1] :
	0  0  1  0  1  0  1  0
	0  0  1  0  1  1  0  0
	1  1  1  1  1  1  1  1
	1  0  0  0  0  1  0  1
	0  1  1  0  1  1  1  1
	1  0  0  1  0  0  0  1
	1  0  0  1  1  1  1  0
	0  0  1  1  0  0  1  0

	Layer[2] :
	1  1  0  1  1  0  1  1
	0  0  0  1  0  1  0  0
	0  1  0  0  1  0  0  1
	0  0  1  1  1  0  0  0
	1  1  0  0  0  1  0  0
	1  0  0  0  1  0  0  0
	1  1  1  1  0  1  0  1
	0  1  0  0  1  0  1  0

	Layer[3] :
	1  0  0  0  1  0  0  0
	1  0  1  1  0  0  1  0
	1  0  0  1  1  0  0  0
	1  1  1  0  1  1  1  1
	1  1  0  0  0  0  0  0
	0  1  1  0  0  1  0  0
	0  1  1  1  0  0  1  1
	1  1  0  1  1  1  1  0

	Layer[4] :
	0  0  0  0  0  1  1  1
	1  0  0  1  0  0  0  0
	0  1  0  1  0  1  0  0
	0  0  0  1  1  1  1  0
	1  0  1  1  1  1  1  1
	1  0  1  1  0  0  0  1
	0  0  1  0  0  1  0  1
	0  0  1  0  1  0  1  1

	Layer[5] :
	0  1  1  0  1  0  0  0
	1  0  0  1  1  0  1  0
	1  0  1  0  1  1  0  0
	1  1  1  0  0  1  1  0
	1  0  0  1  0  0  1  1
	0  1  1  0  0  1  1  0
	0  0  1  0  1  0  1  0
	1  0  0  1  1  0  0  0

	Layer[6] :
	1  0  1  0  0  1  0  0
	1  0  0  1  1  0  0  1
	0  1  0  0  1  0  0  0
	1  1  0  1  1  0  1  0
	1  0  1  1  1  1  1  1
	1  1  1  0  0  1  1  0
	1  1  1  1  0  1  0  0
	0  1  0  0  1  0  0  0

	Layer[7] :
	1  0  1  0  1  1  1  1
	0  0  0  1  0  1  1  0
	0  1  1  1  1  0  0  0
	1  0  1  0  0  1  0  0
	1  0  1  1  0  1  0  1
	0  1  0  1  1  0  1  0
	0  0  0  0  1  0  1  0
	1  1  1  0  1  0  1  0
*/
func GetTestBytes() (fbs []byte) {
	return []byte{0b11111010, 0b00010100, 0b11101000, 0b11100110, 0b11101110, 0b10011100, 0b11111101, 0b00111010, 0b00101010, 0b00101100, 0b11111111, 0b10000101, 0b01101111, 0b10010001, 0b10011110, 0b00110010, 0b11011011, 0b00010100, 0b01001001, 0b00111000, 0b11000100, 0b10001000, 0b11110101, 0b01001010, 0b10001000, 0b10110010, 0b10011000, 0b11101111, 0b11000000, 0b01100100, 0b01110011, 0b11011110, 0b00000111, 0b10010000, 0b01010100, 0b00011110, 0b10111111, 0b10110001, 0b00100101, 0b00101011, 0b01101000, 0b10011010, 0b10101100, 0b11100110, 0b10010011, 0b01100110, 0b00101010, 0b10011000, 0b10100100, 0b10011001, 0b01001000, 0b11011010, 0b10111111, 0b11100110, 0b11110100, 0b01001000, 0b10101111, 0b00010110, 0b01111000, 0b10100100, 0b10110101, 0b01011010, 0b00001010, 0b11101010}
}

func TestRandomBitsYLayer(t *testing.T) {
	for i := 0; i < 8; i++ {
		//fmt.Printf("Layer[%v] : \n",i)
		for j := 1; j <= Side*Side*Side/8; j++ {
			fmt.Print(rand.Intn(2), "")
			if j%8 == 0 {
				//fmt.Println()
			}
		}
		//fmt.Println()
	}
}

func TestRandomBytes(t *testing.T) {
	// 0b11111010, 0b00010100, 0b11101000,...
	for i := 0; i < 8; i++ {
		for j := 1; j <= Side*Side*Side/8; j++ {
			if (j-1)%8 == 0 {
				fmt.Print("0b")
			}
			fmt.Print(rand.Intn(2))
			if j%8 == 0 {
				fmt.Print(", ")
			}
		}
	}
}
