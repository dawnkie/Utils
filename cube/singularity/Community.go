package singularity

import (
	"fmt"
	"log"
	"main/bitset"
	"math"
)

var Dim = 0

const M0 = "00100"
const M1 = "11011"

// Republic 走向共和
func Republic(original *bitset.Bitset) (link []*bitset.Bitset, union *bitset.Bitset) {
	// link
	link = make([]*bitset.Bitset, 0)
	for i, n := 0, 0; i < original.Len(); {
		n = ConditionToInt(original.At(i), continuous1(original, i), continuous0(original, i))
		link = append(link, bitset.New().AppendNBools(n, original.At(i)))
		i += n
	}
	//// secondLink
	//secondLink := make([]*bitset.Bitset, 0)
	//firstTwoRows := bitset.NewWithSize(len(link) * 2)
	//for i := 0; i < len(link); i++ {
	//	if link[i].Len() == 1 {
	//		firstTwoRows.Set(link[i].At(0), i)
	//		firstTwoRows.Set(!link[i].At(0), i+len(link))
	//	} else {
	//		firstTwoRows.Set(link[i].At(0), i)
	//		firstTwoRows.Set(link[i].At(1), i+len(link))
	//	}
	//}
	//// 说明：i 为当前循环回合，j 为第二行游标
	//for i, j, n := 0, len(link), 0; i < len(link); i++ {
	//	n = ConditionToInt(firstTwoRows.At(j), continuous1(firstTwoRows, j), continuous0(firstTwoRows, j))
	//	// 存储列元素
	//	if firstTwoRows.At(j) == firstTwoRows.At(j-len(link)) {
	//		secondLink = append(secondLink, bitset.New().Append(link[j-len(link)].Sub(1, link[j-len(link)].Len())))
	//	} else {
	//		secondLink = append(secondLink, bitset.New().AppendBools(firstTwoRows.At(j)))
	//	}
	//	// 存储相邻列元素
	//	for k := 1; k < n; k++ {
	//		if firstTwoRows.At(j+k) == firstTwoRows.At(j-len(link)+k) {
	//			secondLink[i].AppendBools(!firstTwoRows.At(j + k))
	//			secondLink[i].Append(link[j-len(link)+k].Sub(1, link[j-len(link)+k].Len()))
	//		} else {
	//			secondLink[i].AppendBools(!firstTwoRows.At(j+k), firstTwoRows.At(j+k))
	//		}
	//	}
	//	if n == 1 {
	//		j++
	//	} else {
	//		j += n
	//	}
	//	if j == len(link)*2 {
	//		break
	//	}
	//}
	//// union
	//union = bitset.New()
	//maxBranch := MaxBranch(secondLink)
	//for i := 0; i < maxBranch; i++ {
	//	for j := 0; j < len(secondLink); j++ {
	//		if secondLink[j].Len() > i {
	//			union.AppendBools(ConditionToBool(secondLink[j].At(i), true, false))
	//		} else {
	//			union.AppendBools(ConditionToBool(secondLink[j].At(0), false, true))
	//		}
	//	}
	//}

	union = bitset.New()
	maxBranch := MaxBranch(link)
	for i := 0; i < maxBranch; i++ {
		for j := 0; j < len(link); j++ {
			if link[j].Len() > i {
				union.AppendBools(ConditionToBool(link[j].At(i), true, false))
			} else {
				union.AppendBools(ConditionToBool(link[j].At(0), false, true))
			}
		}
	}
	// Dimension
	Dim++
	return link, union
}

