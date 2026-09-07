import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {SalesRegion} from '../models/SalesRegion';
import {OperatorService} from '../services/Operator.service';
import {SalesCampaignService} from '../services/SalesCampaign.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class SalesRegionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	salesRegion : SalesRegion;

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
	// add a SalesRegion
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addSalesRegion(name, regionCode, Operators, SalesCampaigns) : Observable<any> {
		const uri_ = this.apiUrl + '/SalesRegion/create';
		const obj = {
			      		name: name,
      		regionCode: regionCode,
      		Operators: Operators != null && Operators.length > 0 ? Operators : null,
			SalesCampaigns: SalesCampaigns != null && SalesCampaigns.length > 0 ? SalesCampaigns : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a SalesRegion
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateSalesRegion(name, regionCode, Operators, SalesCampaigns, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/SalesRegion/update/' + id;
		const obj = {
				      		name: name,
      		regionCode: regionCode,
      		Operators: Operators != null && Operators.length > 0 ? Operators : null,
			SalesCampaigns: SalesCampaigns != null && SalesCampaigns.length > 0 ? SalesCampaigns : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a SalesRegion
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteSalesRegion(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/SalesRegion/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a SalesRegion
	// returns the results untouched as an Observable SalesRegion
	// SalesRegion model
	// delegates via URI
	//********************************************************************
	getSalesRegion(id) : Observable<SalesRegion> {
		const uri_ = this.apiUrl + '/SalesRegion/load/' + id;

		return this.http.get<SalesRegion>(uri_);
	}
	
	//********************************************************************
	// gets all SalesRegion
	// returns the results untouched as JSON representation of an
	// Observable array of SalesRegion models
	// delegates via URI
	//********************************************************************
	getSalesRegions() : Observable<SalesRegion[]> {
		const uri_ = this.apiUrl + '/SalesRegion/';

		return this
			.http.get<SalesRegion[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more operatorsIds as a Operators
	// to a SalesRegion
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOperators( salesRegionId, operatorsIds ): Observable<any> {

		// get the SalesRegion
		this.loadHelper( salesRegionId );

	// split on a comma with no spaces
	var idList = operatorsIds.split(',')

	// iterate over array of operators ids
	idList.forEach(function (id) {
		// read the Operator
		var operator = new OperatorService(this.http).getOperator(id);
		// add the Operator if not already assigned
		if ( this.salesRegion.operators.indexOf(operator) == -1 )
		this.salesRegion.operators.push(operator);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more operatorsIds as a Operators
	// from a SalesRegion
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOperators( salesRegionId, operatorsIds ): Observable<any> {

		// get the SalesRegion
		this.loadHelper( salesRegionId );


	// split on a comma with no spaces
	var idList 					= operatorsIds.split(',');
	var operators 	= this.salesRegion.operators;

	if ( operators != null && operatorsIds != null ) {

		// iterate over array of operators ids
		operators.forEach(function (obj) {
			if ( operatorsIds.indexOf(obj._id) > -1 ) {
				// remove the Operator
				this.salesRegion.operators.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more salesCampaignsIds as a SalesCampaigns
	// to a SalesRegion
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSalesCampaigns( salesRegionId, salesCampaignsIds ): Observable<any> {

		// get the SalesRegion
		this.loadHelper( salesRegionId );

	// split on a comma with no spaces
	var idList = salesCampaignsIds.split(',')

	// iterate over array of salesCampaigns ids
	idList.forEach(function (id) {
		// read the SalesCampaign
		var salesCampaign = new SalesCampaignService(this.http).getSalesCampaign(id);
		// add the SalesCampaign if not already assigned
		if ( this.salesRegion.salesCampaigns.indexOf(salesCampaign) == -1 )
		this.salesRegion.salesCampaigns.push(salesCampaign);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more salesCampaignsIds as a SalesCampaigns
	// from a SalesRegion
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSalesCampaigns( salesRegionId, salesCampaignsIds ): Observable<any> {

		// get the SalesRegion
		this.loadHelper( salesRegionId );


	// split on a comma with no spaces
	var idList 					= salesCampaignsIds.split(',');
	var salesCampaigns 	= this.salesRegion.salesCampaigns;

	if ( salesCampaigns != null && salesCampaignsIds != null ) {

		// iterate over array of salesCampaigns ids
		salesCampaigns.forEach(function (obj) {
			if ( salesCampaignsIds.indexOf(obj._id) > -1 ) {
				// remove the SalesCampaign
				this.salesRegion.salesCampaigns.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a SalesRegion
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/SalesRegion/update/' + this.salesRegion;

	return  this.http.post(uri_, this.salesRegion );
}

	//********************************************************************
	// loadHelper - internal helper to load a SalesRegion
	//********************************************************************	
	loadHelper( id ) {
		this.getSalesRegion(id)
			.subscribe((res : SalesRegion) => {
				this.salesRegion = res;
			});
	}
}