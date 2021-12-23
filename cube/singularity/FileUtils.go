package singularity

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"main/bitset"
	"os"
)

// ReadFile 将目标路径的文件读取为字节数组
func ReadFile(path string) (fbs []byte) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fbs, err = ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	return fbs
}

// WriteFile 向目标路径的文件覆盖写入字节数组，不存在文件则创建
func WriteFile(path string, fbs []byte) {
	_ = ioutil.WriteFile(path, fbs, 0644)
}

// AppendFile 向目标路径的文件追加写入字节数组，不存在文件则创建
func AppendFile(path string, fbs []byte) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = file.Write(fbs)
}

// Compress zlib压缩
func Compress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	_, _ = w.Write(src)
	_ = w.Close()
	return in.Bytes()
}

// CompressCube Z立方体的zlib压缩
func CompressCube(src []byte) []byte {
	bits := bitset.New()
	bits.AppendBytes(src)
	cube, carryover := SplitCube(src)
	return Compress(DimensionOfZ(cube).AppendBytes(carryover).Bytes())
}

// Uncompress zlib解压
func Uncompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	_, _ = io.Copy(&out, r)
	return out.Bytes()
}

// Statistics 统计连续0和1的数量，flag 代表0和1
func Statistics(name string, bits *bitset.Bitset, flag bool) {
	// 统计重复2、3、4、5、6+次
	rep1, rep2, rep3, rep4, rep5, rep6, rep7, rep8, rep9, repn := 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
	counter := 1
	for i := 0; i < bits.Len()-1; i++ {
		if bits.At(i) == flag {
			if i+1 < bits.Len() && bits.At(i+1) == flag {
				counter++
				continue
			} else {
				switch counter {
				case 0:
				case 1:
					if i+1 != bits.Len() {
						rep1++
					}
				case 2:
					rep2++
				case 3:
					rep3++
				case 4:
					rep4++
				case 5:
					rep5++
				case 6:
					rep6++
				case 7:
					rep7++
				case 8:
					rep8++
				case 9:
					rep9++
				default:
					repn++
				}
				counter = 1
			}
		}
	}
	fmt.Printf("%s\n", bytes.Repeat([]byte("-"), 50))
	fmt.Printf("名字：\"%v\"(总位数 = %v), 其中\"%v\"的统计情况如下：\n", name, bits.Len(), ConditionToInt(flag,1,0))
	fmt.Printf("\tep1 = %v\n", rep1)
	fmt.Printf("\trep2 = %v\n", rep2)
	fmt.Printf("\trep3 = %v\n", rep3)
	fmt.Printf("\trep4 = %v\n", rep4)
	fmt.Printf("\trep5 = %v\n", rep5)
	fmt.Printf("\trep6 = %v\n", rep6)
	fmt.Printf("\trep7 = %v\n", rep7)
	fmt.Printf("\trep8 = %v\n", rep8)
	fmt.Printf("\trep9 = %v\n", rep9)
	fmt.Printf("\trepn = %v\n", repn)
	fmt.Printf("%s\n", bytes.Repeat([]byte("-"), 50))
}
