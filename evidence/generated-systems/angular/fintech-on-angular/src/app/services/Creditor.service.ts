import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Creditor} from '../models/Creditor';
import {DirectDebitMandateService} from '../services/DirectDebitMandate.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CreditorService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	creditor : Creditor;

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
	// add a Creditor
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCreditor(name, bic, address, Mandates) : Observable<any> {
		const uri_ = this.apiUrl + '/Creditor/create';
		const obj = {
			      		name: name,
      		bic: bic,
      		address: address,
			Mandates: Mandates != null && Mandates.length > 0 ? Mandates : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Creditor
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCreditor(name, bic, address, Mandates, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Creditor/update/' + id;
		const obj = {
				      		name: name,
      		bic: bic,
      		address: address,
			Mandates: Mandates != null && Mandates.length > 0 ? Mandates : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Creditor
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCreditor(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Creditor/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Creditor
	// returns the results untouched as an Observable Creditor
	// Creditor model
	// delegates via URI
	//********************************************************************
	getCreditor(id) : Observable<Creditor> {
		const uri_ = this.apiUrl + '/Creditor/load/' + id;

		return this.http.get<Creditor>(uri_);
	}
	
	//********************************************************************
	// gets all Creditor
	// returns the results untouched as JSON representation of an
	// Observable array of Creditor models
	// delegates via URI
	//********************************************************************
	getCreditors() : Observable<Creditor[]> {
		const uri_ = this.apiUrl + '/Creditor/';

		return this
			.http.get<Creditor[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more mandatesIds as a Mandates
	// to a Creditor
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMandates( creditorId, mandatesIds ): Observable<any> {

		// get the Creditor
		this.loadHelper( creditorId );

	// split on a comma with no spaces
	var idList = mandatesIds.split(',')

	// iterate over array of mandates ids
	idList.forEach(function (id) {
		// read the DirectDebitMandate
		var directDebitMandate = new DirectDebitMandateService(this.http).getDirectDebitMandate(id);
		// add the DirectDebitMandate if not already assigned
		if ( this.creditor.mandates.indexOf(directDebitMandate) == -1 )
		this.creditor.mandates.push(directDebitMandate);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more mandatesIds as a Mandates
	// from a Creditor
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMandates( creditorId, mandatesIds ): Observable<any> {

		// get the Creditor
		this.loadHelper( creditorId );


	// split on a comma with no spaces
	var idList 					= mandatesIds.split(',');
	var mandates 	= this.creditor.mandates;

	if ( mandates != null && mandatesIds != null ) {

		// iterate over array of mandates ids
		mandates.forEach(function (obj) {
			if ( mandatesIds.indexOf(obj._id) > -1 ) {
				// remove the DirectDebitMandate
				this.creditor.mandates.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Creditor
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Creditor/update/' + this.creditor;

	return  this.http.post(uri_, this.creditor );
}

	//********************************************************************
	// loadHelper - internal helper to load a Creditor
	//********************************************************************	
	loadHelper( id ) {
		this.getCreditor(id)
			.subscribe((res : Creditor) => {
				this.creditor = res;
			});
	}
}