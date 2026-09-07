import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PaymentOrder} from '../models/PaymentOrder';
import {AccountService} from '../services/Account.service';
import {BeneficiaryService} from '../services/Beneficiary.service';
import {TransactionService} from '../services/Transaction.service';
import {FXDealService} from '../services/FXDeal.service';
import {AppliedFeeService} from '../services/AppliedFee.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PaymentOrderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	paymentOrder : PaymentOrder;

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
	// add a PaymentOrder
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPaymentOrder(orderReference, requestedExecutionDate, purpose, SourceAccount, DestinationAccount, Beneficiary, Transactions, FxDeal, Fees, PaymentMethod, Status, Priority) : Observable<any> {
		const uri_ = this.apiUrl + '/PaymentOrder/create';
		const obj = {
			      		orderReference: orderReference,
      		requestedExecutionDate: requestedExecutionDate,
      		purpose: purpose,
      		SourceAccount: SourceAccount != null && SourceAccount.length > 0 ? SourceAccount : null,
      		DestinationAccount: DestinationAccount != null && DestinationAccount.length > 0 ? DestinationAccount : null,
      		Beneficiary: Beneficiary != null && Beneficiary.length > 0 ? Beneficiary : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
      		FxDeal: FxDeal != null && FxDeal.length > 0 ? FxDeal : null,
      		Fees: Fees != null && Fees.length > 0 ? Fees : null,
      		PaymentMethod: PaymentMethod,
      		Status: Status,
			Priority: Priority
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PaymentOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePaymentOrder(orderReference, requestedExecutionDate, purpose, SourceAccount, DestinationAccount, Beneficiary, Transactions, FxDeal, Fees, PaymentMethod, Status, Priority, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PaymentOrder/update/' + id;
		const obj = {
				      		orderReference: orderReference,
      		requestedExecutionDate: requestedExecutionDate,
      		purpose: purpose,
      		SourceAccount: SourceAccount != null && SourceAccount.length > 0 ? SourceAccount : null,
      		DestinationAccount: DestinationAccount != null && DestinationAccount.length > 0 ? DestinationAccount : null,
      		Beneficiary: Beneficiary != null && Beneficiary.length > 0 ? Beneficiary : null,
      		Transactions: Transactions != null && Transactions.length > 0 ? Transactions : null,
      		FxDeal: FxDeal != null && FxDeal.length > 0 ? FxDeal : null,
      		Fees: Fees != null && Fees.length > 0 ? Fees : null,
      		PaymentMethod: PaymentMethod,
      		Status: Status,
			Priority: Priority
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PaymentOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePaymentOrder(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PaymentOrder/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PaymentOrder
	// returns the results untouched as an Observable PaymentOrder
	// PaymentOrder model
	// delegates via URI
	//********************************************************************
	getPaymentOrder(id) : Observable<PaymentOrder> {
		const uri_ = this.apiUrl + '/PaymentOrder/load/' + id;

		return this.http.get<PaymentOrder>(uri_);
	}
	
	//********************************************************************
	// gets all PaymentOrder
	// returns the results untouched as JSON representation of an
	// Observable array of PaymentOrder models
	// delegates via URI
	//********************************************************************
	getPaymentOrders() : Observable<PaymentOrder[]> {
		const uri_ = this.apiUrl + '/PaymentOrder/';

		return this
			.http.get<PaymentOrder[]>(uri_);
	}
	
			//********************************************************************
	// assigns a SourceAccount on a PaymentOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSourceAccount( paymentOrderId, _sourceAccountId ): Observable<any> {

		// get the PaymentOrder from storage
		this.loadHelper( paymentOrderId );

	// get the Account from storage
	var tmp 	= new AccountService(this.http).getAccount(_sourceAccountId);

	// assign the SourceAccount
	this.paymentOrder.sourceAccount = tmp;

	// save the PaymentOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a SourceAccount on a PaymentOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSourceAccount( paymentOrderId ): Observable<any> {

		// get the PaymentOrder from storage
		this.loadHelper( paymentOrderId );

	// assign SourceAccount to null
	this.paymentOrder.sourceAccount = null;

	// save the PaymentOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a DestinationAccount on a PaymentOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDestinationAccount( paymentOrderId, _destinationAccountId ): Observable<any> {

		// get the PaymentOrder from storage
		this.loadHelper( paymentOrderId );

	// get the Account from storage
	var tmp 	= new AccountService(this.http).getAccount(_destinationAccountId);

	// assign the DestinationAccount
	this.paymentOrder.destinationAccount = tmp;

	// save the PaymentOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a DestinationAccount on a PaymentOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDestinationAccount( paymentOrderId ): Observable<any> {

		// get the PaymentOrder from storage
		this.loadHelper( paymentOrderId );

	// assign DestinationAccount to null
	this.paymentOrder.destinationAccount = null;

	// save the PaymentOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Beneficiary on a PaymentOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignBeneficiary( paymentOrderId, _beneficiaryId ): Observable<any> {

		// get the PaymentOrder from storage
		this.loadHelper( paymentOrderId );

	// get the Beneficiary from storage
	var tmp 	= new BeneficiaryService(this.http).getBeneficiary(_beneficiaryId);

	// assign the Beneficiary
	this.paymentOrder.beneficiary = tmp;

	// save the PaymentOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Beneficiary on a PaymentOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignBeneficiary( paymentOrderId ): Observable<any> {

		// get the PaymentOrder from storage
		this.loadHelper( paymentOrderId );

	// assign Beneficiary to null
	this.paymentOrder.beneficiary = null;

	// save the PaymentOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a FxDeal on a PaymentOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignFxDeal( paymentOrderId, _fxDealId ): Observable<any> {

		// get the PaymentOrder from storage
		this.loadHelper( paymentOrderId );

	// get the FXDeal from storage
	var tmp 	= new FXDealService(this.http).getFXDeal(_fxDealId);

	// assign the FxDeal
	this.paymentOrder.fxDeal = tmp;

	// save the PaymentOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a FxDeal on a PaymentOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignFxDeal( paymentOrderId ): Observable<any> {

		// get the PaymentOrder from storage
		this.loadHelper( paymentOrderId );

	// assign FxDeal to null
	this.paymentOrder.fxDeal = null;

	// save the PaymentOrder
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more transactionsIds as a Transactions
	// to a PaymentOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTransactions( paymentOrderId, transactionsIds ): Observable<any> {

		// get the PaymentOrder
		this.loadHelper( paymentOrderId );

	// split on a comma with no spaces
	var idList = transactionsIds.split(',')

	// iterate over array of transactions ids
	idList.forEach(function (id) {
		// read the Transaction
		var transaction = new TransactionService(this.http).getTransaction(id);
		// add the Transaction if not already assigned
		if ( this.paymentOrder.transactions.indexOf(transaction) == -1 )
		this.paymentOrder.transactions.push(transaction);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more transactionsIds as a Transactions
	// from a PaymentOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTransactions( paymentOrderId, transactionsIds ): Observable<any> {

		// get the PaymentOrder
		this.loadHelper( paymentOrderId );


	// split on a comma with no spaces
	var idList 					= transactionsIds.split(',');
	var transactions 	= this.paymentOrder.transactions;

	if ( transactions != null && transactionsIds != null ) {

		// iterate over array of transactions ids
		transactions.forEach(function (obj) {
			if ( transactionsIds.indexOf(obj._id) > -1 ) {
				// remove the Transaction
				this.paymentOrder.transactions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more feesIds as a Fees
	// to a PaymentOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addFees( paymentOrderId, feesIds ): Observable<any> {

		// get the PaymentOrder
		this.loadHelper( paymentOrderId );

	// split on a comma with no spaces
	var idList = feesIds.split(',')

	// iterate over array of fees ids
	idList.forEach(function (id) {
		// read the AppliedFee
		var appliedFee = new AppliedFeeService(this.http).getAppliedFee(id);
		// add the AppliedFee if not already assigned
		if ( this.paymentOrder.fees.indexOf(appliedFee) == -1 )
		this.paymentOrder.fees.push(appliedFee);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more feesIds as a Fees
	// from a PaymentOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeFees( paymentOrderId, feesIds ): Observable<any> {

		// get the PaymentOrder
		this.loadHelper( paymentOrderId );


	// split on a comma with no spaces
	var idList 					= feesIds.split(',');
	var fees 	= this.paymentOrder.fees;

	if ( fees != null && feesIds != null ) {

		// iterate over array of fees ids
		fees.forEach(function (obj) {
			if ( feesIds.indexOf(obj._id) > -1 ) {
				// remove the AppliedFee
				this.paymentOrder.fees.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a PaymentOrder
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PaymentOrder/update/' + this.paymentOrder;

	return  this.http.post(uri_, this.paymentOrder );
}

	//********************************************************************
	// loadHelper - internal helper to load a PaymentOrder
	//********************************************************************	
	loadHelper( id ) {
		this.getPaymentOrder(id)
			.subscribe((res : PaymentOrder) => {
				this.paymentOrder = res;
			});
	}
}