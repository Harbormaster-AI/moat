import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {UsageLimit} from '../models/UsageLimit';
import {PricingPlanService} from '../services/PricingPlan.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class UsageLimitService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	usageLimit : UsageLimit;

	//********************************************************************
	// Catch all for the return value of a service call
	//********************************************************************
	result: any;

	//********************************************************************
	// sole constructor, injected with the HttpClient
	//********************************************************************
	constructor(private http: HttpClient) {
		super();
	}

		//********************************************************************
	// add a UsageLimit
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addUsageLimit(name, amount, count, PricingPlan, Scope, Period) : Observable<any> {
		const uri_ = this.apiUrl + '/UsageLimit/create';
		const obj = {
			      		name: name,
      		amount: amount,
      		count: count,
      		PricingPlan: PricingPlan != null && PricingPlan.length > 0 ? PricingPlan : null,
      		Scope: Scope,
			Period: Period
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a UsageLimit
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateUsageLimit(name, amount, count, PricingPlan, Scope, Period, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/UsageLimit/update/' + id;
		const obj = {
				      		name: name,
      		amount: amount,
      		count: count,
      		PricingPlan: PricingPlan != null && PricingPlan.length > 0 ? PricingPlan : null,
      		Scope: Scope,
			Period: Period
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a UsageLimit
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteUsageLimit(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/UsageLimit/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a UsageLimit
	// returns the results untouched as an Observable UsageLimit
	// UsageLimit model
	// delegates via URI
	//********************************************************************
	getUsageLimit(id) : Observable<UsageLimit> {
		const uri_ = this.apiUrl + '/UsageLimit/load/' + id;

		return this.http.get<UsageLimit>(uri_);
	}
	
	//********************************************************************
	// gets all UsageLimit
	// returns the results untouched as JSON representation of an
	// Observable array of UsageLimit models
	// delegates via URI
	//********************************************************************
	getUsageLimits() : Observable<UsageLimit[]> {
		const uri_ = this.apiUrl + '/UsageLimit/';

		return this
			.http.get<UsageLimit[]>(uri_);
	}
	
			//********************************************************************
	// assigns a PricingPlan on a UsageLimit
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPricingPlan( usageLimitId, _pricingPlanId ): Observable<any> {

		// get the UsageLimit from storage
		this.loadHelper( usageLimitId );

	// get the PricingPlan from storage
	var tmp 	= new PricingPlanService(this.http).getPricingPlan(_pricingPlanId);

	// assign the PricingPlan
	this.usageLimit.pricingPlan = tmp;

	// save the UsageLimit
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PricingPlan on a UsageLimit
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPricingPlan( usageLimitId ): Observable<any> {

		// get the UsageLimit from storage
		this.loadHelper( usageLimitId );

	// assign PricingPlan to null
	this.usageLimit.pricingPlan = null;

	// save the UsageLimit
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a UsageLimit
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/UsageLimit/update/' + this.usageLimit;

	return  this.http.post(uri_, this.usageLimit );
}

	//********************************************************************
	// loadHelper - internal helper to load a UsageLimit
	//********************************************************************	
	loadHelper( id ) {
		this.getUsageLimit(id)
			.subscribe((res : UsageLimit) => {
				this.usageLimit = res;
			});
	}
}