/*
* BSD 3-Clause License
* Copyright © 2021. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*
* utility.renderJSON($activity[LiveAppsFindCases].caseRefs.cases, false)
* URI: "https://www.example.com/logicscripts/dyngetTop.js"
*
*/

//console.log("input: " + input);

var arr = JSON.parse(input);

arr.sort(function(a, b) {
    return b.summary.OverallPoints - a.summary.OverallPoints;
});

var list = arr.slice(0, 5);
var result = '{"top":[';

for (var key in list) {
    if (list.hasOwnProperty(key)) {
        console.log(key + " -> " + list[key].summary.Nickname + " | " + list[key].summary.OverallPoints);
        result += '{"name":"'+list[key].summary.Nickname+'","points":'+list[key].summary.OverallPoints+'},';
    }
}

result = result.slice(0, -1);
result += ']}'

var feedback = result;
feedback;