import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {DirectDebitMandate} from '../models/DirectDebitMandate';
import {AccountService} from '../services/Account.service';
import {CreditorService} from '../services/Creditor.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DirectDebitMandateService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	directDebitMandate : DirectDebitMandate;

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
	// add a DirectDebitMandate
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDirectDebitMandate(mandateId, signedAt, Account, Creditor, Scheme, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/DirectDebitMandate/create';
		const obj = {
			      		mandateId: mandateId,
      		signedAt: signedAt,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Creditor: Creditor != null && Creditor.length > 0 ? Creditor : null,
      		Scheme: Scheme,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a DirectDebitMandate
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDirectDebitMandate(mandateId, signedAt, Account, Creditor, Scheme, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/DirectDebitMandate/update/' + id;
		const obj = {
				      		mandateId: mandateId,
      		signedAt: signedAt,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Creditor: Creditor != null && Creditor.length > 0 ? Creditor : null,
      		Scheme: Scheme,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a DirectDebitMandate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDirectDebitMandate(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/DirectDebitMandate/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a DirectDebitMandate
	// returns the results untouched as an Observable DirectDebitMandate
	// DirectDebitMandate model
	// delegates via URI
	//********************************************************************
	getDirectDebitMandate(id) : Observable<DirectDebitMandate> {
		const uri_ = this.apiUrl + '/DirectDebitMandate/load/' + id;

		return this.http.get<DirectDebitMandate>(uri_);
	}
	
	//********************************************************************
	// gets all DirectDebitMandate
	// returns the results untouched as JSON representation of an
	// Observable array of DirectDebitMandate models
	// delegates via URI
	//********************************************************************
	getDirectDebitMandates() : Observable<DirectDebitMandate[]> {
		const uri_ = this.apiUrl + '/DirectDebitMandate/';

		return this
			.http.get<DirectDebitMandate[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Account on a DirectDebitMandate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAccount( directDebitMandateId, _accountId ): Observable<any> {

		// get the DirectDebitMandate from storage
		this.loadHelper( directDebitMandateId );

	// get the Account from storage
	var tmp 	= new AccountService(this.http).getAccount(_accountId);

	// assign the Account
	this.directDebitMandate.account = tmp;

	// save the DirectDebitMandate
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Account on a DirectDebitMandate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAccount( directDebitMandateId ): Observable<any> {

		// get the DirectDebitMandate from storage
		this.loadHelper( directDebitMandateId );

	// assign Account to null
	this.directDebitMandate.account = null;

	// save the DirectDebitMandate
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Creditor on a DirectDebitMandate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCreditor( directDebitMandateId, _creditorId ): Observable<any> {

		// get the DirectDebitMandate from storage
		this.loadHelper( directDebitMandateId );

	// get the Creditor from storage
	var tmp 	= new CreditorService(this.http).getCreditor(_creditorId);

	// assign the Creditor
	this.directDebitMandate.creditor = tmp;

	// save the DirectDebitMandate
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Creditor on a DirectDebitMandate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCreditor( directDebitMandateId ): Observable<any> {

		// get the DirectDebitMandate from storage
		this.loadHelper( directDebitMandateId );

	// assign Creditor to null
	this.directDebitMandate.creditor = null;

	// save the DirectDebitMandate
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a DirectDebitMandate
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/DirectDebitMandate/update/' + this.directDebitMandate;

	return  this.http.post(uri_, this.directDebitMandate );
}

	//********************************************************************
	// loadHelper - internal helper to load a DirectDebitMandate
	//********************************************************************	
	loadHelper( id ) {
		this.getDirectDebitMandate(id)
			.subscribe((res : DirectDebitMandate) => {
				this.directDebitMandate = res;
			});
	}
}