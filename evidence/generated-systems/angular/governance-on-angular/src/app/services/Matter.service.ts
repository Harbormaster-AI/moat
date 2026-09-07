import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Matter} from '../models/Matter';
import {LegalHoldService} from '../services/LegalHold.service';
import {OrganizationService} from '../services/Organization.service';
import {DataBreachService} from '../services/DataBreach.service';
import {ContractService} from '../services/Contract.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class MatterService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	matter : Matter;

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
	// add a Matter
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addMatter(matterName, leadCounsel, LegalHolds, Organization, DataBreaches, Contracts, MatterType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Matter/create';
		const obj = {
			      		matterName: matterName,
      		leadCounsel: leadCounsel,
      		LegalHolds: LegalHolds != null && LegalHolds.length > 0 ? LegalHolds : null,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		DataBreaches: DataBreaches != null && DataBreaches.length > 0 ? DataBreaches : null,
      		Contracts: Contracts != null && Contracts.length > 0 ? Contracts : null,
      		MatterType: MatterType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Matter
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateMatter(matterName, leadCounsel, LegalHolds, Organization, DataBreaches, Contracts, MatterType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Matter/update/' + id;
		const obj = {
				      		matterName: matterName,
      		leadCounsel: leadCounsel,
      		LegalHolds: LegalHolds != null && LegalHolds.length > 0 ? LegalHolds : null,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		DataBreaches: DataBreaches != null && DataBreaches.length > 0 ? DataBreaches : null,
      		Contracts: Contracts != null && Contracts.length > 0 ? Contracts : null,
      		MatterType: MatterType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Matter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteMatter(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Matter/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Matter
	// returns the results untouched as an Observable Matter
	// Matter model
	// delegates via URI
	//********************************************************************
	getMatter(id) : Observable<Matter> {
		const uri_ = this.apiUrl + '/Matter/load/' + id;

		return this.http.get<Matter>(uri_);
	}
	
	//********************************************************************
	// gets all Matter
	// returns the results untouched as JSON representation of an
	// Observable array of Matter models
	// delegates via URI
	//********************************************************************
	getMatters() : Observable<Matter[]> {
		const uri_ = this.apiUrl + '/Matter/';

		return this
			.http.get<Matter[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Matter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( matterId, _organizationId ): Observable<any> {

		// get the Matter from storage
		this.loadHelper( matterId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.matter.organization = tmp;

	// save the Matter
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Matter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( matterId ): Observable<any> {

		// get the Matter from storage
		this.loadHelper( matterId );

	// assign Organization to null
	this.matter.organization = null;

	// save the Matter
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more legalHoldsIds as a LegalHolds
	// to a Matter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLegalHolds( matterId, legalHoldsIds ): Observable<any> {

		// get the Matter
		this.loadHelper( matterId );

	// split on a comma with no spaces
	var idList = legalHoldsIds.split(',')

	// iterate over array of legalHolds ids
	idList.forEach(function (id) {
		// read the LegalHold
		var legalHold = new LegalHoldService(this.http).getLegalHold(id);
		// add the LegalHold if not already assigned
		if ( this.matter.legalHolds.indexOf(legalHold) == -1 )
		this.matter.legalHolds.push(legalHold);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more legalHoldsIds as a LegalHolds
	// from a Matter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLegalHolds( matterId, legalHoldsIds ): Observable<any> {

		// get the Matter
		this.loadHelper( matterId );


	// split on a comma with no spaces
	var idList 					= legalHoldsIds.split(',');
	var legalHolds 	= this.matter.legalHolds;

	if ( legalHolds != null && legalHoldsIds != null ) {

		// iterate over array of legalHolds ids
		legalHolds.forEach(function (obj) {
			if ( legalHoldsIds.indexOf(obj._id) > -1 ) {
				// remove the LegalHold
				this.matter.legalHolds.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dataBreachesIds as a DataBreaches
	// to a Matter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDataBreaches( matterId, dataBreachesIds ): Observable<any> {

		// get the Matter
		this.loadHelper( matterId );

	// split on a comma with no spaces
	var idList = dataBreachesIds.split(',')

	// iterate over array of dataBreaches ids
	idList.forEach(function (id) {
		// read the DataBreach
		var dataBreach = new DataBreachService(this.http).getDataBreach(id);
		// add the DataBreach if not already assigned
		if ( this.matter.dataBreaches.indexOf(dataBreach) == -1 )
		this.matter.dataBreaches.push(dataBreach);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dataBreachesIds as a DataBreaches
	// from a Matter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDataBreaches( matterId, dataBreachesIds ): Observable<any> {

		// get the Matter
		this.loadHelper( matterId );


	// split on a comma with no spaces
	var idList 					= dataBreachesIds.split(',');
	var dataBreaches 	= this.matter.dataBreaches;

	if ( dataBreaches != null && dataBreachesIds != null ) {

		// iterate over array of dataBreaches ids
		dataBreaches.forEach(function (obj) {
			if ( dataBreachesIds.indexOf(obj._id) > -1 ) {
				// remove the DataBreach
				this.matter.dataBreaches.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more contractsIds as a Contracts
	// to a Matter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addContracts( matterId, contractsIds ): Observable<any> {

		// get the Matter
		this.loadHelper( matterId );

	// split on a comma with no spaces
	var idList = contractsIds.split(',')

	// iterate over array of contracts ids
	idList.forEach(function (id) {
		// read the Contract
		var contract = new ContractService(this.http).getContract(id);
		// add the Contract if not already assigned
		if ( this.matter.contracts.indexOf(contract) == -1 )
		this.matter.contracts.push(contract);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more contractsIds as a Contracts
	// from a Matter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeContracts( matterId, contractsIds ): Observable<any> {

		// get the Matter
		this.loadHelper( matterId );


	// split on a comma with no spaces
	var idList 					= contractsIds.split(',');
	var contracts 	= this.matter.contracts;

	if ( contracts != null && contractsIds != null ) {

		// iterate over array of contracts ids
		contracts.forEach(function (obj) {
			if ( contractsIds.indexOf(obj._id) > -1 ) {
				// remove the Contract
				this.matter.contracts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Matter
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Matter/update/' + this.matter;

	return  this.http.post(uri_, this.matter );
}

	//********************************************************************
	// loadHelper - internal helper to load a Matter
	//********************************************************************	
	loadHelper( id ) {
		this.getMatter(id)
			.subscribe((res : Matter) => {
				this.matter = res;
			});
	}
}