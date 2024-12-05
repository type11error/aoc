package main

import (
  "fmt"
  "os"
  "bufio"
  "io"
)

type State int 

const (
  Start State = iota
  M
  U
  L
  LParen
  ANum
  Comma
  BNum
  RParen
  D
  O
  N
  Ap
  T
  DoLP
  DontLP
)

func isDigit(char rune) bool {
    return char >= '0' && char <= '9'
}

func main() {
  f, err := os.Open("./input.txt")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  state := Start
  var a int
  var b int
  result := 0
  enabled := true

  scn := bufio.NewReader(f)
  for {
    if r,s,err := scn.ReadRune(); err != nil {
      if err == io.EOF {
        break
      } else {
       panic(err)
      }
    } else {

      switch state {
      case Start:
        a = 0
        b = 0
        if string(r) == "m" {
          state = M
        } else if string(r) == "d" {
          state = D
        }
      case M: 
        fmt.Printf("m\n")
        if string(r) == "u" {
          state = U
        } else {
          state = Start
        }
      case U: 
        fmt.Printf("u\n")
        if string(r) == "l" {
          state = L
        } else {
          state = Start
        }
      case L: 
        fmt.Printf("l\n")
        if string(r) == "(" {
          state = LParen
        } else {
          state = Start
        }
      case LParen: 
        fmt.Printf("(\n")
        if isDigit(r) {
          a = a*10 + int(r- '0')
        } else if string(r) == "," {
          state = Comma
        } else {
          state = Start
        }
      case Comma: 
        if isDigit(r) {
          b= b*10 + int(r- '0')
        } else if string(r) == ")" {
          if enabled {
            result = result + (a*b)
          }
          state = Start
        } else {
          state = Start
        }
      case D: 
        if string(r) == "o" {
          state = O
        } else {
          state = Start
        }
      case O: 
        if string(r) == "n" {
          state = N
        } else if string(r) == "(" {
          state = DoLP
        } else {
          state = Start
        }
      case DoLP: 
        if string(r) == ")" {
          state = Start
          enabled = true
        } else {
          state = Start
        }
      case N: 
        if string(r) == "'" {
          state = Ap
        } else {
          state = Start
        }
      case Ap: 
        if string(r) == "t" {
          state = T
        } else {
          state = Start
        }
      case T: 
        if string(r) == "(" {
          state = DontLP
        } else {
          state = Start
        }
      case DontLP: 
        if string(r) == ")" {
          state = Start
          enabled = false
        } else {
          state = Start
        }

      }
      fmt.Printf("r=%q s=%d\n",string(r),s)
    }

  }
  fmt.Printf("result=%d\n",result)
}
