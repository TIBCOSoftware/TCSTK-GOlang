import { WiServiceContribution } from "wi-studio/app/contrib/wi-contrib";
import { NgModule } from "@angular/core";
import { CommonModule } from "@angular/common";
import { HttpModule } from "@angular/http";
import { LiveAppsGetArtifactActivityContribution } from "./activity";

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
            useClass: LiveAppsGetArtifactActivityContribution
        }
    ],
    bootstrap: []
})

export default class LiveAppsGetArtifactActivityModule {

}
