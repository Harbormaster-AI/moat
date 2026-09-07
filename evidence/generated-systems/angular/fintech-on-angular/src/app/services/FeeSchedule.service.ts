import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {FeeSchedule} from '../models/FeeSchedule';
import {PricingPlanService} from '../services/PricingPlan.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class FeeScheduleService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	feeSchedule : FeeSchedule;

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
	// add a FeeSchedule
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addFeeSchedule(name, amount, percentage, minimum, maximum, PricingPlan, FeeType, CalculationMethod) : Observable<any> {
		const uri_ = this.apiUrl + '/FeeSchedule/create';
		const obj = {
			      		name: name,
      		amount: amount,
      		percentage: percentage,
      		minimum: minimum,
      		maximum: maximum,
      		PricingPlan: PricingPlan != null && PricingPlan.length > 0 ? PricingPlan : null,
      		FeeType: FeeType,
			CalculationMethod: CalculationMethod
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a FeeSchedule
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateFeeSchedule(name, amount, percentage, minimum, maximum, PricingPlan, FeeType, CalculationMethod, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/FeeSchedule/update/' + id;
		const obj = {
				      		name: name,
      		amount: amount,
      		percentage: percentage,
      		minimum: minimum,
      		maximum: maximum,
      		PricingPlan: PricingPlan != null && PricingPlan.length > 0 ? PricingPlan : null,
      		FeeType: FeeType,
			CalculationMethod: CalculationMethod
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a FeeSchedule
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteFeeSchedule(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/FeeSchedule/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a FeeSchedule
	// returns the results untouched as an Observable FeeSchedule
	// FeeSchedule model
	// delegates via URI
	//********************************************************************
	getFeeSchedule(id) : Observable<FeeSchedule> {
		const uri_ = this.apiUrl + '/FeeSchedule/load/' + id;

		return this.http.get<FeeSchedule>(uri_);
	}
	
	//********************************************************************
	// gets all FeeSchedule
	// returns the results untouched as JSON representation of an
	// Observable array of FeeSchedule models
	// delegates via URI
	//********************************************************************
	getFeeSchedules() : Observable<FeeSchedule[]> {
		const uri_ = this.apiUrl + '/FeeSchedule/';

		return this
			.http.get<FeeSchedule[]>(uri_);
	}
	
			//********************************************************************
	// assigns a PricingPlan on a FeeSchedule
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPricingPlan( feeScheduleId, _pricingPlanId ): Observable<any> {

		// get the FeeSchedule from storage
		this.loadHelper( feeScheduleId );

	// get the PricingPlan from storage
	var tmp 	= new PricingPlanService(this.http).getPricingPlan(_pricingPlanId);

	// assign the PricingPlan
	this.feeSchedule.pricingPlan = tmp;

	// save the FeeSchedule
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PricingPlan on a FeeSchedule
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPricingPlan( feeScheduleId ): Observable<any> {

		// get the FeeSchedule from storage
		this.loadHelper( feeScheduleId );

	// assign PricingPlan to null
	this.feeSchedule.pricingPlan = null;

	// save the FeeSchedule
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a FeeSchedule
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/FeeSchedule/update/' + this.feeSchedule;

	return  this.http.post(uri_, this.feeSchedule );
}

	//********************************************************************
	// loadHelper - internal helper to load a FeeSchedule
	//********************************************************************	
	loadHelper( id ) {
		this.getFeeSchedule(id)
			.subscribe((res : FeeSchedule) => {
				this.feeSchedule = res;
			});
	}
}