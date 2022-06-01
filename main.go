package main

import (
	"github.com/gusanmaz/jbst-processor/generator"
)

func main(){
	csvFiles := []string{"2022_April.csv", "2022_August.csv", "2022_December.csv"}
	generator.GenerateWebPage(csvFiles, "test", "www_static")
}
