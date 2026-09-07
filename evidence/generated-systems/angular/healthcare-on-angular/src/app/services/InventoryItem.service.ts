import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InventoryItem} from '../models/InventoryItem';
import {FacilityService} from '../services/Facility.service';
import {MedicalSupplierService} from '../services/MedicalSupplier.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InventoryItemService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	inventoryItem : InventoryItem;

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
	// add a InventoryItem
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInventoryItem(sku, name, quantityOnHand, quantityReserved, Facility, Supplier) : Observable<any> {
		const uri_ = this.apiUrl + '/InventoryItem/create';
		const obj = {
			      		sku: sku,
      		name: name,
      		quantityOnHand: quantityOnHand,
      		quantityReserved: quantityReserved,
      		Facility: Facility != null && Facility.length > 0 ? Facility : null,
			Supplier: Supplier != null && Supplier.length > 0 ? Supplier : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInventoryItem(sku, name, quantityOnHand, quantityReserved, Facility, Supplier, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InventoryItem/update/' + id;
		const obj = {
				      		sku: sku,
      		name: name,
      		quantityOnHand: quantityOnHand,
      		quantityReserved: quantityReserved,
      		Facility: Facility != null && Facility.length > 0 ? Facility : null,
			Supplier: Supplier != null && Supplier.length > 0 ? Supplier : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInventoryItem(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InventoryItem/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InventoryItem
	// returns the results untouched as an Observable InventoryItem
	// InventoryItem model
	// delegates via URI
	//********************************************************************
	getInventoryItem(id) : Observable<InventoryItem> {
		const uri_ = this.apiUrl + '/InventoryItem/load/' + id;

		return this.http.get<InventoryItem>(uri_);
	}
	
	//********************************************************************
	// gets all InventoryItem
	// returns the results untouched as JSON representation of an
	// Observable array of InventoryItem models
	// delegates via URI
	//********************************************************************
	getInventoryItems() : Observable<InventoryItem[]> {
		const uri_ = this.apiUrl + '/InventoryItem/';

		return this
			.http.get<InventoryItem[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Facility on a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignFacility( inventoryItemId, _facilityId ): Observable<any> {

		// get the InventoryItem from storage
		this.loadHelper( inventoryItemId );

	// get the Facility from storage
	var tmp 	= new FacilityService(this.http).getFacility(_facilityId);

	// assign the Facility
	this.inventoryItem.facility = tmp;

	// save the InventoryItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Facility on a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignFacility( inventoryItemId ): Observable<any> {

		// get the InventoryItem from storage
		this.loadHelper( inventoryItemId );

	// assign Facility to null
	this.inventoryItem.facility = null;

	// save the InventoryItem
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Supplier on a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSupplier( inventoryItemId, _supplierId ): Observable<any> {

		// get the InventoryItem from storage
		this.loadHelper( inventoryItemId );

	// get the MedicalSupplier from storage
	var tmp 	= new MedicalSupplierService(this.http).getMedicalSupplier(_supplierId);

	// assign the Supplier
	this.inventoryItem.supplier = tmp;

	// save the InventoryItem
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Supplier on a InventoryItem
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSupplier( inventoryItemId ): Observable<any> {

		// get the InventoryItem from storage
		this.loadHelper( inventoryItemId );

	// assign Supplier to null
	this.inventoryItem.supplier = null;

	// save the InventoryItem
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a InventoryItem
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InventoryItem/update/' + this.inventoryItem;

	return  this.http.post(uri_, this.inventoryItem );
}

	//********************************************************************
	// loadHelper - internal helper to load a InventoryItem
	//********************************************************************	
	loadHelper( id ) {
		this.getInventoryItem(id)
			.subscribe((res : InventoryItem) => {
				this.inventoryItem = res;
			});
	}
}