import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {SubrogationRecovery} from '../models/SubrogationRecovery';
import {ClaimService} from '../services/Claim.service';
import {ExposureService} from '../services/Exposure.service';
import {ThirdPartyService} from '../services/ThirdParty.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class SubrogationRecoveryService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	subrogationRecovery : SubrogationRecovery;

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
	// add a SubrogationRecovery
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addSubrogationRecovery(recoveryReference, amount, recoveryDate, Claim, Exposure, Counterparty, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/SubrogationRecovery/create';
		const obj = {
			      		recoveryReference: recoveryReference,
      		amount: amount,
      		recoveryDate: recoveryDate,
      		Claim: Claim != null && Claim.length > 0 ? Claim : null,
      		Exposure: Exposure != null && Exposure.length > 0 ? Exposure : null,
      		Counterparty: Counterparty != null && Counterparty.length > 0 ? Counterparty : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a SubrogationRecovery
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateSubrogationRecovery(recoveryReference, amount, recoveryDate, Claim, Exposure, Counterparty, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/SubrogationRecovery/update/' + id;
		const obj = {
				      		recoveryReference: recoveryReference,
      		amount: amount,
      		recoveryDate: recoveryDate,
      		Claim: Claim != null && Claim.length > 0 ? Claim : null,
      		Exposure: Exposure != null && Exposure.length > 0 ? Exposure : null,
      		Counterparty: Counterparty != null && Counterparty.length > 0 ? Counterparty : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a SubrogationRecovery
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteSubrogationRecovery(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/SubrogationRecovery/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a SubrogationRecovery
	// returns the results untouched as an Observable SubrogationRecovery
	// SubrogationRecovery model
	// delegates via URI
	//********************************************************************
	getSubrogationRecovery(id) : Observable<SubrogationRecovery> {
		const uri_ = this.apiUrl + '/SubrogationRecovery/load/' + id;

		return this.http.get<SubrogationRecovery>(uri_);
	}
	
	//********************************************************************
	// gets all SubrogationRecovery
	// returns the results untouched as JSON representation of an
	// Observable array of SubrogationRecovery models
	// delegates via URI
	//********************************************************************
	getSubrogationRecoverys() : Observable<SubrogationRecovery[]> {
		const uri_ = this.apiUrl + '/SubrogationRecovery/';

		return this
			.http.get<SubrogationRecovery[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Claim on a SubrogationRecovery
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignClaim( subrogationRecoveryId, _claimId ): Observable<any> {

		// get the SubrogationRecovery from storage
		this.loadHelper( subrogationRecoveryId );

	// get the Claim from storage
	var tmp 	= new ClaimService(this.http).getClaim(_claimId);

	// assign the Claim
	this.subrogationRecovery.claim = tmp;

	// save the SubrogationRecovery
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Claim on a SubrogationRecovery
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignClaim( subrogationRecoveryId ): Observable<any> {

		// get the SubrogationRecovery from storage
		this.loadHelper( subrogationRecoveryId );

	// assign Claim to null
	this.subrogationRecovery.claim = null;

	// save the SubrogationRecovery
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Exposure on a SubrogationRecovery
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignExposure( subrogationRecoveryId, _exposureId ): Observable<any> {

		// get the SubrogationRecovery from storage
		this.loadHelper( subrogationRecoveryId );

	// get the Exposure from storage
	var tmp 	= new ExposureService(this.http).getExposure(_exposureId);

	// assign the Exposure
	this.subrogationRecovery.exposure = tmp;

	// save the SubrogationRecovery
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Exposure on a SubrogationRecovery
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignExposure( subrogationRecoveryId ): Observable<any> {

		// get the SubrogationRecovery from storage
		this.loadHelper( subrogationRecoveryId );

	// assign Exposure to null
	this.subrogationRecovery.exposure = null;

	// save the SubrogationRecovery
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Counterparty on a SubrogationRecovery
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCounterparty( subrogationRecoveryId, _counterpartyId ): Observable<any> {

		// get the SubrogationRecovery from storage
		this.loadHelper( subrogationRecoveryId );

	// get the ThirdParty from storage
	var tmp 	= new ThirdPartyService(this.http).getThirdParty(_counterpartyId);

	// assign the Counterparty
	this.subrogationRecovery.counterparty = tmp;

	// save the SubrogationRecovery
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Counterparty on a SubrogationRecovery
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCounterparty( subrogationRecoveryId ): Observable<any> {

		// get the SubrogationRecovery from storage
		this.loadHelper( subrogationRecoveryId );

	// assign Counterparty to null
	this.subrogationRecovery.counterparty = null;

	// save the SubrogationRecovery
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a SubrogationRecovery
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/SubrogationRecovery/update/' + this.subrogationRecovery;

	return  this.http.post(uri_, this.subrogationRecovery );
}

	//********************************************************************
	// loadHelper - internal helper to load a SubrogationRecovery
	//********************************************************************	
	loadHelper( id ) {
		this.getSubrogationRecovery(id)
			.subscribe((res : SubrogationRecovery) => {
				this.subrogationRecovery = res;
			});
	}
}