/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/
import { WiServiceContribution } from "wi-studio/app/contrib/wi-contrib";
import { NgModule } from "@angular/core";
import { CommonModule } from "@angular/common";
import { HttpModule } from "@angular/http";
import { LiveAppsGetTypesActivityContribution } from "./activity";

@NgModule({
    imports: [
        CommonModule,
        HttpModule
    ],
    exports: [],
    declarations: [],
    entryComponents: [],
    providers: [
        {
            provide: WiServiceContribution,
            useClass: LiveAppsGetTypesActivityContribution
        }
    ],
    bootstrap: []
})

export default class LiveAppsGetTypesActivityModule {

}
