import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Underwriter} from '../models/Underwriter';
import {UnderwritingDecisionService} from '../services/UnderwritingDecision.service';
import {InsurerService} from '../services/Insurer.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class UnderwriterService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	underwriter : Underwriter;

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
	// add a Underwriter
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addUnderwriter(firstName, lastName, employeeId, authorityLimit, Decisions, Insurer) : Observable<any> {
		const uri_ = this.apiUrl + '/Underwriter/create';
		const obj = {
			      		firstName: firstName,
      		lastName: lastName,
      		employeeId: employeeId,
      		authorityLimit: authorityLimit,
      		Decisions: Decisions != null && Decisions.length > 0 ? Decisions : null,
			Insurer: Insurer != null && Insurer.length > 0 ? Insurer : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Underwriter
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateUnderwriter(firstName, lastName, employeeId, authorityLimit, Decisions, Insurer, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Underwriter/update/' + id;
		const obj = {
				      		firstName: firstName,
      		lastName: lastName,
      		employeeId: employeeId,
      		authorityLimit: authorityLimit,
      		Decisions: Decisions != null && Decisions.length > 0 ? Decisions : null,
			Insurer: Insurer != null && Insurer.length > 0 ? Insurer : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Underwriter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteUnderwriter(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Underwriter/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Underwriter
	// returns the results untouched as an Observable Underwriter
	// Underwriter model
	// delegates via URI
	//********************************************************************
	getUnderwriter(id) : Observable<Underwriter> {
		const uri_ = this.apiUrl + '/Underwriter/load/' + id;

		return this.http.get<Underwriter>(uri_);
	}
	
	//********************************************************************
	// gets all Underwriter
	// returns the results untouched as JSON representation of an
	// Observable array of Underwriter models
	// delegates via URI
	//********************************************************************
	getUnderwriters() : Observable<Underwriter[]> {
		const uri_ = this.apiUrl + '/Underwriter/';

		return this
			.http.get<Underwriter[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Insurer on a Underwriter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInsurer( underwriterId, _insurerId ): Observable<any> {

		// get the Underwriter from storage
		this.loadHelper( underwriterId );

	// get the Insurer from storage
	var tmp 	= new InsurerService(this.http).getInsurer(_insurerId);

	// assign the Insurer
	this.underwriter.insurer = tmp;

	// save the Underwriter
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Insurer on a Underwriter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInsurer( underwriterId ): Observable<any> {

		// get the Underwriter from storage
		this.loadHelper( underwriterId );

	// assign Insurer to null
	this.underwriter.insurer = null;

	// save the Underwriter
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more decisionsIds as a Decisions
	// to a Underwriter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDecisions( underwriterId, decisionsIds ): Observable<any> {

		// get the Underwriter
		this.loadHelper( underwriterId );

	// split on a comma with no spaces
	var idList = decisionsIds.split(',')

	// iterate over array of decisions ids
	idList.forEach(function (id) {
		// read the UnderwritingDecision
		var underwritingDecision = new UnderwritingDecisionService(this.http).getUnderwritingDecision(id);
		// add the UnderwritingDecision if not already assigned
		if ( this.underwriter.decisions.indexOf(underwritingDecision) == -1 )
		this.underwriter.decisions.push(underwritingDecision);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more decisionsIds as a Decisions
	// from a Underwriter
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDecisions( underwriterId, decisionsIds ): Observable<any> {

		// get the Underwriter
		this.loadHelper( underwriterId );


	// split on a comma with no spaces
	var idList 					= decisionsIds.split(',');
	var decisions 	= this.underwriter.decisions;

	if ( decisions != null && decisionsIds != null ) {

		// iterate over array of decisions ids
		decisions.forEach(function (obj) {
			if ( decisionsIds.indexOf(obj._id) > -1 ) {
				// remove the UnderwritingDecision
				this.underwriter.decisions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Underwriter
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Underwriter/update/' + this.underwriter;

	return  this.http.post(uri_, this.underwriter );
}

	//********************************************************************
	// loadHelper - internal helper to load a Underwriter
	//********************************************************************	
	loadHelper( id ) {
		this.getUnderwriter(id)
			.subscribe((res : Underwriter) => {
				this.underwriter = res;
			});
	}
}