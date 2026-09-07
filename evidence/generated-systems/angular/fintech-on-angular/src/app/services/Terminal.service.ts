import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Terminal} from '../models/Terminal';
import {MerchantService} from '../services/Merchant.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TerminalService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	terminal : Terminal;

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
	// add a Terminal
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTerminal(location, Merchant, Type, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Terminal/create';
		const obj = {
			      		location: location,
      		Merchant: Merchant != null && Merchant.length > 0 ? Merchant : null,
      		Type: Type,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Terminal
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTerminal(location, Merchant, Type, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Terminal/update/' + id;
		const obj = {
				      		location: location,
      		Merchant: Merchant != null && Merchant.length > 0 ? Merchant : null,
      		Type: Type,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Terminal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTerminal(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Terminal/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Terminal
	// returns the results untouched as an Observable Terminal
	// Terminal model
	// delegates via URI
	//********************************************************************
	getTerminal(id) : Observable<Terminal> {
		const uri_ = this.apiUrl + '/Terminal/load/' + id;

		return this.http.get<Terminal>(uri_);
	}
	
	//********************************************************************
	// gets all Terminal
	// returns the results untouched as JSON representation of an
	// Observable array of Terminal models
	// delegates via URI
	//********************************************************************
	getTerminals() : Observable<Terminal[]> {
		const uri_ = this.apiUrl + '/Terminal/';

		return this
			.http.get<Terminal[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Merchant on a Terminal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignMerchant( terminalId, _merchantId ): Observable<any> {

		// get the Terminal from storage
		this.loadHelper( terminalId );

	// get the Merchant from storage
	var tmp 	= new MerchantService(this.http).getMerchant(_merchantId);

	// assign the Merchant
	this.terminal.merchant = tmp;

	// save the Terminal
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Merchant on a Terminal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignMerchant( terminalId ): Observable<any> {

		// get the Terminal from storage
		this.loadHelper( terminalId );

	// assign Merchant to null
	this.terminal.merchant = null;

	// save the Terminal
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Terminal
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Terminal/update/' + this.terminal;

	return  this.http.post(uri_, this.terminal );
}

	//********************************************************************
	// loadHelper - internal helper to load a Terminal
	//********************************************************************	
	loadHelper( id ) {
		this.getTerminal(id)
			.subscribe((res : Terminal) => {
				this.terminal = res;
			});
	}
}