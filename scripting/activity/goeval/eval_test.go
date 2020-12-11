/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/
package goeval

import (
	"fmt"
	"testing"
)

func TestExecute(t *testing.T) {
	s := NewScope()
	s.Set("print", fmt.Println)
	t.Log(s.Eval(`count := 0`))
	t.Log(s.Eval(`for i:=0;i<100;i=i+1 { 
			count=count+i
		}`))
	t.Log(s.Eval(`print(count)`))
}
