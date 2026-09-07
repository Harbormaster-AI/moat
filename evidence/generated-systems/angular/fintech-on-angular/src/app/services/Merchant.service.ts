import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Merchant} from '../models/Merchant';
import {TerminalService} from '../services/Terminal.service';
import {PaymentContractService} from '../services/PaymentContract.service';
import {PayoutService} from '../services/Payout.service';
import {SettlementBatchService} from '../services/SettlementBatch.service';
import {DisputeService} from '../services/Dispute.service';
import {InvoiceService} from '../services/Invoice.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class MerchantService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	merchant : Merchant;

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
	// add a Merchant
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addMerchant(name, mcc, url, country, settlementCurrency, Terminals, PaymentContracts, Payouts, Settlements, Disputes, Invoices) : Observable<any> {
		const uri_ = this.apiUrl + '/Merchant/create';
		const obj = {
			      		name: name,
      		mcc: mcc,
      		url: url,
      		country: country,
      		settlementCurrency: settlementCurrency,
      		Terminals: Terminals != null && Terminals.length > 0 ? Terminals : null,
      		PaymentContracts: PaymentContracts != null && PaymentContracts.length > 0 ? PaymentContracts : null,
      		Payouts: Payouts != null && Payouts.length > 0 ? Payouts : null,
      		Settlements: Settlements != null && Settlements.length > 0 ? Settlements : null,
      		Disputes: Disputes != null && Disputes.length > 0 ? Disputes : null,
			Invoices: Invoices != null && Invoices.length > 0 ? Invoices : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Merchant
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateMerchant(name, mcc, url, country, settlementCurrency, Terminals, PaymentContracts, Payouts, Settlements, Disputes, Invoices, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Merchant/update/' + id;
		const obj = {
				      		name: name,
      		mcc: mcc,
      		url: url,
      		country: country,
      		settlementCurrency: settlementCurrency,
      		Terminals: Terminals != null && Terminals.length > 0 ? Terminals : null,
      		PaymentContracts: PaymentContracts != null && PaymentContracts.length > 0 ? PaymentContracts : null,
      		Payouts: Payouts != null && Payouts.length > 0 ? Payouts : null,
      		Settlements: Settlements != null && Settlements.length > 0 ? Settlements : null,
      		Disputes: Disputes != null && Disputes.length > 0 ? Disputes : null,
			Invoices: Invoices != null && Invoices.length > 0 ? Invoices : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Merchant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteMerchant(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Merchant/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Merchant
	// returns the results untouched as an Observable Merchant
	// Merchant model
	// delegates via URI
	//********************************************************************
	getMerchant(id) : Observable<Merchant> {
		const uri_ = this.apiUrl + '/Merchant/load/' + id;

		return this.http.get<Merchant>(uri_);
	}
	
	//********************************************************************
	// gets all Merchant
	// returns the results untouched as JSON representation of an
	// Observable array of Merchant models
	// delegates via URI
	//********************************************************************
	getMerchants() : Observable<Merchant[]> {
		const uri_ = this.apiUrl + '/Merchant/';

		return this
			.http.get<Merchant[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more terminalsIds as a Terminals
	// to a Merchant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTerminals( merchantId, terminalsIds ): Observable<any> {

		// get the Merchant
		this.loadHelper( merchantId );

	// split on a comma with no spaces
	var idList = terminalsIds.split(',')

	// iterate over array of terminals ids
	idList.forEach(function (id) {
		// read the Terminal
		var terminal = new TerminalService(this.http).getTerminal(id);
		// add the Terminal if not already assigned
		if ( this.merchant.terminals.indexOf(terminal) == -1 )
		this.merchant.terminals.push(terminal);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more terminalsIds as a Terminals
	// from a Merchant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTerminals( merchantId, terminalsIds ): Observable<any> {

		// get the Merchant
		this.loadHelper( merchantId );


	// split on a comma with no spaces
	var idList 					= terminalsIds.split(',');
	var terminals 	= this.merchant.terminals;

	if ( terminals != null && terminalsIds != null ) {

		// iterate over array of terminals ids
		terminals.forEach(function (obj) {
			if ( terminalsIds.indexOf(obj._id) > -1 ) {
				// remove the Terminal
				this.merchant.terminals.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more paymentContractsIds as a PaymentContracts
	// to a Merchant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPaymentContracts( merchantId, paymentContractsIds ): Observable<any> {

		// get the Merchant
		this.loadHelper( merchantId );

	// split on a comma with no spaces
	var idList = paymentContractsIds.split(',')

	// iterate over array of paymentContracts ids
	idList.forEach(function (id) {
		// read the PaymentContract
		var paymentContract = new PaymentContractService(this.http).getPaymentContract(id);
		// add the PaymentContract if not already assigned
		if ( this.merchant.paymentContracts.indexOf(paymentContract) == -1 )
		this.merchant.paymentContracts.push(paymentContract);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more paymentContractsIds as a PaymentContracts
	// from a Merchant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePaymentContracts( merchantId, paymentContractsIds ): Observable<any> {

		// get the Merchant
		this.loadHelper( merchantId );


	// split on a comma with no spaces
	var idList 					= paymentContractsIds.split(',');
	var paymentContracts 	= this.merchant.paymentContracts;

	if ( paymentContracts != null && paymentContractsIds != null ) {

		// iterate over array of paymentContracts ids
		paymentContracts.forEach(function (obj) {
			if ( paymentContractsIds.indexOf(obj._id) > -1 ) {
				// remove the PaymentContract
				this.merchant.paymentContracts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more payoutsIds as a Payouts
	// to a Merchant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPayouts( merchantId, payoutsIds ): Observable<any> {

		// get the Merchant
		this.loadHelper( merchantId );

	// split on a comma with no spaces
	var idList = payoutsIds.split(',')

	// iterate over array of payouts ids
	idList.forEach(function (id) {
		// read the Payout
		var payout = new PayoutService(this.http).getPayout(id);
		// add the Payout if not already assigned
		if ( this.merchant.payouts.indexOf(payout) == -1 )
		this.merchant.payouts.push(payout);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more payoutsIds as a Payouts
	// from a Merchant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePayouts( merchantId, payoutsIds ): Observable<any> {

		// get the Merchant
		this.loadHelper( merchantId );


	// split on a comma with no spaces
	var idList 					= payoutsIds.split(',');
	var payouts 	= this.merchant.payouts;

	if ( payouts != null && payoutsIds != null ) {

		// iterate over array of payouts ids
		payouts.forEach(function (obj) {
			if ( payoutsIds.indexOf(obj._id) > -1 ) {
				// remove the Payout
				this.merchant.payouts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more settlementsIds as a Settlements
	// to a Merchant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSettlements( merchantId, settlementsIds ): Observable<any> {

		// get the Merchant
		this.loadHelper( merchantId );

	// split on a comma with no spaces
	var idList = settlementsIds.split(',')

	// iterate over array of settlements ids
	idList.forEach(function (id) {
		// read the SettlementBatch
		var settlementBatch = new SettlementBatchService(this.http).getSettlementBatch(id);
		// add the SettlementBatch if not already assigned
		if ( this.merchant.settlements.indexOf(settlementBatch) == -1 )
		this.merchant.settlements.push(settlementBatch);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more settlementsIds as a Settlements
	// from a Merchant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSettlements( merchantId, settlementsIds ): Observable<any> {

		// get the Merchant
		this.loadHelper( merchantId );


	// split on a comma with no spaces
	var idList 					= settlementsIds.split(',');
	var settlements 	= this.merchant.settlements;

	if ( settlements != null && settlementsIds != null ) {

		// iterate over array of settlements ids
		settlements.forEach(function (obj) {
			if ( settlementsIds.indexOf(obj._id) > -1 ) {
				// remove the SettlementBatch
				this.merchant.settlements.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more disputesIds as a Disputes
	// to a Merchant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDisputes( merchantId, disputesIds ): Observable<any> {

		// get the Merchant
		this.loadHelper( merchantId );

	// split on a comma with no spaces
	var idList = disputesIds.split(',')

	// iterate over array of disputes ids
	idList.forEach(function (id) {
		// read the Dispute
		var dispute = new DisputeService(this.http).getDispute(id);
		// add the Dispute if not already assigned
		if ( this.merchant.disputes.indexOf(dispute) == -1 )
		this.merchant.disputes.push(dispute);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more disputesIds as a Disputes
	// from a Merchant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDisputes( merchantId, disputesIds ): Observable<any> {

		// get the Merchant
		this.loadHelper( merchantId );


	// split on a comma with no spaces
	var idList 					= disputesIds.split(',');
	var disputes 	= this.merchant.disputes;

	if ( disputes != null && disputesIds != null ) {

		// iterate over array of disputes ids
		disputes.forEach(function (obj) {
			if ( disputesIds.indexOf(obj._id) > -1 ) {
				// remove the Dispute
				this.merchant.disputes.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more invoicesIds as a Invoices
	// to a Merchant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInvoices( merchantId, invoicesIds ): Observable<any> {

		// get the Merchant
		this.loadHelper( merchantId );

	// split on a comma with no spaces
	var idList = invoicesIds.split(',')

	// iterate over array of invoices ids
	idList.forEach(function (id) {
		// read the Invoice
		var invoice = new InvoiceService(this.http).getInvoice(id);
		// add the Invoice if not already assigned
		if ( this.merchant.invoices.indexOf(invoice) == -1 )
		this.merchant.invoices.push(invoice);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more invoicesIds as a Invoices
	// from a Merchant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInvoices( merchantId, invoicesIds ): Observable<any> {

		// get the Merchant
		this.loadHelper( merchantId );


	// split on a comma with no spaces
	var idList 					= invoicesIds.split(',');
	var invoices 	= this.merchant.invoices;

	if ( invoices != null && invoicesIds != null ) {

		// iterate over array of invoices ids
		invoices.forEach(function (obj) {
			if ( invoicesIds.indexOf(obj._id) > -1 ) {
				// remove the Invoice
				this.merchant.invoices.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Merchant
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Merchant/update/' + this.merchant;

	return  this.http.post(uri_, this.merchant );
}

	//********************************************************************
	// loadHelper - internal helper to load a Merchant
	//********************************************************************	
	loadHelper( id ) {
		this.getMerchant(id)
			.subscribe((res : Merchant) => {
				this.merchant = res;
			});
	}
}