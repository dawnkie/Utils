package singularity

import (
	"fmt"
	"main/bitset"
	"math"
)

// Side 边长(8的倍数)，注意：X = Side/8 Bytes, Y = Side Bytes, Z = Side Bytes
// 测试环境：const Side = 8
var Side = 8
//var Area = Side * Side

const MinSize = 512
const KB = 1024
const MB = 1024 * KB
const GB = 1024 * MB
const TB = 1024 * GB
const PB = 1024 * TB

// BytesToBitset 将字节数组转为bit数组
func BytesToBitset(bs []byte) (bits *bitset.Bitset) {
	bits = bitset.New()
	bits.AppendBytes(bs)
	return bits
}

// InitialSide 根据数据集长度初始化最大"立方体"的边长
func InitialSide(size int) {
	if size < MinSize {
		panic("[ERROR] The Param `size` is too small!")
	}
	Side = int(math.Cbrt(float64(size))) / 8 * 8
}

// CalculateSide 根据数据集长度计算最大"立方体"的边长
func CalculateSide(size int64) (calSide int) {
	if size < MinSize {
		panic("[ERROR] The Param `size` is too small!")
	}
	return int(math.Cbrt(float64(size))) / 8 * 8
}

// SplitCube 将数据分解为立方体和尾数据 X = Side/8 Bytes, Y = Side Bytes, Z = Side Bytes
func SplitCube(fbs []byte) (cube *bitset.Bitset, carryover []byte) {
	InitialSide(len(fbs))
	cube = bitset.New()
	spots := Side * Side * Side / 8
	cube.AppendBytes(fbs[:spots])
	carryover = fbs[spots:]
	return cube, carryover
}

//DimensionOfX 沿X轴将立方体上的一组组(一条条)地bit抽取组合为新的立方体(注意：该函数生成的新立方体与旧方体完全一致)
func DimensionOfX(bits *bitset.Bitset) (planeX *bitset.Bitset) {
	planeX = bitset.New()
	for i := 0; i < Side*Side; i++ {
		planeX.Append(PickXPillar(bits, i))
	}
	return planeX
}

// DimensionOfY 沿Y轴将立方体上的一组组(一条条)地bit抽取组合为新的立方体
func DimensionOfY(bits *bitset.Bitset) (planeY *bitset.Bitset) {
	planeY = bitset.New()
	for i := 0; i < Side*Side; i++ {
		planeY.Append(PickYPillar(bits, i))
	}
	return planeY
}

//func UndimensionOfY(planeY *bitset.Bitset) (bits *bitset.Bitset) {
//	bits = bitset.NewWithSize(planeY.Len())
//	for i := 0; i < Side*Side; i++ {
//		planeY.Append(PickYPillar(bits, i))
//	}
//	pillar = bitset.New()
//	for i, j := 0, bi; i < Side; i++ {
//		pillar.AppendBools(bits.At(j))
//		j += Side*Side
//	}
//	return pillar
//	return planeY
//}

// DimensionOfZ 沿Z轴将立方体上的一组组(一条条)地bit抽取组合为新的立方体
func DimensionOfZ(bits *bitset.Bitset) (planeZ *bitset.Bitset) {
	fmt.Printf("\nSide = %v\n",Side)
	planeZ = bitset.New()
	for i := 0; i < Side*Side; i++ {
		planeZ.Append(PickZPillar(bits, i))
	}
	return planeZ
}

// PickXPillar 前提：假设X面的bit自右向左自增编号 实现：抽取指定编号沿X轴的一组(一条)bit
func PickXPillar(bits *bitset.Bitset, bi int) (pillar *bitset.Bitset) {
	return bits.Sub(Side*bi, Side*bi+Side)
}

// PickYPillar 前提：假设Y面的bit自左向右自增编号 实现：抽取指定编号沿Y轴的一组(一条)bit
func PickYPillar(bits *bitset.Bitset, bi int) (pillar *bitset.Bitset) {
	pillar = bitset.New()
	for i, j := 0, bi; i < Side; i++ {
		pillar.AppendBools(bits.At(j))
		j += Side*Side
	}
	return pillar
}

// PickZPillar 前提：假设Z面的bit自左向右自增编号 实现：抽取指定编号沿Z轴的一组(一条)bit
func PickZPillar(bits *bitset.Bitset, bi int) (pillar *bitset.Bitset) {
	pillar = bitset.New()
	for i, j := 0, bi%Side+bi/Side*Side*Side; i < Side; i++ {
		pillar.AppendBools(bits.At(j))
		j += Side
	}
	return pillar
}

func Reverse(bits []byte) {
	for i := 0; i < len(bits)/2; i++ {
		bits[i], bits[len(bits)-1-i] = bits[len(bits)-1-i], bits[i]
	}
}