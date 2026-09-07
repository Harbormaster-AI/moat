import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AircraftOrder} from '../models/AircraftOrder';
import {OperatorService} from '../services/Operator.service';
import {AircraftVariantService} from '../services/AircraftVariant.service';
import {QuoteService} from '../services/Quote.service';
import {PurchaseAgreementService} from '../services/PurchaseAgreement.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AircraftOrderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	aircraftOrder : AircraftOrder;

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
	// add a AircraftOrder
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAircraftOrder(orderNumber, totalAmount, Operator, Variant, Quote, PurchaseAgreement, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/AircraftOrder/create';
		const obj = {
			      		orderNumber: orderNumber,
      		totalAmount: totalAmount,
      		Operator: Operator != null && Operator.length > 0 ? Operator : null,
      		Variant: Variant != null && Variant.length > 0 ? Variant : null,
      		Quote: Quote != null && Quote.length > 0 ? Quote : null,
      		PurchaseAgreement: PurchaseAgreement != null && PurchaseAgreement.length > 0 ? PurchaseAgreement : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AircraftOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAircraftOrder(orderNumber, totalAmount, Operator, Variant, Quote, PurchaseAgreement, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AircraftOrder/update/' + id;
		const obj = {
				      		orderNumber: orderNumber,
      		totalAmount: totalAmount,
      		Operator: Operator != null && Operator.length > 0 ? Operator : null,
      		Variant: Variant != null && Variant.length > 0 ? Variant : null,
      		Quote: Quote != null && Quote.length > 0 ? Quote : null,
      		PurchaseAgreement: PurchaseAgreement != null && PurchaseAgreement.length > 0 ? PurchaseAgreement : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AircraftOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAircraftOrder(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AircraftOrder/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AircraftOrder
	// returns the results untouched as an Observable AircraftOrder
	// AircraftOrder model
	// delegates via URI
	//********************************************************************
	getAircraftOrder(id) : Observable<AircraftOrder> {
		const uri_ = this.apiUrl + '/AircraftOrder/load/' + id;

		return this.http.get<AircraftOrder>(uri_);
	}
	
	//********************************************************************
	// gets all AircraftOrder
	// returns the results untouched as JSON representation of an
	// Observable array of AircraftOrder models
	// delegates via URI
	//********************************************************************
	getAircraftOrders() : Observable<AircraftOrder[]> {
		const uri_ = this.apiUrl + '/AircraftOrder/';

		return this
			.http.get<AircraftOrder[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Operator on a AircraftOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOperator( aircraftOrderId, _operatorId ): Observable<any> {

		// get the AircraftOrder from storage
		this.loadHelper( aircraftOrderId );

	// get the Operator from storage
	var tmp 	= new OperatorService(this.http).getOperator(_operatorId);

	// assign the Operator
	this.aircraftOrder.operator = tmp;

	// save the AircraftOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Operator on a AircraftOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOperator( aircraftOrderId ): Observable<any> {

		// get the AircraftOrder from storage
		this.loadHelper( aircraftOrderId );

	// assign Operator to null
	this.aircraftOrder.operator = null;

	// save the AircraftOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Variant on a AircraftOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignVariant( aircraftOrderId, _variantId ): Observable<any> {

		// get the AircraftOrder from storage
		this.loadHelper( aircraftOrderId );

	// get the AircraftVariant from storage
	var tmp 	= new AircraftVariantService(this.http).getAircraftVariant(_variantId);

	// assign the Variant
	this.aircraftOrder.variant = tmp;

	// save the AircraftOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Variant on a AircraftOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignVariant( aircraftOrderId ): Observable<any> {

		// get the AircraftOrder from storage
		this.loadHelper( aircraftOrderId );

	// assign Variant to null
	this.aircraftOrder.variant = null;

	// save the AircraftOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Quote on a AircraftOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignQuote( aircraftOrderId, _quoteId ): Observable<any> {

		// get the AircraftOrder from storage
		this.loadHelper( aircraftOrderId );

	// get the Quote from storage
	var tmp 	= new QuoteService(this.http).getQuote(_quoteId);

	// assign the Quote
	this.aircraftOrder.quote = tmp;

	// save the AircraftOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Quote on a AircraftOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignQuote( aircraftOrderId ): Observable<any> {

		// get the AircraftOrder from storage
		this.loadHelper( aircraftOrderId );

	// assign Quote to null
	this.aircraftOrder.quote = null;

	// save the AircraftOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a PurchaseAgreement on a AircraftOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPurchaseAgreement( aircraftOrderId, _purchaseAgreementId ): Observable<any> {

		// get the AircraftOrder from storage
		this.loadHelper( aircraftOrderId );

	// get the PurchaseAgreement from storage
	var tmp 	= new PurchaseAgreementService(this.http).getPurchaseAgreement(_purchaseAgreementId);

	// assign the PurchaseAgreement
	this.aircraftOrder.purchaseAgreement = tmp;

	// save the AircraftOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PurchaseAgreement on a AircraftOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPurchaseAgreement( aircraftOrderId ): Observable<any> {

		// get the AircraftOrder from storage
		this.loadHelper( aircraftOrderId );

	// assign PurchaseAgreement to null
	this.aircraftOrder.purchaseAgreement = null;

	// save the AircraftOrder
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a AircraftOrder
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AircraftOrder/update/' + this.aircraftOrder;

	return  this.http.post(uri_, this.aircraftOrder );
}

	//********************************************************************
	// loadHelper - internal helper to load a AircraftOrder
	//********************************************************************	
	loadHelper( id ) {
		this.getAircraftOrder(id)
			.subscribe((res : AircraftOrder) => {
				this.aircraftOrder = res;
			});
	}
}