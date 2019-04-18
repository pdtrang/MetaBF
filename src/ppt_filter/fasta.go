//-----------------------------------------------------------------------------
// Author: Vinhthuy Phan, 2018
//-----------------------------------------------------------------------------
package ppt_filter

import (
	"bufio"
	"io"
	"log"
)

type FastaScanner struct {
	Header     string
	NextHeader string
	Seq        []byte
	Finished   bool
	Scanner    *bufio.Scanner
}

//-----------------------------------------------------------------------------
func NewFastaScanner(r io.Reader) *FastaScanner {
	scanner := &FastaScanner{Scanner: bufio.NewScanner(r), Finished: false}
	return scanner
}

//-----------------------------------------------------------------------------
func (s *FastaScanner) Scan() bool {
	if s.Finished {
		return false
	}
	var line []byte
	var flag bool
	// 1. Read Fasta header
	if s.NextHeader == "" {
		for flag = s.Scanner.Scan(); flag; flag = s.Scanner.Scan() {
			line = s.Scanner.Bytes()
			if len(line)==0 { continue }
			if line[0] == '>' {
				s.Header = string(line)
				break
			}
		}
		if err := s.Scanner.Err(); err != nil {
			log.Fatal(err)
		}
		// This happens when the file has no FASTA sequence
		if flag == false {
			return false
		}
	} else {
		s.Header = s.NextHeader
	}
	// 2. Read Fasta sequence
	var seq []byte
	for flag = s.Scanner.Scan(); flag; flag = s.Scanner.Scan() {
		line = s.Scanner.Bytes()
		if len(line)==0 { continue }
		if line[0] == '>' {
			s.NextHeader = string(line)
			break
		} else {
			seq = append(seq, line...)
		}
	}
	s.Seq = seq
	if err := s.Scanner.Err(); err != nil {
		log.Fatal(err)
	}
	s.Finished = !flag
	return true
}
