import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {UnderwritingDecision} from '../models/UnderwritingDecision';
import {QuoteService} from '../services/Quote.service';
import {UnderwriterService} from '../services/Underwriter.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class UnderwritingDecisionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	underwritingDecision : UnderwritingDecision;

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
	// add a UnderwritingDecision
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addUnderwritingDecision(notes, decisionDate, Quote, Underwriter, Decision) : Observable<any> {
		const uri_ = this.apiUrl + '/UnderwritingDecision/create';
		const obj = {
			      		notes: notes,
      		decisionDate: decisionDate,
      		Quote: Quote != null && Quote.length > 0 ? Quote : null,
      		Underwriter: Underwriter != null && Underwriter.length > 0 ? Underwriter : null,
			Decision: Decision
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a UnderwritingDecision
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateUnderwritingDecision(notes, decisionDate, Quote, Underwriter, Decision, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/UnderwritingDecision/update/' + id;
		const obj = {
				      		notes: notes,
      		decisionDate: decisionDate,
      		Quote: Quote != null && Quote.length > 0 ? Quote : null,
      		Underwriter: Underwriter != null && Underwriter.length > 0 ? Underwriter : null,
			Decision: Decision
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a UnderwritingDecision
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteUnderwritingDecision(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/UnderwritingDecision/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a UnderwritingDecision
	// returns the results untouched as an Observable UnderwritingDecision
	// UnderwritingDecision model
	// delegates via URI
	//********************************************************************
	getUnderwritingDecision(id) : Observable<UnderwritingDecision> {
		const uri_ = this.apiUrl + '/UnderwritingDecision/load/' + id;

		return this.http.get<UnderwritingDecision>(uri_);
	}
	
	//********************************************************************
	// gets all UnderwritingDecision
	// returns the results untouched as JSON representation of an
	// Observable array of UnderwritingDecision models
	// delegates via URI
	//********************************************************************
	getUnderwritingDecisions() : Observable<UnderwritingDecision[]> {
		const uri_ = this.apiUrl + '/UnderwritingDecision/';

		return this
			.http.get<UnderwritingDecision[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Quote on a UnderwritingDecision
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignQuote( underwritingDecisionId, _quoteId ): Observable<any> {

		// get the UnderwritingDecision from storage
		this.loadHelper( underwritingDecisionId );

	// get the Quote from storage
	var tmp 	= new QuoteService(this.http).getQuote(_quoteId);

	// assign the Quote
	this.underwritingDecision.quote = tmp;

	// save the UnderwritingDecision
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Quote on a UnderwritingDecision
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignQuote( underwritingDecisionId ): Observable<any> {

		// get the UnderwritingDecision from storage
		this.loadHelper( underwritingDecisionId );

	// assign Quote to null
	this.underwritingDecision.quote = null;

	// save the UnderwritingDecision
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Underwriter on a UnderwritingDecision
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignUnderwriter( underwritingDecisionId, _underwriterId ): Observable<any> {

		// get the UnderwritingDecision from storage
		this.loadHelper( underwritingDecisionId );

	// get the Underwriter from storage
	var tmp 	= new UnderwriterService(this.http).getUnderwriter(_underwriterId);

	// assign the Underwriter
	this.underwritingDecision.underwriter = tmp;

	// save the UnderwritingDecision
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Underwriter on a UnderwritingDecision
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignUnderwriter( underwritingDecisionId ): Observable<any> {

		// get the UnderwritingDecision from storage
		this.loadHelper( underwritingDecisionId );

	// assign Underwriter to null
	this.underwritingDecision.underwriter = null;

	// save the UnderwritingDecision
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a UnderwritingDecision
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/UnderwritingDecision/update/' + this.underwritingDecision;

	return  this.http.post(uri_, this.underwritingDecision );
}

	//********************************************************************
	// loadHelper - internal helper to load a UnderwritingDecision
	//********************************************************************	
	loadHelper( id ) {
		this.getUnderwritingDecision(id)
			.subscribe((res : UnderwritingDecision) => {
				this.underwritingDecision = res;
			});
	}
}