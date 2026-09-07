import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ThirdParty} from '../models/ThirdParty';
import {OrganizationService} from '../services/Organization.service';
import {DataProcessingActivityService} from '../services/DataProcessingActivity.service';
import {ThirdPartyAssessmentService} from '../services/ThirdPartyAssessment.service';
import {ContractService} from '../services/Contract.service';
import {ObligationService} from '../services/Obligation.service';
import {DataBreachService} from '../services/DataBreach.service';
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
	addThirdParty(name, country, contactEmail, Organization, ProcessingActivities, Assessments, Contracts, Obligations, DataBreaches, ThirdPartyType, Criticality) : Observable<any> {
		const uri_ = this.apiUrl + '/ThirdParty/create';
		const obj = {
			      		name: name,
      		country: country,
      		contactEmail: contactEmail,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		ProcessingActivities: ProcessingActivities != null && ProcessingActivities.length > 0 ? ProcessingActivities : null,
      		Assessments: Assessments != null && Assessments.length > 0 ? Assessments : null,
      		Contracts: Contracts != null && Contracts.length > 0 ? Contracts : null,
      		Obligations: Obligations != null && Obligations.length > 0 ? Obligations : null,
      		DataBreaches: DataBreaches != null && DataBreaches.length > 0 ? DataBreaches : null,
      		ThirdPartyType: ThirdPartyType,
			Criticality: Criticality
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ThirdParty
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateThirdParty(name, country, contactEmail, Organization, ProcessingActivities, Assessments, Contracts, Obligations, DataBreaches, ThirdPartyType, Criticality, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ThirdParty/update/' + id;
		const obj = {
				      		name: name,
      		country: country,
      		contactEmail: contactEmail,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		ProcessingActivities: ProcessingActivities != null && ProcessingActivities.length > 0 ? ProcessingActivities : null,
      		Assessments: Assessments != null && Assessments.length > 0 ? Assessments : null,
      		Contracts: Contracts != null && Contracts.length > 0 ? Contracts : null,
      		Obligations: Obligations != null && Obligations.length > 0 ? Obligations : null,
      		DataBreaches: DataBreaches != null && DataBreaches.length > 0 ? DataBreaches : null,
      		ThirdPartyType: ThirdPartyType,
			Criticality: Criticality
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
	// assigns a Organization on a ThirdParty
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( thirdPartyId, _organizationId ): Observable<any> {

		// get the ThirdParty from storage
		this.loadHelper( thirdPartyId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.thirdParty.organization = tmp;

	// save the ThirdParty
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a ThirdParty
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( thirdPartyId ): Observable<any> {

		// get the ThirdParty from storage
		this.loadHelper( thirdPartyId );

	// assign Organization to null
	this.thirdParty.organization = null;

	// save the ThirdParty
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more processingActivitiesIds as a ProcessingActivities
	// to a ThirdParty
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProcessingActivities( thirdPartyId, processingActivitiesIds ): Observable<any> {

		// get the ThirdParty
		this.loadHelper( thirdPartyId );

	// split on a comma with no spaces
	var idList = processingActivitiesIds.split(',')

	// iterate over array of processingActivities ids
	idList.forEach(function (id) {
		// read the DataProcessingActivity
		var dataProcessingActivity = new DataProcessingActivityService(this.http).getDataProcessingActivity(id);
		// add the DataProcessingActivity if not already assigned
		if ( this.thirdParty.processingActivities.indexOf(dataProcessingActivity) == -1 )
		this.thirdParty.processingActivities.push(dataProcessingActivity);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more processingActivitiesIds as a ProcessingActivities
	// from a ThirdParty
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProcessingActivities( thirdPartyId, processingActivitiesIds ): Observable<any> {

		// get the ThirdParty
		this.loadHelper( thirdPartyId );


	// split on a comma with no spaces
	var idList 					= processingActivitiesIds.split(',');
	var processingActivities 	= this.thirdParty.processingActivities;

	if ( processingActivities != null && processingActivitiesIds != null ) {

		// iterate over array of processingActivities ids
		processingActivities.forEach(function (obj) {
			if ( processingActivitiesIds.indexOf(obj._id) > -1 ) {
				// remove the DataProcessingActivity
				this.thirdParty.processingActivities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more assessmentsIds as a Assessments
	// to a ThirdParty
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAssessments( thirdPartyId, assessmentsIds ): Observable<any> {

		// get the ThirdParty
		this.loadHelper( thirdPartyId );

	// split on a comma with no spaces
	var idList = assessmentsIds.split(',')

	// iterate over array of assessments ids
	idList.forEach(function (id) {
		// read the ThirdPartyAssessment
		var thirdPartyAssessment = new ThirdPartyAssessmentService(this.http).getThirdPartyAssessment(id);
		// add the ThirdPartyAssessment if not already assigned
		if ( this.thirdParty.assessments.indexOf(thirdPartyAssessment) == -1 )
		this.thirdParty.assessments.push(thirdPartyAssessment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more assessmentsIds as a Assessments
	// from a ThirdParty
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAssessments( thirdPartyId, assessmentsIds ): Observable<any> {

		// get the ThirdParty
		this.loadHelper( thirdPartyId );


	// split on a comma with no spaces
	var idList 					= assessmentsIds.split(',');
	var assessments 	= this.thirdParty.assessments;

	if ( assessments != null && assessmentsIds != null ) {

		// iterate over array of assessments ids
		assessments.forEach(function (obj) {
			if ( assessmentsIds.indexOf(obj._id) > -1 ) {
				// remove the ThirdPartyAssessment
				this.thirdParty.assessments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more contractsIds as a Contracts
	// to a ThirdParty
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addContracts( thirdPartyId, contractsIds ): Observable<any> {

		// get the ThirdParty
		this.loadHelper( thirdPartyId );

	// split on a comma with no spaces
	var idList = contractsIds.split(',')

	// iterate over array of contracts ids
	idList.forEach(function (id) {
		// read the Contract
		var contract = new ContractService(this.http).getContract(id);
		// add the Contract if not already assigned
		if ( this.thirdParty.contracts.indexOf(contract) == -1 )
		this.thirdParty.contracts.push(contract);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more contractsIds as a Contracts
	// from a ThirdParty
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeContracts( thirdPartyId, contractsIds ): Observable<any> {

		// get the ThirdParty
		this.loadHelper( thirdPartyId );


	// split on a comma with no spaces
	var idList 					= contractsIds.split(',');
	var contracts 	= this.thirdParty.contracts;

	if ( contracts != null && contractsIds != null ) {

		// iterate over array of contracts ids
		contracts.forEach(function (obj) {
			if ( contractsIds.indexOf(obj._id) > -1 ) {
				// remove the Contract
				this.thirdParty.contracts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more obligationsIds as a Obligations
	// to a ThirdParty
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addObligations( thirdPartyId, obligationsIds ): Observable<any> {

		// get the ThirdParty
		this.loadHelper( thirdPartyId );

	// split on a comma with no spaces
	var idList = obligationsIds.split(',')

	// iterate over array of obligations ids
	idList.forEach(function (id) {
		// read the Obligation
		var obligation = new ObligationService(this.http).getObligation(id);
		// add the Obligation if not already assigned
		if ( this.thirdParty.obligations.indexOf(obligation) == -1 )
		this.thirdParty.obligations.push(obligation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more obligationsIds as a Obligations
	// from a ThirdParty
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeObligations( thirdPartyId, obligationsIds ): Observable<any> {

		// get the ThirdParty
		this.loadHelper( thirdPartyId );


	// split on a comma with no spaces
	var idList 					= obligationsIds.split(',');
	var obligations 	= this.thirdParty.obligations;

	if ( obligations != null && obligationsIds != null ) {

		// iterate over array of obligations ids
		obligations.forEach(function (obj) {
			if ( obligationsIds.indexOf(obj._id) > -1 ) {
				// remove the Obligation
				this.thirdParty.obligations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dataBreachesIds as a DataBreaches
	// to a ThirdParty
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDataBreaches( thirdPartyId, dataBreachesIds ): Observable<any> {

		// get the ThirdParty
		this.loadHelper( thirdPartyId );

	// split on a comma with no spaces
	var idList = dataBreachesIds.split(',')

	// iterate over array of dataBreaches ids
	idList.forEach(function (id) {
		// read the DataBreach
		var dataBreach = new DataBreachService(this.http).getDataBreach(id);
		// add the DataBreach if not already assigned
		if ( this.thirdParty.dataBreaches.indexOf(dataBreach) == -1 )
		this.thirdParty.dataBreaches.push(dataBreach);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dataBreachesIds as a DataBreaches
	// from a ThirdParty
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDataBreaches( thirdPartyId, dataBreachesIds ): Observable<any> {

		// get the ThirdParty
		this.loadHelper( thirdPartyId );


	// split on a comma with no spaces
	var idList 					= dataBreachesIds.split(',');
	var dataBreaches 	= this.thirdParty.dataBreaches;

	if ( dataBreaches != null && dataBreachesIds != null ) {

		// iterate over array of dataBreaches ids
		dataBreaches.forEach(function (obj) {
			if ( dataBreachesIds.indexOf(obj._id) > -1 ) {
				// remove the DataBreach
				this.thirdParty.dataBreaches.pop(obj);
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