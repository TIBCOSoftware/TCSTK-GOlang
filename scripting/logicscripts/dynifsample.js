/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/
var feedback;

if (input=="green"){
	console.log("JS VM: green path selected");
	feedback = "25";
} else {
	console.log("JS VM: other path selected");
	feedback = "75";
};

feedback;