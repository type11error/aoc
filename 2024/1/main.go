package main

import (
  "fmt"
  "os"
  "bufio"
  "io"
  "sort"
)

func diff(a,b int) int {
  if a < b {
    return b - a
  }
  return a - b
}

func main() {
  f, err := os.Open("./input.txt")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  scn := bufio.NewScanner(f)

  var l1 []int
  var l2 []int
  d := 0
  s := 0

  for scn.Scan() {
    var a, b int
    _, err := fmt.Sscanf(scn.Text(), "%d %d", &a, &b)

    if err == io.EOF {
      break
    }

    if err != nil {
      fmt.Printf("%s\n", err)
      panic(err)
    }

    l1 = append(l1,a)
    l2 = append(l2,b)

  }

  sort.Ints(l1)
  sort.Ints(l2)

  for i,_ := range l1 {
    fmt.Printf("a=%d b=%d\n", l1[i], l2[i])
    d += diff(l1[i], l2[i])
  }

  fmt.Printf("d=%d\n", d)

  // Part 2

  // Create map of 2nd list
  l2Map := map[int]int{}
  for i,_ := range l1 {
    if _,exist := l2Map[l2[i]]; !exist {
      l2Map[l2[i]] = 1
    } else {
      l2Map[l2[i]]++
    }
  }

  // Iterate over 1st list, look up out of map and add
  for _,key := range l1 {
    v,exists := l2Map[key]
    if exists {
      x := key * v
      s += x
    }
  }

  fmt.Printf("s=%d\n", s)
}