// Reformed 数据格式：[自编码标志位 1bit][编码方法标志位 1bit][step 8bit][基准 step位][evils... step的倍数位][evils结尾0串 step位]
func Reformed(origin *bitset.Bitset) (newWorld *bitset.Bitset, tail *bitset.Bitset) {
	newWorld = bitset.New()
	for i, n := 0, 0; i < origin.Len()/6; i++ {
		newWorld.Append(AngelKeep(int(origin.GetUint64(n, n+6))))
		n += 6
	}
	if origin.Len()%6 != 0 {
		tail = origin.Sub(origin.Len()-origin.Len()%6, origin.Len())
		fmt.Printf("newWorld-tail \n%s\n", tail)
		return newWorld, tail

	}
	fmt.Printf("newWorld \n%s\n", newWorld.InitStringFunc("", 12))
	return newWorld, nil
}
func DevilCome(angel *bitset.Bitset) int {
	switch angel.String() {
	case "100000000010":
		return 0
	case "100000000110":
		return 1
	case "100000001010":
		return 2
	case "100000001110":
		return 3
	case "100000010110":
		return 4
	case "100000011010":
		return 5
	case "100000011110":
		return 6
	case "100000101010":
		return 7
	case "100000101110":
		return 8
	case "100000110010":
		return 9
	case "100000111010":
		return 10
	case "100000111110":
		return 11
	case "100001010010":
		return 12
	case "100001010110":
		return 13
	case "100001011010":
		return 14
	case "100001011110":
		return 15
	case "100001100010":
		return 16
	case "100001100110":
		return 17
	case "100001101010":
		return 18
	case "100001110010":
		return 19
	case "100001111010":
		return 20
	case "100001111110":
		return 21
	case "100010100010":
		return 22
	case "100010100110":
		return 23
	case "100010101010":
		return 24
	case "100010101110":
		return 25
	case "100010110010":
		return 26
	case "100010111010":
		return 27
	case "100010111110":
		return 28
	case "100011000010":
		return 29
	case "100011000110":
		return 30
	case "100011001010":
		return 31
	case "100011001110":
		return 32
	case "100011010010":
		return 33
	case "100011010110":
		return 34
	case "100011100010":
		return 35
	case "100011100110":
		return 36
	case "100011101010":
		return 37
	case "100011110010":
		return 38
	case "100011111010":
		return 39
	case "100011111110":
		return 40
	case "100101000010":
		return 41
	case "100101000110":
		return 42
	case "100101001010":
		return 43
	case "100101001110":
		return 44
	case "100101010010":
		return 45
	case "100101010110":
		return 46
	case "100101011010":
		return 47
	case "100101011110":
		return 48
	case "100101100010":
		return 49
	case "100101100110":
		return 50
	case "100101101010":
		return 51
	case "100101110010":
		return 52
	case "100101111010":
		return 53
	case "100101111110":
		return 54
	case "100110000010":
		return 55
	case "100110000110":
		return 56
	case "100110001010":
		return 57
	case "100110001110":
		return 58
	case "100110010110":
		return 59
	case "100110011010":
		return 60
	case "100110011110":
		return 61
	case "100110100010":
		return 62
	case "100110100110":
		return 63
	default:
		log.Panicln("[ERROR] Angel Number must be 0-63!")
		return -1
	}
}
func AngelKeep(devil int) *bitset.Bitset {
	switch devil {
	case 0:
		return bitset.NewFrom01String("100000000010")
	case 1:
		return bitset.NewFrom01String("100000000110")
	case 2:
		return bitset.NewFrom01String("100000001010")
	case 3:
		return bitset.NewFrom01String("100000001110")
	case 4:
		return bitset.NewFrom01String("100000010110")
	case 5:
		return bitset.NewFrom01String("100000011010")
	case 6:
		return bitset.NewFrom01String("100000011110")
	case 7:
		return bitset.NewFrom01String("100000101010")
	case 8:
		return bitset.NewFrom01String("100000101110")
	case 9:
		return bitset.NewFrom01String("100000110010")
	case 10:
		return bitset.NewFrom01String("100000111010")
	case 11:
		return bitset.NewFrom01String("100000111110")
	case 12:
		return bitset.NewFrom01String("100001010010")
	case 13:
		return bitset.NewFrom01String("100001010110")
	case 14:
		return bitset.NewFrom01String("100001011010")
	case 15:
		return bitset.NewFrom01String("100001011110")
	case 16:
		return bitset.NewFrom01String("100001100010")
	case 17:
		return bitset.NewFrom01String("100001100110")
	case 18:
		return bitset.NewFrom01String("100001101010")
	case 19:
		return bitset.NewFrom01String("100001110010")
	case 20:
		return bitset.NewFrom01String("100001111010")
	case 21:
		return bitset.NewFrom01String("100001111110")
	case 22:
		return bitset.NewFrom01String("100010100010")
	case 23:
		return bitset.NewFrom01String("100010100110")
	case 24:
		return bitset.NewFrom01String("100010101010")
	case 25:
		return bitset.NewFrom01String("100010101110")
	case 26:
		return bitset.NewFrom01String("100010110010")
	case 27:
		return bitset.NewFrom01String("100010111010")
	case 28:
		return bitset.NewFrom01String("100010111110")
	case 29:
		return bitset.NewFrom01String("100011000010")
	case 30:
		return bitset.NewFrom01String("100011000110")
	case 31:
		return bitset.NewFrom01String("100011001010")
	case 32:
		return bitset.NewFrom01String("100011001110")
	case 33:
		return bitset.NewFrom01String("100011010010")
	case 34:
		return bitset.NewFrom01String("100011010110")
	case 35:
		return bitset.NewFrom01String("100011100010")
	case 36:
		return bitset.NewFrom01String("100011100110")
	case 37:
		return bitset.NewFrom01String("100011101010")
	case 38:
		return bitset.NewFrom01String("100011110010")
	case 39:
		return bitset.NewFrom01String("100011111010")
	case 40:
		return bitset.NewFrom01String("100011111110")
	case 41:
		return bitset.NewFrom01String("100101000010")
	case 42:
		return bitset.NewFrom01String("100101000110")
	case 43:
		return bitset.NewFrom01String("100101001010")
	case 44:
		return bitset.NewFrom01String("100101001110")
	case 45:
		return bitset.NewFrom01String("100101010010")
	case 46:
		return bitset.NewFrom01String("100101010110")
	case 47:
		return bitset.NewFrom01String("100101011010")
	case 48:
		return bitset.NewFrom01String("100101011110")
	case 49:
		return bitset.NewFrom01String("100101100010")
	case 50:
		return bitset.NewFrom01String("100101100110")
	case 51:
		return bitset.NewFrom01String("100101101010")
	case 52:
		return bitset.NewFrom01String("100101110010")
	case 53:
		return bitset.NewFrom01String("100101111010")
	case 54:
		return bitset.NewFrom01String("100101111110")
	case 55:
		return bitset.NewFrom01String("100110000010")
	case 56:
		return bitset.NewFrom01String("100110000110")
	case 57:
		return bitset.NewFrom01String("100110001010")
	case 58:
		return bitset.NewFrom01String("100110001110")
	case 59:
		return bitset.NewFrom01String("100110010110")
	case 60:
		return bitset.NewFrom01String("100110011010")
	case 61:
		return bitset.NewFrom01String("100110011110")
	case 62:
		return bitset.NewFrom01String("100110100010")
	case 63:
		return bitset.NewFrom01String("100110100110")
	default:
		log.Panicln("[ERROR] Angel Number must be 0-63!")
		return nil
	}
}

