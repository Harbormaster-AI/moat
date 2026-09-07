import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {MedicationOrder} from '../models/MedicationOrder';
import {ClinicalOrderService} from '../services/ClinicalOrder.service';
import {PharmacyService} from '../services/Pharmacy.service';
import {MedicationDispenseService} from '../services/MedicationDispense.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class MedicationOrderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	medicationOrder : MedicationOrder;

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
	// add a MedicationOrder
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addMedicationOrder(medicationCode, dose, frequency, duration, Order, Pharmacy, Dispenses, Route) : Observable<any> {
		const uri_ = this.apiUrl + '/MedicationOrder/create';
		const obj = {
			      		medicationCode: medicationCode,
      		dose: dose,
      		frequency: frequency,
      		duration: duration,
      		Order: Order != null && Order.length > 0 ? Order : null,
      		Pharmacy: Pharmacy != null && Pharmacy.length > 0 ? Pharmacy : null,
      		Dispenses: Dispenses != null && Dispenses.length > 0 ? Dispenses : null,
			Route: Route
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a MedicationOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateMedicationOrder(medicationCode, dose, frequency, duration, Order, Pharmacy, Dispenses, Route, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/MedicationOrder/update/' + id;
		const obj = {
				      		medicationCode: medicationCode,
      		dose: dose,
      		frequency: frequency,
      		duration: duration,
      		Order: Order != null && Order.length > 0 ? Order : null,
      		Pharmacy: Pharmacy != null && Pharmacy.length > 0 ? Pharmacy : null,
      		Dispenses: Dispenses != null && Dispenses.length > 0 ? Dispenses : null,
			Route: Route
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a MedicationOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteMedicationOrder(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/MedicationOrder/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a MedicationOrder
	// returns the results untouched as an Observable MedicationOrder
	// MedicationOrder model
	// delegates via URI
	//********************************************************************
	getMedicationOrder(id) : Observable<MedicationOrder> {
		const uri_ = this.apiUrl + '/MedicationOrder/load/' + id;

		return this.http.get<MedicationOrder>(uri_);
	}
	
	//********************************************************************
	// gets all MedicationOrder
	// returns the results untouched as JSON representation of an
	// Observable array of MedicationOrder models
	// delegates via URI
	//********************************************************************
	getMedicationOrders() : Observable<MedicationOrder[]> {
		const uri_ = this.apiUrl + '/MedicationOrder/';

		return this
			.http.get<MedicationOrder[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Order on a MedicationOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrder( medicationOrderId, _orderId ): Observable<any> {

		// get the MedicationOrder from storage
		this.loadHelper( medicationOrderId );

	// get the ClinicalOrder from storage
	var tmp 	= new ClinicalOrderService(this.http).getClinicalOrder(_orderId);

	// assign the Order
	this.medicationOrder.order = tmp;

	// save the MedicationOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Order on a MedicationOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrder( medicationOrderId ): Observable<any> {

		// get the MedicationOrder from storage
		this.loadHelper( medicationOrderId );

	// assign Order to null
	this.medicationOrder.order = null;

	// save the MedicationOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Pharmacy on a MedicationOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPharmacy( medicationOrderId, _pharmacyId ): Observable<any> {

		// get the MedicationOrder from storage
		this.loadHelper( medicationOrderId );

	// get the Pharmacy from storage
	var tmp 	= new PharmacyService(this.http).getPharmacy(_pharmacyId);

	// assign the Pharmacy
	this.medicationOrder.pharmacy = tmp;

	// save the MedicationOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Pharmacy on a MedicationOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPharmacy( medicationOrderId ): Observable<any> {

		// get the MedicationOrder from storage
		this.loadHelper( medicationOrderId );

	// assign Pharmacy to null
	this.medicationOrder.pharmacy = null;

	// save the MedicationOrder
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more dispensesIds as a Dispenses
	// to a MedicationOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDispenses( medicationOrderId, dispensesIds ): Observable<any> {

		// get the MedicationOrder
		this.loadHelper( medicationOrderId );

	// split on a comma with no spaces
	var idList = dispensesIds.split(',')

	// iterate over array of dispenses ids
	idList.forEach(function (id) {
		// read the MedicationDispense
		var medicationDispense = new MedicationDispenseService(this.http).getMedicationDispense(id);
		// add the MedicationDispense if not already assigned
		if ( this.medicationOrder.dispenses.indexOf(medicationDispense) == -1 )
		this.medicationOrder.dispenses.push(medicationDispense);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dispensesIds as a Dispenses
	// from a MedicationOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDispenses( medicationOrderId, dispensesIds ): Observable<any> {

		// get the MedicationOrder
		this.loadHelper( medicationOrderId );


	// split on a comma with no spaces
	var idList 					= dispensesIds.split(',');
	var dispenses 	= this.medicationOrder.dispenses;

	if ( dispenses != null && dispensesIds != null ) {

		// iterate over array of dispenses ids
		dispenses.forEach(function (obj) {
			if ( dispensesIds.indexOf(obj._id) > -1 ) {
				// remove the MedicationDispense
				this.medicationOrder.dispenses.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a MedicationOrder
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/MedicationOrder/update/' + this.medicationOrder;

	return  this.http.post(uri_, this.medicationOrder );
}

	//********************************************************************
	// loadHelper - internal helper to load a MedicationOrder
	//********************************************************************	
	loadHelper( id ) {
		this.getMedicationOrder(id)
			.subscribe((res : MedicationOrder) => {
				this.medicationOrder = res;
			});
	}
}