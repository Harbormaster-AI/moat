import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Exposure} from '../models/Exposure';
import {ClaimService} from '../services/Claim.service';
import {PolicyCoverageService} from '../services/PolicyCoverage.service';
import {InsuredObjectService} from '../services/InsuredObject.service';
import {ClaimReserveService} from '../services/ClaimReserve.service';
import {ClaimPaymentService} from '../services/ClaimPayment.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ExposureService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	exposure : Exposure;

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
	// add a Exposure
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addExposure(Claim, PolicyCoverage, InsuredObject, Reserves, Payments, ExposureType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Exposure/create';
		const obj = {
			      		Claim: Claim != null && Claim.length > 0 ? Claim : null,
      		PolicyCoverage: PolicyCoverage != null && PolicyCoverage.length > 0 ? PolicyCoverage : null,
      		InsuredObject: InsuredObject != null && InsuredObject.length > 0 ? InsuredObject : null,
      		Reserves: Reserves != null && Reserves.length > 0 ? Reserves : null,
      		Payments: Payments != null && Payments.length > 0 ? Payments : null,
      		ExposureType: ExposureType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Exposure
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateExposure(Claim, PolicyCoverage, InsuredObject, Reserves, Payments, ExposureType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Exposure/update/' + id;
		const obj = {
				      		Claim: Claim != null && Claim.length > 0 ? Claim : null,
      		PolicyCoverage: PolicyCoverage != null && PolicyCoverage.length > 0 ? PolicyCoverage : null,
      		InsuredObject: InsuredObject != null && InsuredObject.length > 0 ? InsuredObject : null,
      		Reserves: Reserves != null && Reserves.length > 0 ? Reserves : null,
      		Payments: Payments != null && Payments.length > 0 ? Payments : null,
      		ExposureType: ExposureType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Exposure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteExposure(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Exposure/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Exposure
	// returns the results untouched as an Observable Exposure
	// Exposure model
	// delegates via URI
	//********************************************************************
	getExposure(id) : Observable<Exposure> {
		const uri_ = this.apiUrl + '/Exposure/load/' + id;

		return this.http.get<Exposure>(uri_);
	}
	
	//********************************************************************
	// gets all Exposure
	// returns the results untouched as JSON representation of an
	// Observable array of Exposure models
	// delegates via URI
	//********************************************************************
	getExposures() : Observable<Exposure[]> {
		const uri_ = this.apiUrl + '/Exposure/';

		return this
			.http.get<Exposure[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Claim on a Exposure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignClaim( exposureId, _claimId ): Observable<any> {

		// get the Exposure from storage
		this.loadHelper( exposureId );

	// get the Claim from storage
	var tmp 	= new ClaimService(this.http).getClaim(_claimId);

	// assign the Claim
	this.exposure.claim = tmp;

	// save the Exposure
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Claim on a Exposure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignClaim( exposureId ): Observable<any> {

		// get the Exposure from storage
		this.loadHelper( exposureId );

	// assign Claim to null
	this.exposure.claim = null;

	// save the Exposure
	return this.saveHelper();
}

		//********************************************************************
	// assigns a PolicyCoverage on a Exposure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPolicyCoverage( exposureId, _policyCoverageId ): Observable<any> {

		// get the Exposure from storage
		this.loadHelper( exposureId );

	// get the PolicyCoverage from storage
	var tmp 	= new PolicyCoverageService(this.http).getPolicyCoverage(_policyCoverageId);

	// assign the PolicyCoverage
	this.exposure.policyCoverage = tmp;

	// save the Exposure
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PolicyCoverage on a Exposure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPolicyCoverage( exposureId ): Observable<any> {

		// get the Exposure from storage
		this.loadHelper( exposureId );

	// assign PolicyCoverage to null
	this.exposure.policyCoverage = null;

	// save the Exposure
	return this.saveHelper();
}

		//********************************************************************
	// assigns a InsuredObject on a Exposure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInsuredObject( exposureId, _insuredObjectId ): Observable<any> {

		// get the Exposure from storage
		this.loadHelper( exposureId );

	// get the InsuredObject from storage
	var tmp 	= new InsuredObjectService(this.http).getInsuredObject(_insuredObjectId);

	// assign the InsuredObject
	this.exposure.insuredObject = tmp;

	// save the Exposure
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a InsuredObject on a Exposure
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInsuredObject( exposureId ): Observable<any> {

		// get the Exposure from storage
		this.loadHelper( exposureId );

	// assign InsuredObject to null
	this.exposure.insuredObject = null;

	// save the Exposure
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more reservesIds as a Reserves
	// to a Exposure
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addReserves( exposureId, reservesIds ): Observable<any> {

		// get the Exposure
		this.loadHelper( exposureId );

	// split on a comma with no spaces
	var idList = reservesIds.split(',')

	// iterate over array of reserves ids
	idList.forEach(function (id) {
		// read the ClaimReserve
		var claimReserve = new ClaimReserveService(this.http).getClaimReserve(id);
		// add the ClaimReserve if not already assigned
		if ( this.exposure.reserves.indexOf(claimReserve) == -1 )
		this.exposure.reserves.push(claimReserve);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more reservesIds as a Reserves
	// from a Exposure
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeReserves( exposureId, reservesIds ): Observable<any> {

		// get the Exposure
		this.loadHelper( exposureId );


	// split on a comma with no spaces
	var idList 					= reservesIds.split(',');
	var reserves 	= this.exposure.reserves;

	if ( reserves != null && reservesIds != null ) {

		// iterate over array of reserves ids
		reserves.forEach(function (obj) {
			if ( reservesIds.indexOf(obj._id) > -1 ) {
				// remove the ClaimReserve
				this.exposure.reserves.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more paymentsIds as a Payments
	// to a Exposure
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPayments( exposureId, paymentsIds ): Observable<any> {

		// get the Exposure
		this.loadHelper( exposureId );

	// split on a comma with no spaces
	var idList = paymentsIds.split(',')

	// iterate over array of payments ids
	idList.forEach(function (id) {
		// read the ClaimPayment
		var claimPayment = new ClaimPaymentService(this.http).getClaimPayment(id);
		// add the ClaimPayment if not already assigned
		if ( this.exposure.payments.indexOf(claimPayment) == -1 )
		this.exposure.payments.push(claimPayment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more paymentsIds as a Payments
	// from a Exposure
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePayments( exposureId, paymentsIds ): Observable<any> {

		// get the Exposure
		this.loadHelper( exposureId );


	// split on a comma with no spaces
	var idList 					= paymentsIds.split(',');
	var payments 	= this.exposure.payments;

	if ( payments != null && paymentsIds != null ) {

		// iterate over array of payments ids
		payments.forEach(function (obj) {
			if ( paymentsIds.indexOf(obj._id) > -1 ) {
				// remove the ClaimPayment
				this.exposure.payments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Exposure
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Exposure/update/' + this.exposure;

	return  this.http.post(uri_, this.exposure );
}

	//********************************************************************
	// loadHelper - internal helper to load a Exposure
	//********************************************************************	
	loadHelper( id ) {
		this.getExposure(id)
			.subscribe((res : Exposure) => {
				this.exposure = res;
			});
	}
}