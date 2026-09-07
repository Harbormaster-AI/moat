import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ThirdParty} from '../models/ThirdParty';
import {SubrogationRecoveryService} from '../services/SubrogationRecovery.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ThirdPartyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	thirdParty : ThirdParty;

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
	// add a ThirdParty
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addThirdParty(name, taxId, address, Subrogations, PartyType) : Observable<any> {
		const uri_ = this.apiUrl + '/ThirdParty/create';
		const obj = {
			      		name: name,
      		taxId: taxId,
      		address: address,
      		Subrogations: Subrogations != null && Subrogations.length > 0 ? Subrogations : null,
			PartyType: PartyType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ThirdParty
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateThirdParty(name, taxId, address, Subrogations, PartyType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ThirdParty/update/' + id;
		const obj = {
				      		name: name,
      		taxId: taxId,
      		address: address,
      		Subrogations: Subrogations != null && Subrogations.length > 0 ? Subrogations : null,
			PartyType: PartyType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ThirdParty
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteThirdParty(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ThirdParty/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ThirdParty
	// returns the results untouched as an Observable ThirdParty
	// ThirdParty model
	// delegates via URI
	//********************************************************************
	getThirdParty(id) : Observable<ThirdParty> {
		const uri_ = this.apiUrl + '/ThirdParty/load/' + id;

		return this.http.get<ThirdParty>(uri_);
	}
	
	//********************************************************************
	// gets all ThirdParty
	// returns the results untouched as JSON representation of an
	// Observable array of ThirdParty models
	// delegates via URI
	//********************************************************************
	getThirdPartys() : Observable<ThirdParty[]> {
		const uri_ = this.apiUrl + '/ThirdParty/';

		return this
			.http.get<ThirdParty[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more subrogationsIds as a Subrogations
	// to a ThirdParty
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSubrogations( thirdPartyId, subrogationsIds ): Observable<any> {

		// get the ThirdParty
		this.loadHelper( thirdPartyId );

	// split on a comma with no spaces
	var idList = subrogationsIds.split(',')

	// iterate over array of subrogations ids
	idList.forEach(function (id) {
		// read the SubrogationRecovery
		var subrogationRecovery = new SubrogationRecoveryService(this.http).getSubrogationRecovery(id);
		// add the SubrogationRecovery if not already assigned
		if ( this.thirdParty.subrogations.indexOf(subrogationRecovery) == -1 )
		this.thirdParty.subrogations.push(subrogationRecovery);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more subrogationsIds as a Subrogations
	// from a ThirdParty
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSubrogations( thirdPartyId, subrogationsIds ): Observable<any> {

		// get the ThirdParty
		this.loadHelper( thirdPartyId );


	// split on a comma with no spaces
	var idList 					= subrogationsIds.split(',');
	var subrogations 	= this.thirdParty.subrogations;

	if ( subrogations != null && subrogationsIds != null ) {

		// iterate over array of subrogations ids
		subrogations.forEach(function (obj) {
			if ( subrogationsIds.indexOf(obj._id) > -1 ) {
				// remove the SubrogationRecovery
				this.thirdParty.subrogations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a ThirdParty
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ThirdParty/update/' + this.thirdParty;

	return  this.http.post(uri_, this.thirdParty );
}

	//********************************************************************
	// loadHelper - internal helper to load a ThirdParty
	//********************************************************************	
	loadHelper( id ) {
		this.getThirdParty(id)
			.subscribe((res : ThirdParty) => {
				this.thirdParty = res;
			});
	}
}