func Ultima(primitive *bitset.Bitset, m string, isFirst bool) (newWorld *bitset.Bitset) {
	// 1. 创世阶段，缉查外表为`m`的恶魔
	evils := EvilHunter(primitive, m, isFirst)
	if evils == nil {
		return primitive
	}
	jail := append(make([]*bitset.Bitset, 0), evils)
	ShowEvilsInfo(evils)
	for evils.HasSub(m) {
		evils = EvilHunter(evils, m, false)
		if evils != nil {
			jail = append(jail, evils)
			ShowEvilsInfo(evils)
		}
	}
	fmt.Printf("jail %v\n", jail)

	// 2. 创世阶段，重整秩序
	order := bitset.New()
	for i := len(jail) - 1; i >= 0; i-- {
		order.Append(jail[i])
	}
	order.Append(primitive)
	fmt.Printf("order %s\n", order)

	// 3. 创世阶段，构建新世界
	newWorld = bitset.New()
	for i, t := 0, 0; i < order.Len(); {
		index := order.IndexOfSub(m, i)
		fmt.Printf("index|%v ", index)
		if index != -1 {
			newWorld.Append(order.Sub(i, index))
			t = index
			i = index + len(m)
		} else {
			newWorld.Append(order.Sub(t+len(m), order.Len()))
			break
		}
	}

	fmt.Printf("\nnewWorld %s\n", newWorld)
	return newWorld
}

