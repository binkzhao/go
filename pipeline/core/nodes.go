package core

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

// 开始排序的时间
var startTime = time.Now()

// 数据进入channel
func ArraySource(a ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, value := range a {
			out <- value
		}
		close(out)
	}()

	return out
}

// 单个切片排序(内部排序节点)
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int, 1024)

	go func() {
		// Read into memory
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		fmt.Println("Read done: ", time.Now().Sub(startTime))

		// Sort
		sort.Ints(a)
		fmt.Println("InMemSort done: ", time.Now().Sub(startTime))

		// Output
		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}

// 多个已排序好的切片进行归并Merge排序(归并节点)
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)

	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		fmt.Println("Merge done: ", time.Now().Sub(startTime))
	}()

	return out
}

func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}

	middle := len(inputs) / 2
	return Merge(
		MergeN(inputs[:middle]...),
		MergeN(inputs[middle:]...),
	)
}

// 从文件流读取数据
func ReaderSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)

	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			if n > 0 {
				bytesRead += n
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}

			// chunkSize = -1 表示读到底
			if err != nil || (chunkSize != -1 && bytesRead > chunkSize) {
				break
			}
		}
		close(out)
	}()

	return out
}

// 把数据写入文件流
func WriterSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

// 生成随机数
func RandomSource(count int) <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()

	return out
}
