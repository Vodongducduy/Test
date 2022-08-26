package main

import "fmt"

func main() {
	a := []int{2, 4, 6, 7, 15}
	b := []int{1, 3, 5, 10, 19, 20}
	var c []int
	var m int

	//n := len(a)
	//selectionSort([]int{8, 2, 10, 6, 1, 31, 0})
	//insertSort([]int{8, 2, 10, 6, 1, 31, 0})
	//interChangeSort([]int{8, 2, 10, 6, 1, 31, 0})
	//bubbleSort([]int{8, 2, 10, 6, 1, 31, 0})
	//bubbleSort([]int{8, 6, 34, 22, 40, 5, 11, 23, 44, 18})
	//ChanVeDau([]int{8, 6, 34, 22, 40, 5, 11, 23, 44, 18})
	//quickSort(a, 0, n-1)
	//TaoMang(a, &b, &m)
	Tron(a, b, &c, &m)
	fmt.Println(c, m)

}

//SelectionSort => tim gia tri nho nhat va swap voi vi tri tang dan tu 0
func selectionSort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				a[min], a[j] = a[j], a[min]
			}
		}
	}
	fmt.Println(a)
	return a
}

func insertSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		x := a[i]
		p := i - 1
		for p >= 0 && a[p] > x {
			a[p+1] = a[p]
			p--
		}
		a[p+1] = x
	}
	fmt.Println(a)
	return a
}

func interChangeSort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println(a)
	return a
}

func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(a)
}

func ChanVeDau(a []int) {
	vt := -1
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			vt++
			a[vt], a[i] = a[i], a[vt]
			fmt.Println("Check a", i, a)
		}
	}
}

func quickSort(a []int, left int, right int) {
	pivot := a[(left+right)/2]
	l := left
	r := right
	for l <= r {
		for a[l] < pivot {
			l++
		}

		for a[r] > pivot {
			r--
		}

		if l <= r {
			a[l], a[r] = a[r], a[l]
			l++
			r--
		}
	}
	if left < r {
		quickSort(a, left, r)
	}
	if l < right {
		quickSort(a, l, right)
	}
	fmt.Println(a)
}

func TaoMang(a []int, b *[]int, m *int) {
	*m = 0
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			*b = append(*b, 0)
			(*b)[*m] = a[i]
			*m++
		}
	}
}

func Tron(a []int, b []int, c *[]int, m *int) {
	var (
		i = 0
		j = 0
	)
	*m = 0
	for !(i >= len(a) && j >= len(b)) {
		if (i < len(a) && j < len(b) && a[i] < b[j]) || j >= len(b) {
			*c = append(*c, 0)
			(*c)[*m] = a[i]
			i++
			*m++
		} else {
			*c = append(*c, 0)
			(*c)[*m] = b[j]
			j++
			*m++
		}
		fmt.Println(*c)
	}
}

func MergeSort(a []int, left int, right int) {
	if left < right {

	}
}