// EvilHunter 数据格式：[自编码][编码方法][step][首evil][evils][结尾符]
// 数据格式：[自编码][编码方法][step][首evil][evils][结尾符]
// 自编码: 标志位 [0:1]，自编码=1，无自编码=0
// 编码方法: 标志位 [1:2]，"11011" =1, "00100"=0
// step: [2:10]
// 首evil: 标志位 [10:11], 首evil序号!=0->0, 首evil序号==0->1
// evils： evils 长度step+1的倍数
// 结尾符: 结尾符 长度step+1，若方法为"11011"，则结尾符为`0..`，若方法为"00100"，则结尾符为`1..`
func EvilHunter(primitive *bitset.Bitset, m string, isFirst bool) (evils *bitset.Bitset) {
	if primitive == nil || !primitive.HasSub(m) {
		return nil
	}
	// evilsList
	evilsList := make([]int, 0)
	t := bitset.NewFrom01String(m)
	for i := 0; i < primitive.Len()-t.Len(); {
		if primitive.Sub(i, i+t.Len()).Equals(t) {
			evilsList = append(evilsList, i)
			i += t.Len()
		}
		i++
	}
	// gap
	gap := evilsList[0]
	for i, n := 0, 0; i < len(evilsList)-1; i++ {
		n = evilsList[i+1] - evilsList[i]
		if n > gap {
			gap = n - t.Len()
		}
	}
	// step
	step := 0
	for i, n := 0, gap; i < gap; i++ {
		step++
		n = n / 2
		if n == 0 {
			break
		}
	}
	// 自编码 编码方法
	evils = bitset.NewWithSize(2)
	if !isFirst {
		evils.Set(true, 0)
	}
	// step
	step++
	evils.AppendUint32(uint32(step), 8)
	// 首evil
	evils.AppendBools(ConditionToBool(evilsList[0] == 0, true, false))
	// evils
	evils.AppendUint32(uint32(evilsList[0]), step)
	for i := 0; i < len(evilsList)-1; i++ {
		evils.AppendUint32(uint32(evilsList[i+1]-evilsList[i]-t.Len()), step)
	}
	switch m {
	case "11011":
		// 编码方法
		evils.Set(true, 1)
		// 结尾符
		evils.AppendUint32(uint32(0), step) // [结尾符 step位]
	case "00100":
		// 编码方法
		evils.Set(false, 1)
		// 结尾符
		evils.AppendUint32(math.MaxUint32, step) // [结尾符 step位]
	}

	return evils

}

func ShowEvilsInfo(evils *bitset.Bitset) {
	if evils != nil && evils.Len() > 10 {
		fmt.Printf("1. 是否为自编码 = %v\n", evils.At(0))
		fmt.Printf("2. 编码方法 = %v(1：11011,0：00100)\n", ConditionToInt(evils.At(1), 1, 0))
		step := int(evils.Sub(2, 10).Bytes()[0])
		evilIndex := ConditionToInt(evils.At(10), 1, 0)
		fmt.Printf("3. step = %v\n", step)
		fmt.Printf("4. 是否evilIndex为0 = %v\n", ConditionToBool(evilIndex == 0, false, true))
		fmt.Printf("5. evils %s\n\t\t ", evils.InitStringFunc("", 1))
		onlyEvils := evils.Sub(11, evils.Len()-step)
		for i := 0; i < onlyEvils.Len()-1; i += step {
			fmt.Printf("%s/%d ", onlyEvils.Sub(i, i+step), onlyEvils.Sub(i, i+step).GetUint64(0, onlyEvils.Len()))
		}
		fmt.Println()
	}
}

// MaxBranch 返回最长分支的长度
func MaxBranch(link []*bitset.Bitset) (maxBranch int) {
	for i := 0; i < len(link); i++ {
		if link[i].Len() > maxBranch {
			maxBranch = link[i].Len()
		}
	}
	return maxBranch
}

// Regression 回归原始
func Regression(m *bitset.Bitset, s *bitset.Bitset) (original *bitset.Bitset) {
	original = bitset.New()
	for i, j := 0, 0; i < m.Len(); {
		if !m.At(i) {
			original.AppendBools(false)
			i++
		} else {
			if !s.At(j) {
				original.AppendBools(true, false)
			} else {
				n := continuous1(m, i)
				original.AppendNBools(n+1, true)
				original.AppendBools(false)
				i += n
			}
			j++
		}
	}
	return original
}

func continuous1(bits *bitset.Bitset, from int) (n int) {
	for i := from; i < bits.Len(); i++ {
		if bits.At(i) {
			n++
		} else {
			return n
		}
	}
	return n
}
func continuous0(bits *bitset.Bitset, from int) (n int) {
	for i := from; i < bits.Len(); i++ {
		if !bits.At(i) {
			n++
		} else {
			return n
		}
	}
	return n
}
