package merge

import (
  //"fmt"
)

func mergeSort(arr []int) []int{
  if len(arr) <= 1 {
    return arr
  }

  mid := len(arr) / 2

  l := mergeSort(arr[0:mid])
  r := mergeSort(arr[mid:])

  // Merge two sorted sides of array
  var rv []int
  var ll, rr int = 0, 0

  for (ll < len(l) && rr < len(r)) {
    if l[ll] < r[rr] {
      rv = append(rv, l[ll])
      ll++
    } else {
      rv = append(rv, r[rr])
      rr++
    }
  }

  // copy any remaining items from l
  for ll < len(l) {
    rv = append(rv, l[ll])
    ll++
  }

  // copy any remaining items from r
  for rr < len(r) {
    rv = append(rv, r[rr])
    rr++
  }

  return rv
}
