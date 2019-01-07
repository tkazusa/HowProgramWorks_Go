package main

import "fmt"
 
func main() {
	var sum float64;
  var i int;

  /* clear var sum as zero */
  sum = 0;

  for i = 1; i <= 100; i++ {
    sum += 0.1;
  };

	fmt.Printf("%f", sum);
};
