package singularity

import (
	"bytes"
	"fmt"
	"main/bitset"
	"math"
	"testing"
)

func TestXxx(t *testing.T) {
	for i := 0; i < math.MaxInt8; i++ {
		fmt.Printf("10%08b10\n",i)
	}
}
func TestEvilHunter(t *testing.T) {
	original := bitset.New().AppendBytes(GetTestBytes())
	fmt.Printf("original \n%s\n",original.InitStringFunc("",6))
	//newWorld, _ := Reformed(original)


}

// ShowMaxBranchInfo 打印最长分支的信息
func ShowMaxBranchInfo(name string, link []*bitset.Bitset) {
	index, maxBranch, flag := 0, 0, 0
	for i, j := 0, 0; i < len(link); i++ {
		if link[i].Len() > maxBranch {
			maxBranch = link[i].Len()
			index = j
			flag = ConditionToInt(link[i].At(0), 1, 0)
		}
		j += link[i].Len()
	}
	fmt.Printf("%v: flag = %v, MaxBranch = %-2v, First-index = %-5v\n", name, flag, maxBranch, index)
}

func TestCommunity(t *testing.T) {
	// 1: 0 1 0 2 0 3 0 4 0 5 0 0 1 0 0 0 1 0 0 0 0 1 0 0 0 0 0 1
	// 0: 1 0 1 0 0 1 0 0 0 1 0 0 0 0 1 0 0 0 0 0 2 0 3 0 4 0 5 0
	//original := bitset.NewFrom01String("01011011101111011111001000100001000001")
	//for i := 0; i < original.Len(); i++ {
	//fmt.Print(continuous1(original, 15), " ")
	//}
	// union2 写入文件
	//union2.InitStringFunc("", 8).WriteString("..\\sources\\1")

	// 读取Bytes
	original := bitset.New().AppendBytes(GetTestBytes())
	//original := bitset.NewFrom01String("1011010")
	//original, _ = Reformed(original)

	fmt.Printf("%s\n", bytes.Repeat([]byte("# "), 500))

	// 1. First
	link1, union1 := Republic(original)
	ShowMaxBranchInfo("Dim1", link1)
	fmt.Printf("%s\n\n", original.InitStringFunc("", 1))
	// union1 打印
	fmt.Printf("%s\n\n", union1.InitStringFunc("\n", len(link1)))
	// link1 打印：第二行加01隔离数
	counter := 0
	fmt.Printf("%s\n", union1.Sub(0, len(link1)).InitStringFunc("\n", len(link1)))
	for i := 1; i < union1.Len()/len(link1); i++ {
		for j := 0; j < len(link1); j++ {
			if i == 1 && union1.At(i*len(link1)+j) != union1.At(j) {
				fmt.Printf("%v", ConditionToInt(union1.At(j), 0, 1))
			} else if union1.At(i*len(link1)+j) != union1.At(j) {
				fmt.Print("_")
				counter++
			} else {
				fmt.Printf("%v", ConditionToInt(union1.At(i*len(link1)+j), 1, 0))
			}
		}
		fmt.Println()
	}
	fmt.Printf("`_`的数量 = %v\n", counter)
	// link1 打印：第二行不加01隔离数
	fmt.Printf("\n%s\n", union1.Sub(0, len(link1)).InitStringFunc("\n", len(link1)))
	for i := 1; i < union1.Len()/len(link1); i++ {
		for j := 0; j < len(link1); j++ {
			if union1.At(i*len(link1)+j) != union1.At(j) {
				fmt.Print("_")
			} else {
				fmt.Printf("%v", ConditionToInt(union1.At(i*len(link1)+j), 1, 0))
			}
		}
		fmt.Println()
	}
	fmt.Printf("%s\n", bytes.Repeat([]byte("# "), 500))

	// 2. Second
	link2, union2 := Republic(union1)
	ShowMaxBranchInfo("Dim2", link2)
	fmt.Printf("%s\n\n", union1.InitStringFunc("", 1))
	// union2 打印
	fmt.Printf("%s\n\n", union2.InitStringFunc("\n", len(link2)))
	// link2 打印：第二行加01隔离数
	counter = 0
	fmt.Printf("%s\n", union2.Sub(0, len(link2)).InitStringFunc("\n", len(link2)))
	for i := 1; i < union2.Len()/len(link2); i++ {
		for j := 0; j < len(link2); j++ {
			if i == 1 && union2.At(i*len(link2)+j) != union2.At(j) {
				fmt.Printf("%v", ConditionToInt(union2.At(j), 0, 1))
			} else if union2.At(i*len(link2)+j) != union2.At(j) {
				fmt.Print("_")
				counter++
			} else {
				fmt.Printf("%v", ConditionToInt(union2.At(i*len(link2)+j), 1, 0))
			}
		}
		fmt.Println()
	}
	fmt.Printf("`_`的数量 = %v\n", counter)
	// link2 打印：第二行不加01隔离数
	fmt.Printf("%s\n", union2.Sub(0, len(link2)).InitStringFunc("\n", len(link2)))
	for i := 1; i < union2.Len()/len(link2); i++ {
		for j := 0; j < len(link2); j++ {
			if union2.At(i*len(link2)+j) != union2.At(j) {
				fmt.Print("_")
			} else {
				fmt.Printf("%v", ConditionToInt(union2.At(i*len(link2)+j), 1, 0))
			}
		}
		fmt.Println()
	}
	fmt.Printf("%s\n", bytes.Repeat([]byte("# "), 500))
}

