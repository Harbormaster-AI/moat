import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Quote} from '../models/Quote';
import {ApplicationService} from '../services/Application.service';
import {UnderwritingDecisionService} from '../services/UnderwritingDecision.service';
import {PolicyService} from '../services/Policy.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class QuoteService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	quote : Quote;

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
	// add a Quote
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addQuote(quoteNumber, totalPremium, ratingDate, asBound, Application, UnderwritingDecisions, Policy) : Observable<any> {
		const uri_ = this.apiUrl + '/Quote/create';
		const obj = {
			      		quoteNumber: quoteNumber,
      		totalPremium: totalPremium,
      		ratingDate: ratingDate,
      		asBound: asBound,
      		Application: Application != null && Application.length > 0 ? Application : null,
      		UnderwritingDecisions: UnderwritingDecisions != null && UnderwritingDecisions.length > 0 ? UnderwritingDecisions : null,
			Policy: Policy != null && Policy.length > 0 ? Policy : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateQuote(quoteNumber, totalPremium, ratingDate, asBound, Application, UnderwritingDecisions, Policy, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Quote/update/' + id;
		const obj = {
				      		quoteNumber: quoteNumber,
      		totalPremium: totalPremium,
      		ratingDate: ratingDate,
      		asBound: asBound,
      		Application: Application != null && Application.length > 0 ? Application : null,
      		UnderwritingDecisions: UnderwritingDecisions != null && UnderwritingDecisions.length > 0 ? UnderwritingDecisions : null,
			Policy: Policy != null && Policy.length > 0 ? Policy : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteQuote(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Quote/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Quote
	// returns the results untouched as an Observable Quote
	// Quote model
	// delegates via URI
	//********************************************************************
	getQuote(id) : Observable<Quote> {
		const uri_ = this.apiUrl + '/Quote/load/' + id;

		return this.http.get<Quote>(uri_);
	}
	
	//********************************************************************
	// gets all Quote
	// returns the results untouched as JSON representation of an
	// Observable array of Quote models
	// delegates via URI
	//********************************************************************
	getQuotes() : Observable<Quote[]> {
		const uri_ = this.apiUrl + '/Quote/';

		return this
			.http.get<Quote[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Application on a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignApplication( quoteId, _applicationId ): Observable<any> {

		// get the Quote from storage
		this.loadHelper( quoteId );

	// get the Application from storage
	var tmp 	= new ApplicationService(this.http).getApplication(_applicationId);

	// assign the Application
	this.quote.application = tmp;

	// save the Quote
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Application on a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignApplication( quoteId ): Observable<any> {

		// get the Quote from storage
		this.loadHelper( quoteId );

	// assign Application to null
	this.quote.application = null;

	// save the Quote
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Policy on a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPolicy( quoteId, _policyId ): Observable<any> {

		// get the Quote from storage
		this.loadHelper( quoteId );

	// get the Policy from storage
	var tmp 	= new PolicyService(this.http).getPolicy(_policyId);

	// assign the Policy
	this.quote.policy = tmp;

	// save the Quote
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Policy on a Quote
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPolicy( quoteId ): Observable<any> {

		// get the Quote from storage
		this.loadHelper( quoteId );

	// assign Policy to null
	this.quote.policy = null;

	// save the Quote
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more underwritingDecisionsIds as a UnderwritingDecisions
	// to a Quote
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addUnderwritingDecisions( quoteId, underwritingDecisionsIds ): Observable<any> {

		// get the Quote
		this.loadHelper( quoteId );

	// split on a comma with no spaces
	var idList = underwritingDecisionsIds.split(',')

	// iterate over array of underwritingDecisions ids
	idList.forEach(function (id) {
		// read the UnderwritingDecision
		var underwritingDecision = new UnderwritingDecisionService(this.http).getUnderwritingDecision(id);
		// add the UnderwritingDecision if not already assigned
		if ( this.quote.underwritingDecisions.indexOf(underwritingDecision) == -1 )
		this.quote.underwritingDecisions.push(underwritingDecision);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more underwritingDecisionsIds as a UnderwritingDecisions
	// from a Quote
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeUnderwritingDecisions( quoteId, underwritingDecisionsIds ): Observable<any> {

		// get the Quote
		this.loadHelper( quoteId );


	// split on a comma with no spaces
	var idList 					= underwritingDecisionsIds.split(',');
	var underwritingDecisions 	= this.quote.underwritingDecisions;

	if ( underwritingDecisions != null && underwritingDecisionsIds != null ) {

		// iterate over array of underwritingDecisions ids
		underwritingDecisions.forEach(function (obj) {
			if ( underwritingDecisionsIds.indexOf(obj._id) > -1 ) {
				// remove the UnderwritingDecision
				this.quote.underwritingDecisions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Quote
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Quote/update/' + this.quote;

	return  this.http.post(uri_, this.quote );
}

	//********************************************************************
	// loadHelper - internal helper to load a Quote
	//********************************************************************	
	loadHelper( id ) {
		this.getQuote(id)
			.subscribe((res : Quote) => {
				this.quote = res;
			});
	}
}