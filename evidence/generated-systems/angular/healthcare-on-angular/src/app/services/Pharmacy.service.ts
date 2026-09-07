import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Pharmacy} from '../models/Pharmacy';
import {FacilityService} from '../services/Facility.service';
import {MedicationDispenseService} from '../services/MedicationDispense.service';
import {MedicationOrderService} from '../services/MedicationOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PharmacyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	pharmacy : Pharmacy;

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
	// add a Pharmacy
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPharmacy(name, Facility, MedicationDispenses, MedicationOrders) : Observable<any> {
		const uri_ = this.apiUrl + '/Pharmacy/create';
		const obj = {
			      		name: name,
      		Facility: Facility != null && Facility.length > 0 ? Facility : null,
      		MedicationDispenses: MedicationDispenses != null && MedicationDispenses.length > 0 ? MedicationDispenses : null,
			MedicationOrders: MedicationOrders != null && MedicationOrders.length > 0 ? MedicationOrders : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Pharmacy
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePharmacy(name, Facility, MedicationDispenses, MedicationOrders, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Pharmacy/update/' + id;
		const obj = {
				      		name: name,
      		Facility: Facility != null && Facility.length > 0 ? Facility : null,
      		MedicationDispenses: MedicationDispenses != null && MedicationDispenses.length > 0 ? MedicationDispenses : null,
			MedicationOrders: MedicationOrders != null && MedicationOrders.length > 0 ? MedicationOrders : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Pharmacy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePharmacy(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Pharmacy/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Pharmacy
	// returns the results untouched as an Observable Pharmacy
	// Pharmacy model
	// delegates via URI
	//********************************************************************
	getPharmacy(id) : Observable<Pharmacy> {
		const uri_ = this.apiUrl + '/Pharmacy/load/' + id;

		return this.http.get<Pharmacy>(uri_);
	}
	
	//********************************************************************
	// gets all Pharmacy
	// returns the results untouched as JSON representation of an
	// Observable array of Pharmacy models
	// delegates via URI
	//********************************************************************
	getPharmacys() : Observable<Pharmacy[]> {
		const uri_ = this.apiUrl + '/Pharmacy/';

		return this
			.http.get<Pharmacy[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Facility on a Pharmacy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignFacility( pharmacyId, _facilityId ): Observable<any> {

		// get the Pharmacy from storage
		this.loadHelper( pharmacyId );

	// get the Facility from storage
	var tmp 	= new FacilityService(this.http).getFacility(_facilityId);

	// assign the Facility
	this.pharmacy.facility = tmp;

	// save the Pharmacy
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Facility on a Pharmacy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignFacility( pharmacyId ): Observable<any> {

		// get the Pharmacy from storage
		this.loadHelper( pharmacyId );

	// assign Facility to null
	this.pharmacy.facility = null;

	// save the Pharmacy
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more medicationDispensesIds as a MedicationDispenses
	// to a Pharmacy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMedicationDispenses( pharmacyId, medicationDispensesIds ): Observable<any> {

		// get the Pharmacy
		this.loadHelper( pharmacyId );

	// split on a comma with no spaces
	var idList = medicationDispensesIds.split(',')

	// iterate over array of medicationDispenses ids
	idList.forEach(function (id) {
		// read the MedicationDispense
		var medicationDispense = new MedicationDispenseService(this.http).getMedicationDispense(id);
		// add the MedicationDispense if not already assigned
		if ( this.pharmacy.medicationDispenses.indexOf(medicationDispense) == -1 )
		this.pharmacy.medicationDispenses.push(medicationDispense);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more medicationDispensesIds as a MedicationDispenses
	// from a Pharmacy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMedicationDispenses( pharmacyId, medicationDispensesIds ): Observable<any> {

		// get the Pharmacy
		this.loadHelper( pharmacyId );


	// split on a comma with no spaces
	var idList 					= medicationDispensesIds.split(',');
	var medicationDispenses 	= this.pharmacy.medicationDispenses;

	if ( medicationDispenses != null && medicationDispensesIds != null ) {

		// iterate over array of medicationDispenses ids
		medicationDispenses.forEach(function (obj) {
			if ( medicationDispensesIds.indexOf(obj._id) > -1 ) {
				// remove the MedicationDispense
				this.pharmacy.medicationDispenses.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more medicationOrdersIds as a MedicationOrders
	// to a Pharmacy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMedicationOrders( pharmacyId, medicationOrdersIds ): Observable<any> {

		// get the Pharmacy
		this.loadHelper( pharmacyId );

	// split on a comma with no spaces
	var idList = medicationOrdersIds.split(',')

	// iterate over array of medicationOrders ids
	idList.forEach(function (id) {
		// read the MedicationOrder
		var medicationOrder = new MedicationOrderService(this.http).getMedicationOrder(id);
		// add the MedicationOrder if not already assigned
		if ( this.pharmacy.medicationOrders.indexOf(medicationOrder) == -1 )
		this.pharmacy.medicationOrders.push(medicationOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more medicationOrdersIds as a MedicationOrders
	// from a Pharmacy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMedicationOrders( pharmacyId, medicationOrdersIds ): Observable<any> {

		// get the Pharmacy
		this.loadHelper( pharmacyId );


	// split on a comma with no spaces
	var idList 					= medicationOrdersIds.split(',');
	var medicationOrders 	= this.pharmacy.medicationOrders;

	if ( medicationOrders != null && medicationOrdersIds != null ) {

		// iterate over array of medicationOrders ids
		medicationOrders.forEach(function (obj) {
			if ( medicationOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the MedicationOrder
				this.pharmacy.medicationOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Pharmacy
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Pharmacy/update/' + this.pharmacy;

	return  this.http.post(uri_, this.pharmacy );
}

	//********************************************************************
	// loadHelper - internal helper to load a Pharmacy
	//********************************************************************	
	loadHelper( id ) {
		this.getPharmacy(id)
			.subscribe((res : Pharmacy) => {
				this.pharmacy = res;
			});
	}
}