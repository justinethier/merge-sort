package merge

import (
  //"fmt"
  "math/rand"
  "reflect"
  "testing"
)

func TestMerge(t *testing.T) {
  var arr []int
  var shuffled = make([]int, 100)

  for i := 0; i < 100; i++ {
    arr = append(arr, i)
  }
  copy(shuffled, arr)

  for i := range arr {
    j := rand.Intn(100)
    shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
  }

  shuffled = mergeSort(shuffled)

  if !reflect.DeepEqual(arr, shuffled) {
    t.Error("Array is not sorted", shuffled)
  }
}
