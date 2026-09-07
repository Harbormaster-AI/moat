import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ProductOffering} from '../models/ProductOffering';
import {FinancialInstitutionService} from '../services/FinancialInstitution.service';
import {PricingPlanService} from '../services/PricingPlan.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ProductOfferingService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	productOffering : ProductOffering;

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
	// add a ProductOffering
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addProductOffering(name, productCode, Institution, PricingPlans, Category) : Observable<any> {
		const uri_ = this.apiUrl + '/ProductOffering/create';
		const obj = {
			      		name: name,
      		productCode: productCode,
      		Institution: Institution != null && Institution.length > 0 ? Institution : null,
      		PricingPlans: PricingPlans != null && PricingPlans.length > 0 ? PricingPlans : null,
			Category: Category
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ProductOffering
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateProductOffering(name, productCode, Institution, PricingPlans, Category, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ProductOffering/update/' + id;
		const obj = {
				      		name: name,
      		productCode: productCode,
      		Institution: Institution != null && Institution.length > 0 ? Institution : null,
      		PricingPlans: PricingPlans != null && PricingPlans.length > 0 ? PricingPlans : null,
			Category: Category
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ProductOffering
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteProductOffering(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ProductOffering/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ProductOffering
	// returns the results untouched as an Observable ProductOffering
	// ProductOffering model
	// delegates via URI
	//********************************************************************
	getProductOffering(id) : Observable<ProductOffering> {
		const uri_ = this.apiUrl + '/ProductOffering/load/' + id;

		return this.http.get<ProductOffering>(uri_);
	}
	
	//********************************************************************
	// gets all ProductOffering
	// returns the results untouched as JSON representation of an
	// Observable array of ProductOffering models
	// delegates via URI
	//********************************************************************
	getProductOfferings() : Observable<ProductOffering[]> {
		const uri_ = this.apiUrl + '/ProductOffering/';

		return this
			.http.get<ProductOffering[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Institution on a ProductOffering
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInstitution( productOfferingId, _institutionId ): Observable<any> {

		// get the ProductOffering from storage
		this.loadHelper( productOfferingId );

	// get the FinancialInstitution from storage
	var tmp 	= new FinancialInstitutionService(this.http).getFinancialInstitution(_institutionId);

	// assign the Institution
	this.productOffering.institution = tmp;

	// save the ProductOffering
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Institution on a ProductOffering
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInstitution( productOfferingId ): Observable<any> {

		// get the ProductOffering from storage
		this.loadHelper( productOfferingId );

	// assign Institution to null
	this.productOffering.institution = null;

	// save the ProductOffering
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more pricingPlansIds as a PricingPlans
	// to a ProductOffering
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPricingPlans( productOfferingId, pricingPlansIds ): Observable<any> {

		// get the ProductOffering
		this.loadHelper( productOfferingId );

	// split on a comma with no spaces
	var idList = pricingPlansIds.split(',')

	// iterate over array of pricingPlans ids
	idList.forEach(function (id) {
		// read the PricingPlan
		var pricingPlan = new PricingPlanService(this.http).getPricingPlan(id);
		// add the PricingPlan if not already assigned
		if ( this.productOffering.pricingPlans.indexOf(pricingPlan) == -1 )
		this.productOffering.pricingPlans.push(pricingPlan);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more pricingPlansIds as a PricingPlans
	// from a ProductOffering
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePricingPlans( productOfferingId, pricingPlansIds ): Observable<any> {

		// get the ProductOffering
		this.loadHelper( productOfferingId );


	// split on a comma with no spaces
	var idList 					= pricingPlansIds.split(',');
	var pricingPlans 	= this.productOffering.pricingPlans;

	if ( pricingPlans != null && pricingPlansIds != null ) {

		// iterate over array of pricingPlans ids
		pricingPlans.forEach(function (obj) {
			if ( pricingPlansIds.indexOf(obj._id) > -1 ) {
				// remove the PricingPlan
				this.productOffering.pricingPlans.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a ProductOffering
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ProductOffering/update/' + this.productOffering;

	return  this.http.post(uri_, this.productOffering );
}

	//********************************************************************
	// loadHelper - internal helper to load a ProductOffering
	//********************************************************************	
	loadHelper( id ) {
		this.getProductOffering(id)
			.subscribe((res : ProductOffering) => {
				this.productOffering = res;
			});
	}
}