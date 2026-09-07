import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Regulation} from '../models/Regulation';
import {ObligationService} from '../services/Obligation.service';
import {ComplianceProgramService} from '../services/ComplianceProgram.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class RegulationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	regulation : Regulation;

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
	// add a Regulation
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addRegulation(name, citation, jurisdiction, publicationUrl, Obligations, CompliancePrograms) : Observable<any> {
		const uri_ = this.apiUrl + '/Regulation/create';
		const obj = {
			      		name: name,
      		citation: citation,
      		jurisdiction: jurisdiction,
      		publicationUrl: publicationUrl,
      		Obligations: Obligations != null && Obligations.length > 0 ? Obligations : null,
			CompliancePrograms: CompliancePrograms != null && CompliancePrograms.length > 0 ? CompliancePrograms : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Regulation
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateRegulation(name, citation, jurisdiction, publicationUrl, Obligations, CompliancePrograms, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Regulation/update/' + id;
		const obj = {
				      		name: name,
      		citation: citation,
      		jurisdiction: jurisdiction,
      		publicationUrl: publicationUrl,
      		Obligations: Obligations != null && Obligations.length > 0 ? Obligations : null,
			CompliancePrograms: CompliancePrograms != null && CompliancePrograms.length > 0 ? CompliancePrograms : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Regulation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteRegulation(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Regulation/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Regulation
	// returns the results untouched as an Observable Regulation
	// Regulation model
	// delegates via URI
	//********************************************************************
	getRegulation(id) : Observable<Regulation> {
		const uri_ = this.apiUrl + '/Regulation/load/' + id;

		return this.http.get<Regulation>(uri_);
	}
	
	//********************************************************************
	// gets all Regulation
	// returns the results untouched as JSON representation of an
	// Observable array of Regulation models
	// delegates via URI
	//********************************************************************
	getRegulations() : Observable<Regulation[]> {
		const uri_ = this.apiUrl + '/Regulation/';

		return this
			.http.get<Regulation[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more obligationsIds as a Obligations
	// to a Regulation
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addObligations( regulationId, obligationsIds ): Observable<any> {

		// get the Regulation
		this.loadHelper( regulationId );

	// split on a comma with no spaces
	var idList = obligationsIds.split(',')

	// iterate over array of obligations ids
	idList.forEach(function (id) {
		// read the Obligation
		var obligation = new ObligationService(this.http).getObligation(id);
		// add the Obligation if not already assigned
		if ( this.regulation.obligations.indexOf(obligation) == -1 )
		this.regulation.obligations.push(obligation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more obligationsIds as a Obligations
	// from a Regulation
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeObligations( regulationId, obligationsIds ): Observable<any> {

		// get the Regulation
		this.loadHelper( regulationId );


	// split on a comma with no spaces
	var idList 					= obligationsIds.split(',');
	var obligations 	= this.regulation.obligations;

	if ( obligations != null && obligationsIds != null ) {

		// iterate over array of obligations ids
		obligations.forEach(function (obj) {
			if ( obligationsIds.indexOf(obj._id) > -1 ) {
				// remove the Obligation
				this.regulation.obligations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more complianceProgramsIds as a CompliancePrograms
	// to a Regulation
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCompliancePrograms( regulationId, complianceProgramsIds ): Observable<any> {

		// get the Regulation
		this.loadHelper( regulationId );

	// split on a comma with no spaces
	var idList = complianceProgramsIds.split(',')

	// iterate over array of compliancePrograms ids
	idList.forEach(function (id) {
		// read the ComplianceProgram
		var complianceProgram = new ComplianceProgramService(this.http).getComplianceProgram(id);
		// add the ComplianceProgram if not already assigned
		if ( this.regulation.compliancePrograms.indexOf(complianceProgram) == -1 )
		this.regulation.compliancePrograms.push(complianceProgram);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more complianceProgramsIds as a CompliancePrograms
	// from a Regulation
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCompliancePrograms( regulationId, complianceProgramsIds ): Observable<any> {

		// get the Regulation
		this.loadHelper( regulationId );


	// split on a comma with no spaces
	var idList 					= complianceProgramsIds.split(',');
	var compliancePrograms 	= this.regulation.compliancePrograms;

	if ( compliancePrograms != null && complianceProgramsIds != null ) {

		// iterate over array of compliancePrograms ids
		compliancePrograms.forEach(function (obj) {
			if ( complianceProgramsIds.indexOf(obj._id) > -1 ) {
				// remove the ComplianceProgram
				this.regulation.compliancePrograms.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Regulation
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Regulation/update/' + this.regulation;

	return  this.http.post(uri_, this.regulation );
}

	//********************************************************************
	// loadHelper - internal helper to load a Regulation
	//********************************************************************	
	loadHelper( id ) {
		this.getRegulation(id)
			.subscribe((res : Regulation) => {
				this.regulation = res;
			});
	}
}