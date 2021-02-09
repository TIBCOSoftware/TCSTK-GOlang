/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/
import { Observable } from "rxjs/Observable";
import { Inject, Injectable, Injector } from "@angular/core";
import { Http } from "@angular/http";
import {
    IValidationResult,
    WiContrib,
    WiServiceHandlerContribution,
    IActivityContribution,
    ValidationResult,
    IFieldDefinition
} from "wi-studio/app/contrib/wi-contrib";

@WiContrib({})
@Injectable()

export class LiveAppsRunActionActivityContribution extends WiServiceHandlerContribution {
    constructor(@Inject(Injector) injector, private http: Http) {
        super(injector, http);
    }

    /**
     * The value object allows you to specify what types of values you can pick for a certain field
     */
    value = (fieldName: string, context: IActivityContribution): Observable<any> | any => {
        return null;
    }

    /**
     * The validate object can be used to validate the input of certain fields
     */
    validate = (fieldName: string, context: IActivityContribution): Observable<IValidationResult> | IValidationResult => {
        if (fieldName === 'cookie' || fieldName === 'accessToken') {
            let vResult: IValidationResult = ValidationResult.newValidationResult();
            let security: IFieldDefinition = context.getField("security");

            if (security.value.toLowerCase() === fieldName.toLowerCase()){
                vResult.setVisible(true);
            } else {
                vResult.setVisible(false);
            }
            return vResult;
        }
        return null;
    }
}
