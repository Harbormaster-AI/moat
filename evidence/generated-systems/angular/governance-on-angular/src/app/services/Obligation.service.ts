import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Obligation} from '../models/Obligation';
import {RegulationService} from '../services/Regulation.service';
import {ControlService} from '../services/Control.service';
import {PolicyService} from '../services/Policy.service';
import {ContractService} from '../services/Contract.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ObligationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	obligation : Obligation;

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
	// add a Obligation
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addObligation(referenceNumber, descriptionText, Regulation, Controls, Policies, Contracts, ObligationType, ReviewFrequency) : Observable<any> {
		const uri_ = this.apiUrl + '/Obligation/create';
		const obj = {
			      		referenceNumber: referenceNumber,
      		descriptionText: descriptionText,
      		Regulation: Regulation != null && Regulation.length > 0 ? Regulation : null,
      		Controls: Controls != null && Controls.length > 0 ? Controls : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
      		Contracts: Contracts != null && Contracts.length > 0 ? Contracts : null,
      		ObligationType: ObligationType,
			ReviewFrequency: ReviewFrequency
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Obligation
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateObligation(referenceNumber, descriptionText, Regulation, Controls, Policies, Contracts, ObligationType, ReviewFrequency, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Obligation/update/' + id;
		const obj = {
				      		referenceNumber: referenceNumber,
      		descriptionText: descriptionText,
      		Regulation: Regulation != null && Regulation.length > 0 ? Regulation : null,
      		Controls: Controls != null && Controls.length > 0 ? Controls : null,
      		Policies: Policies != null && Policies.length > 0 ? Policies : null,
      		Contracts: Contracts != null && Contracts.length > 0 ? Contracts : null,
      		ObligationType: ObligationType,
			ReviewFrequency: ReviewFrequency
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Obligation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteObligation(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Obligation/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Obligation
	// returns the results untouched as an Observable Obligation
	// Obligation model
	// delegates via URI
	//********************************************************************
	getObligation(id) : Observable<Obligation> {
		const uri_ = this.apiUrl + '/Obligation/load/' + id;

		return this.http.get<Obligation>(uri_);
	}
	
	//********************************************************************
	// gets all Obligation
	// returns the results untouched as JSON representation of an
	// Observable array of Obligation models
	// delegates via URI
	//********************************************************************
	getObligations() : Observable<Obligation[]> {
		const uri_ = this.apiUrl + '/Obligation/';

		return this
			.http.get<Obligation[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Regulation on a Obligation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRegulation( obligationId, _regulationId ): Observable<any> {

		// get the Obligation from storage
		this.loadHelper( obligationId );

	// get the Regulation from storage
	var tmp 	= new RegulationService(this.http).getRegulation(_regulationId);

	// assign the Regulation
	this.obligation.regulation = tmp;

	// save the Obligation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Regulation on a Obligation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRegulation( obligationId ): Observable<any> {

		// get the Obligation from storage
		this.loadHelper( obligationId );

	// assign Regulation to null
	this.obligation.regulation = null;

	// save the Obligation
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more controlsIds as a Controls
	// to a Obligation
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addControls( obligationId, controlsIds ): Observable<any> {

		// get the Obligation
		this.loadHelper( obligationId );

	// split on a comma with no spaces
	var idList = controlsIds.split(',')

	// iterate over array of controls ids
	idList.forEach(function (id) {
		// read the Control
		var control = new ControlService(this.http).getControl(id);
		// add the Control if not already assigned
		if ( this.obligation.controls.indexOf(control) == -1 )
		this.obligation.controls.push(control);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more controlsIds as a Controls
	// from a Obligation
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeControls( obligationId, controlsIds ): Observable<any> {

		// get the Obligation
		this.loadHelper( obligationId );


	// split on a comma with no spaces
	var idList 					= controlsIds.split(',');
	var controls 	= this.obligation.controls;

	if ( controls != null && controlsIds != null ) {

		// iterate over array of controls ids
		controls.forEach(function (obj) {
			if ( controlsIds.indexOf(obj._id) > -1 ) {
				// remove the Control
				this.obligation.controls.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more policiesIds as a Policies
	// to a Obligation
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPolicies( obligationId, policiesIds ): Observable<any> {

		// get the Obligation
		this.loadHelper( obligationId );

	// split on a comma with no spaces
	var idList = policiesIds.split(',')

	// iterate over array of policies ids
	idList.forEach(function (id) {
		// read the Policy
		var policy = new PolicyService(this.http).getPolicy(id);
		// add the Policy if not already assigned
		if ( this.obligation.policies.indexOf(policy) == -1 )
		this.obligation.policies.push(policy);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more policiesIds as a Policies
	// from a Obligation
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePolicies( obligationId, policiesIds ): Observable<any> {

		// get the Obligation
		this.loadHelper( obligationId );


	// split on a comma with no spaces
	var idList 					= policiesIds.split(',');
	var policies 	= this.obligation.policies;

	if ( policies != null && policiesIds != null ) {

		// iterate over array of policies ids
		policies.forEach(function (obj) {
			if ( policiesIds.indexOf(obj._id) > -1 ) {
				// remove the Policy
				this.obligation.policies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more contractsIds as a Contracts
	// to a Obligation
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addContracts( obligationId, contractsIds ): Observable<any> {

		// get the Obligation
		this.loadHelper( obligationId );

	// split on a comma with no spaces
	var idList = contractsIds.split(',')

	// iterate over array of contracts ids
	idList.forEach(function (id) {
		// read the Contract
		var contract = new ContractService(this.http).getContract(id);
		// add the Contract if not already assigned
		if ( this.obligation.contracts.indexOf(contract) == -1 )
		this.obligation.contracts.push(contract);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more contractsIds as a Contracts
	// from a Obligation
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeContracts( obligationId, contractsIds ): Observable<any> {

		// get the Obligation
		this.loadHelper( obligationId );


	// split on a comma with no spaces
	var idList 					= contractsIds.split(',');
	var contracts 	= this.obligation.contracts;

	if ( contracts != null && contractsIds != null ) {

		// iterate over array of contracts ids
		contracts.forEach(function (obj) {
			if ( contractsIds.indexOf(obj._id) > -1 ) {
				// remove the Contract
				this.obligation.contracts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Obligation
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Obligation/update/' + this.obligation;

	return  this.http.post(uri_, this.obligation );
}

	//********************************************************************
	// loadHelper - internal helper to load a Obligation
	//********************************************************************	
	loadHelper( id ) {
		this.getObligation(id)
			.subscribe((res : Obligation) => {
				this.obligation = res;
			});
	}
}