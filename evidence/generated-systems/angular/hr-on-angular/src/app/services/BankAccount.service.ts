import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {BankAccount} from '../models/BankAccount';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class BankAccountService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	bankAccount : BankAccount;

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
	// add a BankAccount
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addBankAccount(accountHolder, bankName, iban, bic, accountNumber, routingNumber) : Observable<any> {
		const uri_ = this.apiUrl + '/BankAccount/create';
		const obj = {
			      		accountHolder: accountHolder,
      		bankName: bankName,
      		iban: iban,
      		bic: bic,
      		accountNumber: accountNumber,
			routingNumber: routingNumber
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a BankAccount
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateBankAccount(accountHolder, bankName, iban, bic, accountNumber, routingNumber, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/BankAccount/update/' + id;
		const obj = {
				      		accountHolder: accountHolder,
      		bankName: bankName,
      		iban: iban,
      		bic: bic,
      		accountNumber: accountNumber,
			routingNumber: routingNumber
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a BankAccount
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteBankAccount(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/BankAccount/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a BankAccount
	// returns the results untouched as an Observable BankAccount
	// BankAccount model
	// delegates via URI
	//********************************************************************
	getBankAccount(id) : Observable<BankAccount> {
		const uri_ = this.apiUrl + '/BankAccount/load/' + id;

		return this.http.get<BankAccount>(uri_);
	}
	
	//********************************************************************
	// gets all BankAccount
	// returns the results untouched as JSON representation of an
	// Observable array of BankAccount models
	// delegates via URI
	//********************************************************************
	getBankAccounts() : Observable<BankAccount[]> {
		const uri_ = this.apiUrl + '/BankAccount/';

		return this
			.http.get<BankAccount[]>(uri_);
	}
	
		
	
	//********************************************************************
	// saveHelper - internal helper to save a BankAccount
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/BankAccount/update/' + this.bankAccount;

	return  this.http.post(uri_, this.bankAccount );
}

	//********************************************************************
	// loadHelper - internal helper to load a BankAccount
	//********************************************************************	
	loadHelper( id ) {
		this.getBankAccount(id)
			.subscribe((res : BankAccount) => {
				this.bankAccount = res;
			});
	}
}