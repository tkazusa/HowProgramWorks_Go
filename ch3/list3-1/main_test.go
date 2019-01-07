package main

import (
  "testing"
)

func Testmain(t *testing.T){
  expect := 100
  actual := main()

  if expect != actual{
    t.Errorf("%f != %f", expect, actual)
  }
}

