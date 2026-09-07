import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PricingPlan} from '../models/PricingPlan';
import {ProductOfferingService} from '../services/ProductOffering.service';
import {FeeScheduleService} from '../services/FeeSchedule.service';
import {UsageLimitService} from '../services/UsageLimit.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PricingPlanService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	pricingPlan : PricingPlan;

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
	// add a PricingPlan
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPricingPlan(name, planCode, baseCurrency, ProductOffering, FeeSchedules, Limits, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/PricingPlan/create';
		const obj = {
			      		name: name,
      		planCode: planCode,
      		baseCurrency: baseCurrency,
      		ProductOffering: ProductOffering != null && ProductOffering.length > 0 ? ProductOffering : null,
      		FeeSchedules: FeeSchedules != null && FeeSchedules.length > 0 ? FeeSchedules : null,
      		Limits: Limits != null && Limits.length > 0 ? Limits : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PricingPlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePricingPlan(name, planCode, baseCurrency, ProductOffering, FeeSchedules, Limits, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PricingPlan/update/' + id;
		const obj = {
				      		name: name,
      		planCode: planCode,
      		baseCurrency: baseCurrency,
      		ProductOffering: ProductOffering != null && ProductOffering.length > 0 ? ProductOffering : null,
      		FeeSchedules: FeeSchedules != null && FeeSchedules.length > 0 ? FeeSchedules : null,
      		Limits: Limits != null && Limits.length > 0 ? Limits : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PricingPlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePricingPlan(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PricingPlan/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PricingPlan
	// returns the results untouched as an Observable PricingPlan
	// PricingPlan model
	// delegates via URI
	//********************************************************************
	getPricingPlan(id) : Observable<PricingPlan> {
		const uri_ = this.apiUrl + '/PricingPlan/load/' + id;

		return this.http.get<PricingPlan>(uri_);
	}
	
	//********************************************************************
	// gets all PricingPlan
	// returns the results untouched as JSON representation of an
	// Observable array of PricingPlan models
	// delegates via URI
	//********************************************************************
	getPricingPlans() : Observable<PricingPlan[]> {
		const uri_ = this.apiUrl + '/PricingPlan/';

		return this
			.http.get<PricingPlan[]>(uri_);
	}
	
			//********************************************************************
	// assigns a ProductOffering on a PricingPlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignProductOffering( pricingPlanId, _productOfferingId ): Observable<any> {

		// get the PricingPlan from storage
		this.loadHelper( pricingPlanId );

	// get the ProductOffering from storage
	var tmp 	= new ProductOfferingService(this.http).getProductOffering(_productOfferingId);

	// assign the ProductOffering
	this.pricingPlan.productOffering = tmp;

	// save the PricingPlan
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ProductOffering on a PricingPlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignProductOffering( pricingPlanId ): Observable<any> {

		// get the PricingPlan from storage
		this.loadHelper( pricingPlanId );

	// assign ProductOffering to null
	this.pricingPlan.productOffering = null;

	// save the PricingPlan
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more feeSchedulesIds as a FeeSchedules
	// to a PricingPlan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addFeeSchedules( pricingPlanId, feeSchedulesIds ): Observable<any> {

		// get the PricingPlan
		this.loadHelper( pricingPlanId );

	// split on a comma with no spaces
	var idList = feeSchedulesIds.split(',')

	// iterate over array of feeSchedules ids
	idList.forEach(function (id) {
		// read the FeeSchedule
		var feeSchedule = new FeeScheduleService(this.http).getFeeSchedule(id);
		// add the FeeSchedule if not already assigned
		if ( this.pricingPlan.feeSchedules.indexOf(feeSchedule) == -1 )
		this.pricingPlan.feeSchedules.push(feeSchedule);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more feeSchedulesIds as a FeeSchedules
	// from a PricingPlan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeFeeSchedules( pricingPlanId, feeSchedulesIds ): Observable<any> {

		// get the PricingPlan
		this.loadHelper( pricingPlanId );


	// split on a comma with no spaces
	var idList 					= feeSchedulesIds.split(',');
	var feeSchedules 	= this.pricingPlan.feeSchedules;

	if ( feeSchedules != null && feeSchedulesIds != null ) {

		// iterate over array of feeSchedules ids
		feeSchedules.forEach(function (obj) {
			if ( feeSchedulesIds.indexOf(obj._id) > -1 ) {
				// remove the FeeSchedule
				this.pricingPlan.feeSchedules.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more limitsIds as a Limits
	// to a PricingPlan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLimits( pricingPlanId, limitsIds ): Observable<any> {

		// get the PricingPlan
		this.loadHelper( pricingPlanId );

	// split on a comma with no spaces
	var idList = limitsIds.split(',')

	// iterate over array of limits ids
	idList.forEach(function (id) {
		// read the UsageLimit
		var usageLimit = new UsageLimitService(this.http).getUsageLimit(id);
		// add the UsageLimit if not already assigned
		if ( this.pricingPlan.limits.indexOf(usageLimit) == -1 )
		this.pricingPlan.limits.push(usageLimit);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more limitsIds as a Limits
	// from a PricingPlan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLimits( pricingPlanId, limitsIds ): Observable<any> {

		// get the PricingPlan
		this.loadHelper( pricingPlanId );


	// split on a comma with no spaces
	var idList 					= limitsIds.split(',');
	var limits 	= this.pricingPlan.limits;

	if ( limits != null && limitsIds != null ) {

		// iterate over array of limits ids
		limits.forEach(function (obj) {
			if ( limitsIds.indexOf(obj._id) > -1 ) {
				// remove the UsageLimit
				this.pricingPlan.limits.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a PricingPlan
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PricingPlan/update/' + this.pricingPlan;

	return  this.http.post(uri_, this.pricingPlan );
}

	//********************************************************************
	// loadHelper - internal helper to load a PricingPlan
	//********************************************************************	
	loadHelper( id ) {
		this.getPricingPlan(id)
			.subscribe((res : PricingPlan) => {
				this.pricingPlan = res;
			});
	}
}