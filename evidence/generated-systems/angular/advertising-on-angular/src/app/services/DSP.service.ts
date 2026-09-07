import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {DSP} from '../models/DSP';
import {AdAccountService} from '../services/AdAccount.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DSPService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	dSP : DSP;

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
	// add a DSP
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDSP(name, website, region, AdAccounts) : Observable<any> {
		const uri_ = this.apiUrl + '/DSP/create';
		const obj = {
			      		name: name,
      		website: website,
      		region: region,
			AdAccounts: AdAccounts != null && AdAccounts.length > 0 ? AdAccounts : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a DSP
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDSP(name, website, region, AdAccounts, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/DSP/update/' + id;
		const obj = {
				      		name: name,
      		website: website,
      		region: region,
			AdAccounts: AdAccounts != null && AdAccounts.length > 0 ? AdAccounts : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a DSP
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDSP(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/DSP/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a DSP
	// returns the results untouched as an Observable DSP
	// DSP model
	// delegates via URI
	//********************************************************************
	getDSP(id) : Observable<DSP> {
		const uri_ = this.apiUrl + '/DSP/load/' + id;

		return this.http.get<DSP>(uri_);
	}
	
	//********************************************************************
	// gets all DSP
	// returns the results untouched as JSON representation of an
	// Observable array of DSP models
	// delegates via URI
	//********************************************************************
	getDSPs() : Observable<DSP[]> {
		const uri_ = this.apiUrl + '/DSP/';

		return this
			.http.get<DSP[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more adAccountsIds as a AdAccounts
	// to a DSP
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAdAccounts( dSPId, adAccountsIds ): Observable<any> {

		// get the DSP
		this.loadHelper( dSPId );

	// split on a comma with no spaces
	var idList = adAccountsIds.split(',')

	// iterate over array of adAccounts ids
	idList.forEach(function (id) {
		// read the AdAccount
		var adAccount = new AdAccountService(this.http).getAdAccount(id);
		// add the AdAccount if not already assigned
		if ( this.dSP.adAccounts.indexOf(adAccount) == -1 )
		this.dSP.adAccounts.push(adAccount);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more adAccountsIds as a AdAccounts
	// from a DSP
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAdAccounts( dSPId, adAccountsIds ): Observable<any> {

		// get the DSP
		this.loadHelper( dSPId );


	// split on a comma with no spaces
	var idList 					= adAccountsIds.split(',');
	var adAccounts 	= this.dSP.adAccounts;

	if ( adAccounts != null && adAccountsIds != null ) {

		// iterate over array of adAccounts ids
		adAccounts.forEach(function (obj) {
			if ( adAccountsIds.indexOf(obj._id) > -1 ) {
				// remove the AdAccount
				this.dSP.adAccounts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a DSP
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/DSP/update/' + this.dSP;

	return  this.http.post(uri_, this.dSP );
}

	//********************************************************************
	// loadHelper - internal helper to load a DSP
	//********************************************************************	
	loadHelper( id ) {
		this.getDSP(id)
			.subscribe((res : DSP) => {
				this.dSP = res;
			});
	}
}