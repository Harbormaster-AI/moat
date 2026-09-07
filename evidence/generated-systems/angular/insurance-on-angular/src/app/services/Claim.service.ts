import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Claim} from '../models/Claim';
import {PolicyService} from '../services/Policy.service';
import {CustomerService} from '../services/Customer.service';
import {AdjusterService} from '../services/Adjuster.service';
import {IncidentService} from '../services/Incident.service';
import {ExposureService} from '../services/Exposure.service';
import {ClaimReserveService} from '../services/ClaimReserve.service';
import {ClaimPaymentService} from '../services/ClaimPayment.service';
import {ServiceProviderService} from '../services/ServiceProvider.service';
import {SubrogationRecoveryService} from '../services/SubrogationRecovery.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ClaimService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	claim : Claim;

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
	// add a Claim
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addClaim(claimNumber, noticeDate, lossDate, reportedBy, Policy, Customer, Adjuster, Incident, Exposures, Reserves, ClaimPayments, ServiceProviders, Subrogations, Status, LossCause) : Observable<any> {
		const uri_ = this.apiUrl + '/Claim/create';
		const obj = {
			      		claimNumber: claimNumber,
      		noticeDate: noticeDate,
      		lossDate: lossDate,
      		reportedBy: reportedBy,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Adjuster: Adjuster != null && Adjuster.length > 0 ? Adjuster : null,
      		Incident: Incident != null && Incident.length > 0 ? Incident : null,
      		Exposures: Exposures != null && Exposures.length > 0 ? Exposures : null,
      		Reserves: Reserves != null && Reserves.length > 0 ? Reserves : null,
      		ClaimPayments: ClaimPayments != null && ClaimPayments.length > 0 ? ClaimPayments : null,
      		ServiceProviders: ServiceProviders != null && ServiceProviders.length > 0 ? ServiceProviders : null,
      		Subrogations: Subrogations != null && Subrogations.length > 0 ? Subrogations : null,
      		Status: Status,
			LossCause: LossCause
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateClaim(claimNumber, noticeDate, lossDate, reportedBy, Policy, Customer, Adjuster, Incident, Exposures, Reserves, ClaimPayments, ServiceProviders, Subrogations, Status, LossCause, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Claim/update/' + id;
		const obj = {
				      		claimNumber: claimNumber,
      		noticeDate: noticeDate,
      		lossDate: lossDate,
      		reportedBy: reportedBy,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Adjuster: Adjuster != null && Adjuster.length > 0 ? Adjuster : null,
      		Incident: Incident != null && Incident.length > 0 ? Incident : null,
      		Exposures: Exposures != null && Exposures.length > 0 ? Exposures : null,
      		Reserves: Reserves != null && Reserves.length > 0 ? Reserves : null,
      		ClaimPayments: ClaimPayments != null && ClaimPayments.length > 0 ? ClaimPayments : null,
      		ServiceProviders: ServiceProviders != null && ServiceProviders.length > 0 ? ServiceProviders : null,
      		Subrogations: Subrogations != null && Subrogations.length > 0 ? Subrogations : null,
      		Status: Status,
			LossCause: LossCause
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteClaim(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Claim/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Claim
	// returns the results untouched as an Observable Claim
	// Claim model
	// delegates via URI
	//********************************************************************
	getClaim(id) : Observable<Claim> {
		const uri_ = this.apiUrl + '/Claim/load/' + id;

		return this.http.get<Claim>(uri_);
	}
	
	//********************************************************************
	// gets all Claim
	// returns the results untouched as JSON representation of an
	// Observable array of Claim models
	// delegates via URI
	//********************************************************************
	getClaims() : Observable<Claim[]> {
		const uri_ = this.apiUrl + '/Claim/';

		return this
			.http.get<Claim[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Policy on a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPolicy( claimId, _policyId ): Observable<any> {

		// get the Claim from storage
		this.loadHelper( claimId );

	// get the Policy from storage
	var tmp 	= new PolicyService(this.http).getPolicy(_policyId);

	// assign the Policy
	this.claim.policy = tmp;

	// save the Claim
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Policy on a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPolicy( claimId ): Observable<any> {

		// get the Claim from storage
		this.loadHelper( claimId );

	// assign Policy to null
	this.claim.policy = null;

	// save the Claim
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Customer on a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCustomer( claimId, _customerId ): Observable<any> {

		// get the Claim from storage
		this.loadHelper( claimId );

	// get the Customer from storage
	var tmp 	= new CustomerService(this.http).getCustomer(_customerId);

	// assign the Customer
	this.claim.customer = tmp;

	// save the Claim
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Customer on a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCustomer( claimId ): Observable<any> {

		// get the Claim from storage
		this.loadHelper( claimId );

	// assign Customer to null
	this.claim.customer = null;

	// save the Claim
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Adjuster on a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAdjuster( claimId, _adjusterId ): Observable<any> {

		// get the Claim from storage
		this.loadHelper( claimId );

	// get the Adjuster from storage
	var tmp 	= new AdjusterService(this.http).getAdjuster(_adjusterId);

	// assign the Adjuster
	this.claim.adjuster = tmp;

	// save the Claim
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Adjuster on a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAdjuster( claimId ): Observable<any> {

		// get the Claim from storage
		this.loadHelper( claimId );

	// assign Adjuster to null
	this.claim.adjuster = null;

	// save the Claim
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Incident on a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignIncident( claimId, _incidentId ): Observable<any> {

		// get the Claim from storage
		this.loadHelper( claimId );

	// get the Incident from storage
	var tmp 	= new IncidentService(this.http).getIncident(_incidentId);

	// assign the Incident
	this.claim.incident = tmp;

	// save the Claim
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Incident on a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignIncident( claimId ): Observable<any> {

		// get the Claim from storage
		this.loadHelper( claimId );

	// assign Incident to null
	this.claim.incident = null;

	// save the Claim
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more exposuresIds as a Exposures
	// to a Claim
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addExposures( claimId, exposuresIds ): Observable<any> {

		// get the Claim
		this.loadHelper( claimId );

	// split on a comma with no spaces
	var idList = exposuresIds.split(',')

	// iterate over array of exposures ids
	idList.forEach(function (id) {
		// read the Exposure
		var exposure = new ExposureService(this.http).getExposure(id);
		// add the Exposure if not already assigned
		if ( this.claim.exposures.indexOf(exposure) == -1 )
		this.claim.exposures.push(exposure);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more exposuresIds as a Exposures
	// from a Claim
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeExposures( claimId, exposuresIds ): Observable<any> {

		// get the Claim
		this.loadHelper( claimId );


	// split on a comma with no spaces
	var idList 					= exposuresIds.split(',');
	var exposures 	= this.claim.exposures;

	if ( exposures != null && exposuresIds != null ) {

		// iterate over array of exposures ids
		exposures.forEach(function (obj) {
			if ( exposuresIds.indexOf(obj._id) > -1 ) {
				// remove the Exposure
				this.claim.exposures.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more reservesIds as a Reserves
	// to a Claim
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addReserves( claimId, reservesIds ): Observable<any> {

		// get the Claim
		this.loadHelper( claimId );

	// split on a comma with no spaces
	var idList = reservesIds.split(',')

	// iterate over array of reserves ids
	idList.forEach(function (id) {
		// read the ClaimReserve
		var claimReserve = new ClaimReserveService(this.http).getClaimReserve(id);
		// add the ClaimReserve if not already assigned
		if ( this.claim.reserves.indexOf(claimReserve) == -1 )
		this.claim.reserves.push(claimReserve);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more reservesIds as a Reserves
	// from a Claim
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeReserves( claimId, reservesIds ): Observable<any> {

		// get the Claim
		this.loadHelper( claimId );


	// split on a comma with no spaces
	var idList 					= reservesIds.split(',');
	var reserves 	= this.claim.reserves;

	if ( reserves != null && reservesIds != null ) {

		// iterate over array of reserves ids
		reserves.forEach(function (obj) {
			if ( reservesIds.indexOf(obj._id) > -1 ) {
				// remove the ClaimReserve
				this.claim.reserves.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more claimPaymentsIds as a ClaimPayments
	// to a Claim
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addClaimPayments( claimId, claimPaymentsIds ): Observable<any> {

		// get the Claim
		this.loadHelper( claimId );

	// split on a comma with no spaces
	var idList = claimPaymentsIds.split(',')

	// iterate over array of claimPayments ids
	idList.forEach(function (id) {
		// read the ClaimPayment
		var claimPayment = new ClaimPaymentService(this.http).getClaimPayment(id);
		// add the ClaimPayment if not already assigned
		if ( this.claim.claimPayments.indexOf(claimPayment) == -1 )
		this.claim.claimPayments.push(claimPayment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more claimPaymentsIds as a ClaimPayments
	// from a Claim
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeClaimPayments( claimId, claimPaymentsIds ): Observable<any> {

		// get the Claim
		this.loadHelper( claimId );


	// split on a comma with no spaces
	var idList 					= claimPaymentsIds.split(',');
	var claimPayments 	= this.claim.claimPayments;

	if ( claimPayments != null && claimPaymentsIds != null ) {

		// iterate over array of claimPayments ids
		claimPayments.forEach(function (obj) {
			if ( claimPaymentsIds.indexOf(obj._id) > -1 ) {
				// remove the ClaimPayment
				this.claim.claimPayments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more serviceProvidersIds as a ServiceProviders
	// to a Claim
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addServiceProviders( claimId, serviceProvidersIds ): Observable<any> {

		// get the Claim
		this.loadHelper( claimId );

	// split on a comma with no spaces
	var idList = serviceProvidersIds.split(',')

	// iterate over array of serviceProviders ids
	idList.forEach(function (id) {
		// read the ServiceProvider
		var serviceProvider = new ServiceProviderService(this.http).getServiceProvider(id);
		// add the ServiceProvider if not already assigned
		if ( this.claim.serviceProviders.indexOf(serviceProvider) == -1 )
		this.claim.serviceProviders.push(serviceProvider);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more serviceProvidersIds as a ServiceProviders
	// from a Claim
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeServiceProviders( claimId, serviceProvidersIds ): Observable<any> {

		// get the Claim
		this.loadHelper( claimId );


	// split on a comma with no spaces
	var idList 					= serviceProvidersIds.split(',');
	var serviceProviders 	= this.claim.serviceProviders;

	if ( serviceProviders != null && serviceProvidersIds != null ) {

		// iterate over array of serviceProviders ids
		serviceProviders.forEach(function (obj) {
			if ( serviceProvidersIds.indexOf(obj._id) > -1 ) {
				// remove the ServiceProvider
				this.claim.serviceProviders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more subrogationsIds as a Subrogations
	// to a Claim
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSubrogations( claimId, subrogationsIds ): Observable<any> {

		// get the Claim
		this.loadHelper( claimId );

	// split on a comma with no spaces
	var idList = subrogationsIds.split(',')

	// iterate over array of subrogations ids
	idList.forEach(function (id) {
		// read the SubrogationRecovery
		var subrogationRecovery = new SubrogationRecoveryService(this.http).getSubrogationRecovery(id);
		// add the SubrogationRecovery if not already assigned
		if ( this.claim.subrogations.indexOf(subrogationRecovery) == -1 )
		this.claim.subrogations.push(subrogationRecovery);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more subrogationsIds as a Subrogations
	// from a Claim
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSubrogations( claimId, subrogationsIds ): Observable<any> {

		// get the Claim
		this.loadHelper( claimId );


	// split on a comma with no spaces
	var idList 					= subrogationsIds.split(',');
	var subrogations 	= this.claim.subrogations;

	if ( subrogations != null && subrogationsIds != null ) {

		// iterate over array of subrogations ids
		subrogations.forEach(function (obj) {
			if ( subrogationsIds.indexOf(obj._id) > -1 ) {
				// remove the SubrogationRecovery
				this.claim.subrogations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Claim
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Claim/update/' + this.claim;

	return  this.http.post(uri_, this.claim );
}

	//********************************************************************
	// loadHelper - internal helper to load a Claim
	//********************************************************************	
	loadHelper( id ) {
		this.getClaim(id)
			.subscribe((res : Claim) => {
				this.claim = res;
			});
	}
}