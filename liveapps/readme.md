# TCI LiveApps Extension
TIBCO Cloud Integration Flogo custom Live Apps activity.

## Overview
These activites allow you to connect to any TIBCO Cloud Live Apps Subscription, and perform a growing number actions.
Have a look on the standard API documentation to get more details about the here used API calls.

> Note: the standard Live Apps Flogo Extension that comes with TIBCO Cloud Integration just allows to connect to the same Subscription Live Apps Applications.  

## Install URL
to install e.g. the LiveApps Activities use in the Flogo Extension Upload UI:

> github.com/TIBCOSoftware/TCSTK-GOlang/liveapps

## Source
All Source on GitHub here: [github.com/TIBCOSoftware/TCST-GOlang/liveapps](https://github.com/TIBCOSoftware/TCST-GOlang/liveapps)

## Activities
The first JavaScript Extension is the most Advanced Option here as it supports longer Scripts.

### Get Claims
provide details of a Subscription to perform actions.

### Start Case
create a new case instance in a Live Apps Application 

### Get Cases
query for existing cases of an Live Apps Subscription and return a single case, or a filtered list. 

### Get Artifact
download an asset from a Case folder e.g. an PDF

### Upload Artifact
upload an asset to a Case folder e.g. an PDF

## What's next
more to come soon, we have already over 20 more activities ready to release soon:

- getworkitem
- getworklistitems
- openworkitem
- completeworkitem
- putworkitem
- createstate
- deletecases
- deleteusergroupmappingsid
- getgroups
- getinstances
- getinstancestates
- getstates
- gettypes
- getusergroupmappings
- getusers
- postcaseactions
- postnotes
- postusergroupmappings
- putinstancestates
- runaction
- updatestates
