/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/
input := "{input}"
feedback := ""

if (input=="green"){
	feedback = "25"
	
} else {
	feedback = "75"
}

return feedback