func TestCommunityEncode(t *testing.T) {
	//original := bitset.New().AppendBytes(GetTestBytes())
	//// original 打印
	//fmt.Printf("original %s\n", original.InitStringFunc("", 8))
	//// link union 处理
	//link := Republic(original)
	//union := Republic(link)
	//// link union写入文件
	//union.InitStringFunc("", 8).WriteString("..\\sources\\1")
	//// link union打印
	//fmt.Printf("union   %s\n", union.InitStringFunc("\n", len(link)))
	//fmt.Printf("第 1 行\n %s\n", union.Sub(0, len(link)).InitStringFunc("\n", len(link)))
	//for i := 1; i < union.Len()/len(link); i++ {
	//	for j := 0; j < len(link); j++ {
	//		if union.At(i*len(link)+j) != union.At(j) {
	//			fmt.Print(" ")
	//		} else {
	//			fmt.Printf("%v", ConditionToInt(union.At(i*len(link)+j), 1, 0))
	//		}
	//	}
	//	fmt.Println()
	//}
}

func TestCompress2(t *testing.T) {
	bits := bitset.NewFrom01String("011010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101100010011001011111100101100110001010001101001011110001010100111111001000000001010000000110101000101000001100100101110100010111010000101001101111000000100010111110000110101011010110000111010010101001001000000100100001111010110111100000001011001101001000001110101011100010111010111000101110011011111011111011100010100010110010101110001100101010100010001010110000100001101010101010000010101010011100100111110110001010011011100010101010101010100010101010101110101010101000101010100010101010101010001010101110111000100010110010110010101010101010101000101110101110001010101010101010101110101010101010101100101011101011101010101010101010101010101010111010101010101010101010101011001010111101100010100110111010101010101010101010101010101010101010101010101010101010101010101010101010101011101110101010101110101110101010101010101010001010101011101010101010101010101011101010101010101011101010111010101010101010101010101010101010101010101010101010101010101010110010101010001000101000101110101010101010101010101010101010101010101010101010101010101010101010101010101010111010101010101010101010101010101010101010100010101010101010101010101010101010111010101010101010111010101010101010101010101010101010101010101010101010101010101010101010101100101010100010101010101011101010101010101010101010101010101010101010101010101010101010101010101010101010101110101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101110101010101010101010101010101010101010101010101010101010101010101010101010001010101010101010101010111010101010101010101010101010101010101010101010101010101010101010101010101010101011101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101011101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010111010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010111010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101110101010101010101010101010101010101010101010101010101")
	pr1 := Compress(bits.Bytes())
	fmt.Printf("pr1 %v\n%08b\n", len(pr1)*8, pr1)
}

func TestRegression(t *testing.T) {
	template := bitset.NewFrom01String("11011")
	sub := template.Sub(0, template.Len())
	fmt.Printf("%s %v", sub, sub.Equals(template))
}

func TestTemp(t *testing.T) {

	original := bitset.New().AppendBytes(ReadFile("D:\\迅雷下载\\Postman-win64-9.1.5-Setup.exe"))
	original.WriteString("..\\sources\\Temp.txt")

}
