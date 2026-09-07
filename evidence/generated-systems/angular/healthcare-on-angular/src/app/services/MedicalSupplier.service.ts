import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {MedicalSupplier} from '../models/MedicalSupplier';
import {FacilityService} from '../services/Facility.service';
import {InventoryItemService} from '../services/InventoryItem.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class MedicalSupplierService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	medicalSupplier : MedicalSupplier;

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
	// add a MedicalSupplier
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addMedicalSupplier(name, website, Facilities, InventoryItems, SupplierTier) : Observable<any> {
		const uri_ = this.apiUrl + '/MedicalSupplier/create';
		const obj = {
			      		name: name,
      		website: website,
      		Facilities: Facilities != null && Facilities.length > 0 ? Facilities : null,
      		InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null,
			SupplierTier: SupplierTier
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a MedicalSupplier
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateMedicalSupplier(name, website, Facilities, InventoryItems, SupplierTier, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/MedicalSupplier/update/' + id;
		const obj = {
				      		name: name,
      		website: website,
      		Facilities: Facilities != null && Facilities.length > 0 ? Facilities : null,
      		InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null,
			SupplierTier: SupplierTier
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a MedicalSupplier
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteMedicalSupplier(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/MedicalSupplier/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a MedicalSupplier
	// returns the results untouched as an Observable MedicalSupplier
	// MedicalSupplier model
	// delegates via URI
	//********************************************************************
	getMedicalSupplier(id) : Observable<MedicalSupplier> {
		const uri_ = this.apiUrl + '/MedicalSupplier/load/' + id;

		return this.http.get<MedicalSupplier>(uri_);
	}
	
	//********************************************************************
	// gets all MedicalSupplier
	// returns the results untouched as JSON representation of an
	// Observable array of MedicalSupplier models
	// delegates via URI
	//********************************************************************
	getMedicalSuppliers() : Observable<MedicalSupplier[]> {
		const uri_ = this.apiUrl + '/MedicalSupplier/';

		return this
			.http.get<MedicalSupplier[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more facilitiesIds as a Facilities
	// to a MedicalSupplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addFacilities( medicalSupplierId, facilitiesIds ): Observable<any> {

		// get the MedicalSupplier
		this.loadHelper( medicalSupplierId );

	// split on a comma with no spaces
	var idList = facilitiesIds.split(',')

	// iterate over array of facilities ids
	idList.forEach(function (id) {
		// read the Facility
		var facility = new FacilityService(this.http).getFacility(id);
		// add the Facility if not already assigned
		if ( this.medicalSupplier.facilities.indexOf(facility) == -1 )
		this.medicalSupplier.facilities.push(facility);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more facilitiesIds as a Facilities
	// from a MedicalSupplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeFacilities( medicalSupplierId, facilitiesIds ): Observable<any> {

		// get the MedicalSupplier
		this.loadHelper( medicalSupplierId );


	// split on a comma with no spaces
	var idList 					= facilitiesIds.split(',');
	var facilities 	= this.medicalSupplier.facilities;

	if ( facilities != null && facilitiesIds != null ) {

		// iterate over array of facilities ids
		facilities.forEach(function (obj) {
			if ( facilitiesIds.indexOf(obj._id) > -1 ) {
				// remove the Facility
				this.medicalSupplier.facilities.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more inventoryItemsIds as a InventoryItems
	// to a MedicalSupplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInventoryItems( medicalSupplierId, inventoryItemsIds ): Observable<any> {

		// get the MedicalSupplier
		this.loadHelper( medicalSupplierId );

	// split on a comma with no spaces
	var idList = inventoryItemsIds.split(',')

	// iterate over array of inventoryItems ids
	idList.forEach(function (id) {
		// read the InventoryItem
		var inventoryItem = new InventoryItemService(this.http).getInventoryItem(id);
		// add the InventoryItem if not already assigned
		if ( this.medicalSupplier.inventoryItems.indexOf(inventoryItem) == -1 )
		this.medicalSupplier.inventoryItems.push(inventoryItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more inventoryItemsIds as a InventoryItems
	// from a MedicalSupplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInventoryItems( medicalSupplierId, inventoryItemsIds ): Observable<any> {

		// get the MedicalSupplier
		this.loadHelper( medicalSupplierId );


	// split on a comma with no spaces
	var idList 					= inventoryItemsIds.split(',');
	var inventoryItems 	= this.medicalSupplier.inventoryItems;

	if ( inventoryItems != null && inventoryItemsIds != null ) {

		// iterate over array of inventoryItems ids
		inventoryItems.forEach(function (obj) {
			if ( inventoryItemsIds.indexOf(obj._id) > -1 ) {
				// remove the InventoryItem
				this.medicalSupplier.inventoryItems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a MedicalSupplier
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/MedicalSupplier/update/' + this.medicalSupplier;

	return  this.http.post(uri_, this.medicalSupplier );
}

	//********************************************************************
	// loadHelper - internal helper to load a MedicalSupplier
	//********************************************************************	
	loadHelper( id ) {
		this.getMedicalSupplier(id)
			.subscribe((res : MedicalSupplier) => {
				this.medicalSupplier = res;
			});
	}
}