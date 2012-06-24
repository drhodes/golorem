// Copyright 2012 Derek A. Rhodes.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lorem

import "testing"
import "log"

func TestAll(t *testing.T) {
	for i := 1; i < 14; i++ {
		log.Print(word(i))
		for j := 1; j < 14; j++ {
			log.Print(Word(i, j))
			log.Print(Sentence(i, j))
			log.Print(Paragraph(i, j))
		}
		log.Print(Url())
		log.Print(Host())
		log.Print(Email())
	}
}
