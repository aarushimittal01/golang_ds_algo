// Write a function which takes following two arguments:
// 1. An array "a" of integers.
// 2. A positive integer "k".
// The function should return the kth largest number in the argument array “a”.
// 
// Example: a = [-1, 100, 0, 1, 5, 2], k = 3
// Answer = 2 (k = 3, so 3rd largest number in array “a” is 2)

func getElement(a []int, l, r, k int) int {
   if l == r {
       return a[l]
   }
   
   key := sortedList(a, l, r)
   
   if k == key {
       return arr[p]
   } else if k < key {
       return getElement(a, l, key+1, k)
   } else {
       return getElement(a, key+1, r, k)
   }
}


func sortedList(a []int, min, max int) int {
    key := a[max]
    i := min - 1
    
    for j := min; j < max: j++ {
        if a[j] < key {
            i++
            a[j], a[i] = a[i], a[j]
        }
    }
    
    a[i+1], a[max] = a[max], a[j]
    
    return i+1
}

func getKthLargestGolang(a []int, int k) int {
    if len(a) > 0 {
        return getElement(a, 0, len(a) -1)
    } else {
        return 0
    }
}


int getKthLargestJava(int[] a, int k) {
}

def getKthLargest(a, k):
