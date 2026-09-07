import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Contract} from '../models/Contract';
import {ThirdPartyService} from '../services/ThirdParty.service';
import {ObligationService} from '../services/Obligation.service';
import {DataProcessingActivityService} from '../services/DataProcessingActivity.service';
import {MatterService} from '../services/Matter.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ContractService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	contract : Contract;

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
	// add a Contract
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addContract(title, effectiveDate, expiryDate, repositoryUrl, ThirdParty, Obligations, DataProcessingActivities, Matter, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Contract/create';
		const obj = {
			      		title: title,
      		effectiveDate: effectiveDate,
      		expiryDate: expiryDate,
      		repositoryUrl: repositoryUrl,
      		ThirdParty: ThirdParty != null && ThirdParty.length > 0 ? ThirdParty : null,
      		Obligations: Obligations != null && Obligations.length > 0 ? Obligations : null,
      		DataProcessingActivities: DataProcessingActivities != null && DataProcessingActivities.length > 0 ? DataProcessingActivities : null,
      		Matter: Matter != null && Matter.length > 0 ? Matter : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Contract
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateContract(title, effectiveDate, expiryDate, repositoryUrl, ThirdParty, Obligations, DataProcessingActivities, Matter, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Contract/update/' + id;
		const obj = {
				      		title: title,
      		effectiveDate: effectiveDate,
      		expiryDate: expiryDate,
      		repositoryUrl: repositoryUrl,
      		ThirdParty: ThirdParty != null && ThirdParty.length > 0 ? ThirdParty : null,
      		Obligations: Obligations != null && Obligations.length > 0 ? Obligations : null,
      		DataProcessingActivities: DataProcessingActivities != null && DataProcessingActivities.length > 0 ? DataProcessingActivities : null,
      		Matter: Matter != null && Matter.length > 0 ? Matter : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Contract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteContract(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Contract/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Contract
	// returns the results untouched as an Observable Contract
	// Contract model
	// delegates via URI
	//********************************************************************
	getContract(id) : Observable<Contract> {
		const uri_ = this.apiUrl + '/Contract/load/' + id;

		return this.http.get<Contract>(uri_);
	}
	
	//********************************************************************
	// gets all Contract
	// returns the results untouched as JSON representation of an
	// Observable array of Contract models
	// delegates via URI
	//********************************************************************
	getContracts() : Observable<Contract[]> {
		const uri_ = this.apiUrl + '/Contract/';

		return this
			.http.get<Contract[]>(uri_);
	}
	
			//********************************************************************
	// assigns a ThirdParty on a Contract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignThirdParty( contractId, _thirdPartyId ): Observable<any> {

		// get the Contract from storage
		this.loadHelper( contractId );

	// get the ThirdParty from storage
	var tmp 	= new ThirdPartyService(this.http).getThirdParty(_thirdPartyId);

	// assign the ThirdParty
	this.contract.thirdParty = tmp;

	// save the Contract
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ThirdParty on a Contract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignThirdParty( contractId ): Observable<any> {

		// get the Contract from storage
		this.loadHelper( contractId );

	// assign ThirdParty to null
	this.contract.thirdParty = null;

	// save the Contract
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Matter on a Contract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignMatter( contractId, _matterId ): Observable<any> {

		// get the Contract from storage
		this.loadHelper( contractId );

	// get the Matter from storage
	var tmp 	= new MatterService(this.http).getMatter(_matterId);

	// assign the Matter
	this.contract.matter = tmp;

	// save the Contract
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Matter on a Contract
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignMatter( contractId ): Observable<any> {

		// get the Contract from storage
		this.loadHelper( contractId );

	// assign Matter to null
	this.contract.matter = null;

	// save the Contract
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more obligationsIds as a Obligations
	// to a Contract
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addObligations( contractId, obligationsIds ): Observable<any> {

		// get the Contract
		this.loadHelper( contractId );

	// split on a comma with no spaces
	var idList = obligationsIds.split(',')

	// iterate over array of obligations ids
	idList.forEach(function (id) {
		// read the Obligation
		var obligation = new ObligationService(this.http).getObligation(id);
		// add the Obligation if not already assigned
		if ( this.contract.obligations.indexOf(obligation) == -1 )
		this.contract.obligations.push(obligation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more obligationsIds as a Obligations
	// from a Contract
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeObligations( contractId, obligationsIds ): Observable<any> {

		// get the Contract
		this.loadHelper( contractId );


	// split on a comma with no spaces
	var idList 					= obligationsIds.split(',');
	var obligations 	= this.contract.obligations;

	if ( obligations != null && obligationsIds != null ) {

		// iterate over array of obligations ids
		obligations.forEach(function (obj) {
			if ( obligationsIds.indexOf(obj._id) > -1 ) {
				// remove the Obligation
				this.contract.obligations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dataProcessingActivitiesIds as a DataProcessingActivities
	// to a Contract
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDataProcessingActivities( contractId, dataProcessingActivitiesIds ): Observable<any> {

		// get the Contract
		this.loadHelper( contractId );

	// split on a comma with no spaces
	var idList = dataProcessingActivitiesIds.split(',')

	// iterate over array of dataProcessingActivities ids
	idList.forEach(function (id) {
		// read the DataProcessingActivity
		var dataProcessingActivity = new DataProcessingActivityService(this.http).getDataProcessingActivity(id);
		// add the DataProcessingActivity if not already assigned
		if ( this.contract.dataProcessingActivities.indexOf(dataProcessingActivity) == -1 )
		this.contract.dataProcessingActivities.push(dataProcessingActivity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dataProcessingActivitiesIds as a DataProcessingActivities
	// from a Contract
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDataProcessingActivities( contractId, dataProcessingActivitiesIds ): Observable<any> {

		// get the Contract
		this.loadHelper( contractId );


	// split on a comma with no spaces
	var idList 					= dataProcessingActivitiesIds.split(',');
	var dataProcessingActivities 	= this.contract.dataProcessingActivities;

	if ( dataProcessingActivities != null && dataProcessingActivitiesIds != null ) {

		// iterate over array of dataProcessingActivities ids
		dataProcessingActivities.forEach(function (obj) {
			if ( dataProcessingActivitiesIds.indexOf(obj._id) > -1 ) {
				// remove the DataProcessingActivity
				this.contract.dataProcessingActivities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Contract
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Contract/update/' + this.contract;

	return  this.http.post(uri_, this.contract );
}

	//********************************************************************
	// loadHelper - internal helper to load a Contract
	//********************************************************************	
	loadHelper( id ) {
		this.getContract(id)
			.subscribe((res : Contract) => {
				this.contract = res;
			});
	}
}