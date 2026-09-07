import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {SalesCampaign} from '../models/SalesCampaign';
import {SalesRegionService} from '../services/SalesRegion.service';
import {OperatorService} from '../services/Operator.service';
import {QuoteService} from '../services/Quote.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class SalesCampaignService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	salesCampaign : SalesCampaign;

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
	// add a SalesCampaign
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addSalesCampaign(campaignCode, Region, Operator, Quotes, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/SalesCampaign/create';
		const obj = {
			      		campaignCode: campaignCode,
      		Region: Region != null && Region.length > 0 ? Region : null,
      		Operator: Operator != null && Operator.length > 0 ? Operator : null,
      		Quotes: Quotes != null && Quotes.length > 0 ? Quotes : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a SalesCampaign
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateSalesCampaign(campaignCode, Region, Operator, Quotes, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/SalesCampaign/update/' + id;
		const obj = {
				      		campaignCode: campaignCode,
      		Region: Region != null && Region.length > 0 ? Region : null,
      		Operator: Operator != null && Operator.length > 0 ? Operator : null,
      		Quotes: Quotes != null && Quotes.length > 0 ? Quotes : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a SalesCampaign
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteSalesCampaign(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/SalesCampaign/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a SalesCampaign
	// returns the results untouched as an Observable SalesCampaign
	// SalesCampaign model
	// delegates via URI
	//********************************************************************
	getSalesCampaign(id) : Observable<SalesCampaign> {
		const uri_ = this.apiUrl + '/SalesCampaign/load/' + id;

		return this.http.get<SalesCampaign>(uri_);
	}
	
	//********************************************************************
	// gets all SalesCampaign
	// returns the results untouched as JSON representation of an
	// Observable array of SalesCampaign models
	// delegates via URI
	//********************************************************************
	getSalesCampaigns() : Observable<SalesCampaign[]> {
		const uri_ = this.apiUrl + '/SalesCampaign/';

		return this
			.http.get<SalesCampaign[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Region on a SalesCampaign
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRegion( salesCampaignId, _regionId ): Observable<any> {

		// get the SalesCampaign from storage
		this.loadHelper( salesCampaignId );

	// get the SalesRegion from storage
	var tmp 	= new SalesRegionService(this.http).getSalesRegion(_regionId);

	// assign the Region
	this.salesCampaign.region = tmp;

	// save the SalesCampaign
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Region on a SalesCampaign
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRegion( salesCampaignId ): Observable<any> {

		// get the SalesCampaign from storage
		this.loadHelper( salesCampaignId );

	// assign Region to null
	this.salesCampaign.region = null;

	// save the SalesCampaign
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Operator on a SalesCampaign
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOperator( salesCampaignId, _operatorId ): Observable<any> {

		// get the SalesCampaign from storage
		this.loadHelper( salesCampaignId );

	// get the Operator from storage
	var tmp 	= new OperatorService(this.http).getOperator(_operatorId);

	// assign the Operator
	this.salesCampaign.operator = tmp;

	// save the SalesCampaign
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Operator on a SalesCampaign
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOperator( salesCampaignId ): Observable<any> {

		// get the SalesCampaign from storage
		this.loadHelper( salesCampaignId );

	// assign Operator to null
	this.salesCampaign.operator = null;

	// save the SalesCampaign
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more quotesIds as a Quotes
	// to a SalesCampaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addQuotes( salesCampaignId, quotesIds ): Observable<any> {

		// get the SalesCampaign
		this.loadHelper( salesCampaignId );

	// split on a comma with no spaces
	var idList = quotesIds.split(',')

	// iterate over array of quotes ids
	idList.forEach(function (id) {
		// read the Quote
		var quote = new QuoteService(this.http).getQuote(id);
		// add the Quote if not already assigned
		if ( this.salesCampaign.quotes.indexOf(quote) == -1 )
		this.salesCampaign.quotes.push(quote);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more quotesIds as a Quotes
	// from a SalesCampaign
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeQuotes( salesCampaignId, quotesIds ): Observable<any> {

		// get the SalesCampaign
		this.loadHelper( salesCampaignId );


	// split on a comma with no spaces
	var idList 					= quotesIds.split(',');
	var quotes 	= this.salesCampaign.quotes;

	if ( quotes != null && quotesIds != null ) {

		// iterate over array of quotes ids
		quotes.forEach(function (obj) {
			if ( quotesIds.indexOf(obj._id) > -1 ) {
				// remove the Quote
				this.salesCampaign.quotes.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a SalesCampaign
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/SalesCampaign/update/' + this.salesCampaign;

	return  this.http.post(uri_, this.salesCampaign );
}

	//********************************************************************
	// loadHelper - internal helper to load a SalesCampaign
	//********************************************************************	
	loadHelper( id ) {
		this.getSalesCampaign(id)
			.subscribe((res : SalesCampaign) => {
				this.salesCampaign = res;
			});
	}
}