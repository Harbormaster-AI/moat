import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PaymentProcessor} from '../models/PaymentProcessor';
import {FinancialInstitutionService} from '../services/FinancialInstitution.service';
import {PaymentContractService} from '../services/PaymentContract.service';
import {SettlementBatchService} from '../services/SettlementBatch.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PaymentProcessorService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	paymentProcessor : PaymentProcessor;

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
	// add a PaymentProcessor
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPaymentProcessor(name, processorCode, networkSupport, Institutions, Contracts, Settlements) : Observable<any> {
		const uri_ = this.apiUrl + '/PaymentProcessor/create';
		const obj = {
			      		name: name,
      		processorCode: processorCode,
      		networkSupport: networkSupport,
      		Institutions: Institutions != null && Institutions.length > 0 ? Institutions : null,
      		Contracts: Contracts != null && Contracts.length > 0 ? Contracts : null,
			Settlements: Settlements != null && Settlements.length > 0 ? Settlements : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PaymentProcessor
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePaymentProcessor(name, processorCode, networkSupport, Institutions, Contracts, Settlements, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PaymentProcessor/update/' + id;
		const obj = {
				      		name: name,
      		processorCode: processorCode,
      		networkSupport: networkSupport,
      		Institutions: Institutions != null && Institutions.length > 0 ? Institutions : null,
      		Contracts: Contracts != null && Contracts.length > 0 ? Contracts : null,
			Settlements: Settlements != null && Settlements.length > 0 ? Settlements : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PaymentProcessor
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePaymentProcessor(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PaymentProcessor/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PaymentProcessor
	// returns the results untouched as an Observable PaymentProcessor
	// PaymentProcessor model
	// delegates via URI
	//********************************************************************
	getPaymentProcessor(id) : Observable<PaymentProcessor> {
		const uri_ = this.apiUrl + '/PaymentProcessor/load/' + id;

		return this.http.get<PaymentProcessor>(uri_);
	}
	
	//********************************************************************
	// gets all PaymentProcessor
	// returns the results untouched as JSON representation of an
	// Observable array of PaymentProcessor models
	// delegates via URI
	//********************************************************************
	getPaymentProcessors() : Observable<PaymentProcessor[]> {
		const uri_ = this.apiUrl + '/PaymentProcessor/';

		return this
			.http.get<PaymentProcessor[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more institutionsIds as a Institutions
	// to a PaymentProcessor
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInstitutions( paymentProcessorId, institutionsIds ): Observable<any> {

		// get the PaymentProcessor
		this.loadHelper( paymentProcessorId );

	// split on a comma with no spaces
	var idList = institutionsIds.split(',')

	// iterate over array of institutions ids
	idList.forEach(function (id) {
		// read the FinancialInstitution
		var financialInstitution = new FinancialInstitutionService(this.http).getFinancialInstitution(id);
		// add the FinancialInstitution if not already assigned
		if ( this.paymentProcessor.institutions.indexOf(financialInstitution) == -1 )
		this.paymentProcessor.institutions.push(financialInstitution);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more institutionsIds as a Institutions
	// from a PaymentProcessor
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInstitutions( paymentProcessorId, institutionsIds ): Observable<any> {

		// get the PaymentProcessor
		this.loadHelper( paymentProcessorId );


	// split on a comma with no spaces
	var idList 					= institutionsIds.split(',');
	var institutions 	= this.paymentProcessor.institutions;

	if ( institutions != null && institutionsIds != null ) {

		// iterate over array of institutions ids
		institutions.forEach(function (obj) {
			if ( institutionsIds.indexOf(obj._id) > -1 ) {
				// remove the FinancialInstitution
				this.paymentProcessor.institutions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more contractsIds as a Contracts
	// to a PaymentProcessor
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addContracts( paymentProcessorId, contractsIds ): Observable<any> {

		// get the PaymentProcessor
		this.loadHelper( paymentProcessorId );

	// split on a comma with no spaces
	var idList = contractsIds.split(',')

	// iterate over array of contracts ids
	idList.forEach(function (id) {
		// read the PaymentContract
		var paymentContract = new PaymentContractService(this.http).getPaymentContract(id);
		// add the PaymentContract if not already assigned
		if ( this.paymentProcessor.contracts.indexOf(paymentContract) == -1 )
		this.paymentProcessor.contracts.push(paymentContract);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more contractsIds as a Contracts
	// from a PaymentProcessor
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeContracts( paymentProcessorId, contractsIds ): Observable<any> {

		// get the PaymentProcessor
		this.loadHelper( paymentProcessorId );


	// split on a comma with no spaces
	var idList 					= contractsIds.split(',');
	var contracts 	= this.paymentProcessor.contracts;

	if ( contracts != null && contractsIds != null ) {

		// iterate over array of contracts ids
		contracts.forEach(function (obj) {
			if ( contractsIds.indexOf(obj._id) > -1 ) {
				// remove the PaymentContract
				this.paymentProcessor.contracts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more settlementsIds as a Settlements
	// to a PaymentProcessor
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSettlements( paymentProcessorId, settlementsIds ): Observable<any> {

		// get the PaymentProcessor
		this.loadHelper( paymentProcessorId );

	// split on a comma with no spaces
	var idList = settlementsIds.split(',')

	// iterate over array of settlements ids
	idList.forEach(function (id) {
		// read the SettlementBatch
		var settlementBatch = new SettlementBatchService(this.http).getSettlementBatch(id);
		// add the SettlementBatch if not already assigned
		if ( this.paymentProcessor.settlements.indexOf(settlementBatch) == -1 )
		this.paymentProcessor.settlements.push(settlementBatch);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more settlementsIds as a Settlements
	// from a PaymentProcessor
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSettlements( paymentProcessorId, settlementsIds ): Observable<any> {

		// get the PaymentProcessor
		this.loadHelper( paymentProcessorId );


	// split on a comma with no spaces
	var idList 					= settlementsIds.split(',');
	var settlements 	= this.paymentProcessor.settlements;

	if ( settlements != null && settlementsIds != null ) {

		// iterate over array of settlements ids
		settlements.forEach(function (obj) {
			if ( settlementsIds.indexOf(obj._id) > -1 ) {
				// remove the SettlementBatch
				this.paymentProcessor.settlements.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a PaymentProcessor
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PaymentProcessor/update/' + this.paymentProcessor;

	return  this.http.post(uri_, this.paymentProcessor );
}

	//********************************************************************
	// loadHelper - internal helper to load a PaymentProcessor
	//********************************************************************	
	loadHelper( id ) {
		this.getPaymentProcessor(id)
			.subscribe((res : PaymentProcessor) => {
				this.paymentProcessor = res;
			});
	}
}