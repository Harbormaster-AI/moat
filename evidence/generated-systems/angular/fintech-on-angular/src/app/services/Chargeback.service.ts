import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Chargeback} from '../models/Chargeback';
import {DisputeService} from '../services/Dispute.service';
import {TransactionService} from '../services/Transaction.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ChargebackService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	chargeback : Chargeback;

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
	// add a Chargeback
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addChargeback(chargebackReference, amount, postedAt, Dispute, Transaction, Stage, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Chargeback/create';
		const obj = {
			      		chargebackReference: chargebackReference,
      		amount: amount,
      		postedAt: postedAt,
      		Dispute: Dispute != null && Dispute.length > 0 ? Dispute : null,
      		Transaction: Transaction != null && Transaction.length > 0 ? Transaction : null,
      		Stage: Stage,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Chargeback
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateChargeback(chargebackReference, amount, postedAt, Dispute, Transaction, Stage, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Chargeback/update/' + id;
		const obj = {
				      		chargebackReference: chargebackReference,
      		amount: amount,
      		postedAt: postedAt,
      		Dispute: Dispute != null && Dispute.length > 0 ? Dispute : null,
      		Transaction: Transaction != null && Transaction.length > 0 ? Transaction : null,
      		Stage: Stage,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Chargeback
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteChargeback(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Chargeback/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Chargeback
	// returns the results untouched as an Observable Chargeback
	// Chargeback model
	// delegates via URI
	//********************************************************************
	getChargeback(id) : Observable<Chargeback> {
		const uri_ = this.apiUrl + '/Chargeback/load/' + id;

		return this.http.get<Chargeback>(uri_);
	}
	
	//********************************************************************
	// gets all Chargeback
	// returns the results untouched as JSON representation of an
	// Observable array of Chargeback models
	// delegates via URI
	//********************************************************************
	getChargebacks() : Observable<Chargeback[]> {
		const uri_ = this.apiUrl + '/Chargeback/';

		return this
			.http.get<Chargeback[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Dispute on a Chargeback
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDispute( chargebackId, _disputeId ): Observable<any> {

		// get the Chargeback from storage
		this.loadHelper( chargebackId );

	// get the Dispute from storage
	var tmp 	= new DisputeService(this.http).getDispute(_disputeId);

	// assign the Dispute
	this.chargeback.dispute = tmp;

	// save the Chargeback
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Dispute on a Chargeback
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDispute( chargebackId ): Observable<any> {

		// get the Chargeback from storage
		this.loadHelper( chargebackId );

	// assign Dispute to null
	this.chargeback.dispute = null;

	// save the Chargeback
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Transaction on a Chargeback
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTransaction( chargebackId, _transactionId ): Observable<any> {

		// get the Chargeback from storage
		this.loadHelper( chargebackId );

	// get the Transaction from storage
	var tmp 	= new TransactionService(this.http).getTransaction(_transactionId);

	// assign the Transaction
	this.chargeback.transaction = tmp;

	// save the Chargeback
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Transaction on a Chargeback
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTransaction( chargebackId ): Observable<any> {

		// get the Chargeback from storage
		this.loadHelper( chargebackId );

	// assign Transaction to null
	this.chargeback.transaction = null;

	// save the Chargeback
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Chargeback
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Chargeback/update/' + this.chargeback;

	return  this.http.post(uri_, this.chargeback );
}

	//********************************************************************
	// loadHelper - internal helper to load a Chargeback
	//********************************************************************	
	loadHelper( id ) {
		this.getChargeback(id)
			.subscribe((res : Chargeback) => {
				this.chargeback = res;
			});
	}
}