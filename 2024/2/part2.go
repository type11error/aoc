package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
  "math"
)

func direction(a,b int) int {
  if a < b {
    return 1
  }
  return -1
}

func main() {
  f, err := os.Open("./reports.txt")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  scn := bufio.NewScanner(f)

  safeReports := 0

  for scn.Scan() {
    line := scn.Text()
    nums := strings.Fields(line)
    
    var report []int
    for i := range nums {
      num,err := strconv.Atoi(nums[i])
      if err != nil {
        panic(err)
      }
      report = append(report, num)
    }

    previous := report[0]
    reportDirection := direction(report[0], report[1])
    safe := 0
    for _,v := range report[1:] {
      i := v - previous
      if ((i < 0 && reportDirection > 0) ||
         (i > 0 && reportDirection < 0)) ||
         math.Abs(float64(i)) > 3 ||
         i == 0 {
        safe++
      }

      previous = v
    }

    if safe < 2 {
      safeReports++
    }
  }

  fmt.Printf("%d\n", safeReports)
}
