package main

import (
	"./ppt_filter"
	// "fmt"
	"flag"
	"log"
)

func main() {
	// get inputs from commandline
	filter_saved_file := flag.String("load", "", "filter saved file")
    read_1 := flag.String("r1", "", "fastq/fq file")
    read_2 := flag.String("r2", "", "fastq/fq file")
    out := flag.String("out", "result.txt", "output filename")
    strategy := flag.String("strategy", "majority", "querying strategy")
    flag.Parse()
	
	// Load filter
	log.Printf("Load filter")
    f := ppt_filter.Load(*filter_saved_file)
    // fmt.Println(f.Gid)
    // f.Summarize()	
    log.Println("Finish loading filter.")

    upper_threshold := float32(0.3)
    lower_threshold := float32(0.1)

	if *read_2 == "" {
		f.OnlineSingleQuery(*read_1, *out, *strategy, upper_threshold, lower_threshold)	
	} else {
		f.OnlinePairQuery(*read_1, *read_2, *out, *strategy, upper_threshold, lower_threshold)
	}

}


