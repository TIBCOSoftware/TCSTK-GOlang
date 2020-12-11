/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/

/*
    {  "customerName" : "Nam" , "caseRef" : "123", "pieces" : ["Justificatif", "Identite"]}
*/
var feedback;

console.log("JS VM IN: " + input);
var obj = JSON.parse(input);


obj.template = "Bonjour <NAME>,<br><br>" +
    "Pour lancer l'étude de votre réclamation (n°<CASEREF>), il nous manque les documents suivants :<br>" +
    "<PIECES>" +
    "Merci de nous les faire parvenir, soit via votre espace client, soit en réponse à ce mail, ou encore en nous les apportant en agence.<br><br>" +
    "Merci par avance !<br><br>" +
    "Bien Cordialement <br>";

console.log("** input template : " + obj.template);
console.log("** input customerName : " + obj.customerName);
console.log("** input caseRef : " + obj.caseRef);
console.log("** input pieces : " + obj.docList);

var str =obj.template;

feedback = str.replace("<NAME>", obj.customerName);
feedback = feedback.replace("<CASEREF>", obj.caseRef);

var piecesStr = "";
for (var i=0; i< obj.docList; i++) {
    console.log("*********Adding doc : " + obj.docList[i]);
    piecesStr += "- " + obj.docList[i] + "<br>";
}

feedback = feedback.replace("<PIECES>", piecesStr);

console.log("** OUTPUT : " + feedback);

feedback;