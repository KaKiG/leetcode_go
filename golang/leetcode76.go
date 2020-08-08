package main

import "fmt"

func main() {
  s := "AA"
  t := "A"
  fmt.Println(minWindow(s, t))
}


func minWindow(s string, t string) string {
  left := 0
  right := 0
  minLen := len(s)+1
  str := ""
  if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
    return ""
  }
  freqMapt := make(map[string]int)
  for i := range t {
    freqMapt[string(t[i])]++
  }
  freqMaps := make(map[string]int)
  freqMaps[string(s[0])]++
  length := right - left + 1
  for right < len(s) {
    if IsValid(freqMaps,freqMapt) {
      length = right - left + 1
      if minLen > length {
        minLen = length
        str = s[left:right+1]
      }
      freqMaps[string(s[left])]--
      left ++
    } else {
      right++
      if right >= len(s) {
        break
      }
      freqMaps[string(s[right])]++
    }
  }
  return str
}

func IsValid(freqMaps, freqMapt map[string]int) bool {
  for k,v := range freqMapt {
    _,exist := freqMaps[k]
    if (!exist) || v > freqMaps[k] {
      return false
    }
  }
  return true
}
