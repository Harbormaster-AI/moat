import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PurchaseAgreement} from '../models/PurchaseAgreement';
import {AircraftOrderService} from '../services/AircraftOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PurchaseAgreementService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	purchaseAgreement : PurchaseAgreement;

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
	// add a PurchaseAgreement
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPurchaseAgreement(agreementNumber, effectiveDate, AircraftOrder) : Observable<any> {
		const uri_ = this.apiUrl + '/PurchaseAgreement/create';
		const obj = {
			      		agreementNumber: agreementNumber,
      		effectiveDate: effectiveDate,
			AircraftOrder: AircraftOrder != null && AircraftOrder.length > 0 ? AircraftOrder : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PurchaseAgreement
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePurchaseAgreement(agreementNumber, effectiveDate, AircraftOrder, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PurchaseAgreement/update/' + id;
		const obj = {
				      		agreementNumber: agreementNumber,
      		effectiveDate: effectiveDate,
			AircraftOrder: AircraftOrder != null && AircraftOrder.length > 0 ? AircraftOrder : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PurchaseAgreement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePurchaseAgreement(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PurchaseAgreement/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PurchaseAgreement
	// returns the results untouched as an Observable PurchaseAgreement
	// PurchaseAgreement model
	// delegates via URI
	//********************************************************************
	getPurchaseAgreement(id) : Observable<PurchaseAgreement> {
		const uri_ = this.apiUrl + '/PurchaseAgreement/load/' + id;

		return this.http.get<PurchaseAgreement>(uri_);
	}
	
	//********************************************************************
	// gets all PurchaseAgreement
	// returns the results untouched as JSON representation of an
	// Observable array of PurchaseAgreement models
	// delegates via URI
	//********************************************************************
	getPurchaseAgreements() : Observable<PurchaseAgreement[]> {
		const uri_ = this.apiUrl + '/PurchaseAgreement/';

		return this
			.http.get<PurchaseAgreement[]>(uri_);
	}
	
			//********************************************************************
	// assigns a AircraftOrder on a PurchaseAgreement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAircraftOrder( purchaseAgreementId, _aircraftOrderId ): Observable<any> {

		// get the PurchaseAgreement from storage
		this.loadHelper( purchaseAgreementId );

	// get the AircraftOrder from storage
	var tmp 	= new AircraftOrderService(this.http).getAircraftOrder(_aircraftOrderId);

	// assign the AircraftOrder
	this.purchaseAgreement.aircraftOrder = tmp;

	// save the PurchaseAgreement
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a AircraftOrder on a PurchaseAgreement
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAircraftOrder( purchaseAgreementId ): Observable<any> {

		// get the PurchaseAgreement from storage
		this.loadHelper( purchaseAgreementId );

	// assign AircraftOrder to null
	this.purchaseAgreement.aircraftOrder = null;

	// save the PurchaseAgreement
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a PurchaseAgreement
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PurchaseAgreement/update/' + this.purchaseAgreement;

	return  this.http.post(uri_, this.purchaseAgreement );
}

	//********************************************************************
	// loadHelper - internal helper to load a PurchaseAgreement
	//********************************************************************	
	loadHelper( id ) {
		this.getPurchaseAgreement(id)
			.subscribe((res : PurchaseAgreement) => {
				this.purchaseAgreement = res;
			});
	}
}