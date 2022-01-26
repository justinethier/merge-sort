package merge

import (
  //"fmt"
)

func mergeSort(arr []int) []int{
  if len(arr) <= 1 {
    return arr
  }

  mid := len(arr) / 2
  l := arr[0:mid]
  r := arr[mid:]

  l = mergeSort(arr[0:mid])
  r = mergeSort(arr[mid:])

  // TODO: Merge two sorted sides of array

  // TODO: copy elements from l/r to arr
  // TODO: copy any remaining items from l
  // TODO: copy any remaining items from r

  return nil // TODO
}
