import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Item} from '../models/Item';
import {BusinessUnitService} from '../services/BusinessUnit.service';
import {BOMService} from '../services/BOM.service';
import {RoutingService} from '../services/Routing.service';
import {SupplierService} from '../services/Supplier.service';
import {QualitySpecificationService} from '../services/QualitySpecification.service';
import {InventoryItemService} from '../services/InventoryItem.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ItemService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	item : Item;

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
	// add a Item
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addItem(itemNumber, name, standardCost, weight, asSerialControlled, BusinessUnit, Boms, Routings, Suppliers, QualitySpecifications, InventoryItems, ItemType, ProcurementType, UnitOfMeasure, LifecycleStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/Item/create';
		const obj = {
			      		itemNumber: itemNumber,
      		name: name,
      		standardCost: standardCost,
      		weight: weight,
      		asSerialControlled: asSerialControlled,
      		BusinessUnit: BusinessUnit != null && BusinessUnit.length > 0 ? BusinessUnit : null,
      		Boms: Boms != null && Boms.length > 0 ? Boms : null,
      		Routings: Routings != null && Routings.length > 0 ? Routings : null,
      		Suppliers: Suppliers != null && Suppliers.length > 0 ? Suppliers : null,
      		QualitySpecifications: QualitySpecifications != null && QualitySpecifications.length > 0 ? QualitySpecifications : null,
      		InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null,
      		ItemType: ItemType,
      		ProcurementType: ProcurementType,
      		UnitOfMeasure: UnitOfMeasure,
			LifecycleStatus: LifecycleStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Item
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateItem(itemNumber, name, standardCost, weight, asSerialControlled, BusinessUnit, Boms, Routings, Suppliers, QualitySpecifications, InventoryItems, ItemType, ProcurementType, UnitOfMeasure, LifecycleStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Item/update/' + id;
		const obj = {
				      		itemNumber: itemNumber,
      		name: name,
      		standardCost: standardCost,
      		weight: weight,
      		asSerialControlled: asSerialControlled,
      		BusinessUnit: BusinessUnit != null && BusinessUnit.length > 0 ? BusinessUnit : null,
      		Boms: Boms != null && Boms.length > 0 ? Boms : null,
      		Routings: Routings != null && Routings.length > 0 ? Routings : null,
      		Suppliers: Suppliers != null && Suppliers.length > 0 ? Suppliers : null,
      		QualitySpecifications: QualitySpecifications != null && QualitySpecifications.length > 0 ? QualitySpecifications : null,
      		InventoryItems: InventoryItems != null && InventoryItems.length > 0 ? InventoryItems : null,
      		ItemType: ItemType,
      		ProcurementType: ProcurementType,
      		UnitOfMeasure: UnitOfMeasure,
			LifecycleStatus: LifecycleStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Item
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteItem(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Item/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Item
	// returns the results untouched as an Observable Item
	// Item model
	// delegates via URI
	//********************************************************************
	getItem(id) : Observable<Item> {
		const uri_ = this.apiUrl + '/Item/load/' + id;

		return this.http.get<Item>(uri_);
	}
	
	//********************************************************************
	// gets all Item
	// returns the results untouched as JSON representation of an
	// Observable array of Item models
	// delegates via URI
	//********************************************************************
	getItems() : Observable<Item[]> {
		const uri_ = this.apiUrl + '/Item/';

		return this
			.http.get<Item[]>(uri_);
	}
	
			//********************************************************************
	// assigns a BusinessUnit on a Item
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignBusinessUnit( itemId, _businessUnitId ): Observable<any> {

		// get the Item from storage
		this.loadHelper( itemId );

	// get the BusinessUnit from storage
	var tmp 	= new BusinessUnitService(this.http).getBusinessUnit(_businessUnitId);

	// assign the BusinessUnit
	this.item.businessUnit = tmp;

	// save the Item
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a BusinessUnit on a Item
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignBusinessUnit( itemId ): Observable<any> {

		// get the Item from storage
		this.loadHelper( itemId );

	// assign BusinessUnit to null
	this.item.businessUnit = null;

	// save the Item
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more bomsIds as a Boms
	// to a Item
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addBoms( itemId, bomsIds ): Observable<any> {

		// get the Item
		this.loadHelper( itemId );

	// split on a comma with no spaces
	var idList = bomsIds.split(',')

	// iterate over array of boms ids
	idList.forEach(function (id) {
		// read the BOM
		var bOM = new BOMService(this.http).getBOM(id);
		// add the BOM if not already assigned
		if ( this.item.boms.indexOf(bOM) == -1 )
		this.item.boms.push(bOM);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more bomsIds as a Boms
	// from a Item
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeBoms( itemId, bomsIds ): Observable<any> {

		// get the Item
		this.loadHelper( itemId );


	// split on a comma with no spaces
	var idList 					= bomsIds.split(',');
	var boms 	= this.item.boms;

	if ( boms != null && bomsIds != null ) {

		// iterate over array of boms ids
		boms.forEach(function (obj) {
			if ( bomsIds.indexOf(obj._id) > -1 ) {
				// remove the BOM
				this.item.boms.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more routingsIds as a Routings
	// to a Item
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRoutings( itemId, routingsIds ): Observable<any> {

		// get the Item
		this.loadHelper( itemId );

	// split on a comma with no spaces
	var idList = routingsIds.split(',')

	// iterate over array of routings ids
	idList.forEach(function (id) {
		// read the Routing
		var routing = new RoutingService(this.http).getRouting(id);
		// add the Routing if not already assigned
		if ( this.item.routings.indexOf(routing) == -1 )
		this.item.routings.push(routing);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more routingsIds as a Routings
	// from a Item
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRoutings( itemId, routingsIds ): Observable<any> {

		// get the Item
		this.loadHelper( itemId );


	// split on a comma with no spaces
	var idList 					= routingsIds.split(',');
	var routings 	= this.item.routings;

	if ( routings != null && routingsIds != null ) {

		// iterate over array of routings ids
		routings.forEach(function (obj) {
			if ( routingsIds.indexOf(obj._id) > -1 ) {
				// remove the Routing
				this.item.routings.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more suppliersIds as a Suppliers
	// to a Item
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSuppliers( itemId, suppliersIds ): Observable<any> {

		// get the Item
		this.loadHelper( itemId );

	// split on a comma with no spaces
	var idList = suppliersIds.split(',')

	// iterate over array of suppliers ids
	idList.forEach(function (id) {
		// read the Supplier
		var supplier = new SupplierService(this.http).getSupplier(id);
		// add the Supplier if not already assigned
		if ( this.item.suppliers.indexOf(supplier) == -1 )
		this.item.suppliers.push(supplier);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more suppliersIds as a Suppliers
	// from a Item
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSuppliers( itemId, suppliersIds ): Observable<any> {

		// get the Item
		this.loadHelper( itemId );


	// split on a comma with no spaces
	var idList 					= suppliersIds.split(',');
	var suppliers 	= this.item.suppliers;

	if ( suppliers != null && suppliersIds != null ) {

		// iterate over array of suppliers ids
		suppliers.forEach(function (obj) {
			if ( suppliersIds.indexOf(obj._id) > -1 ) {
				// remove the Supplier
				this.item.suppliers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more qualitySpecificationsIds as a QualitySpecifications
	// to a Item
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addQualitySpecifications( itemId, qualitySpecificationsIds ): Observable<any> {

		// get the Item
		this.loadHelper( itemId );

	// split on a comma with no spaces
	var idList = qualitySpecificationsIds.split(',')

	// iterate over array of qualitySpecifications ids
	idList.forEach(function (id) {
		// read the QualitySpecification
		var qualitySpecification = new QualitySpecificationService(this.http).getQualitySpecification(id);
		// add the QualitySpecification if not already assigned
		if ( this.item.qualitySpecifications.indexOf(qualitySpecification) == -1 )
		this.item.qualitySpecifications.push(qualitySpecification);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more qualitySpecificationsIds as a QualitySpecifications
	// from a Item
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeQualitySpecifications( itemId, qualitySpecificationsIds ): Observable<any> {

		// get the Item
		this.loadHelper( itemId );


	// split on a comma with no spaces
	var idList 					= qualitySpecificationsIds.split(',');
	var qualitySpecifications 	= this.item.qualitySpecifications;

	if ( qualitySpecifications != null && qualitySpecificationsIds != null ) {

		// iterate over array of qualitySpecifications ids
		qualitySpecifications.forEach(function (obj) {
			if ( qualitySpecificationsIds.indexOf(obj._id) > -1 ) {
				// remove the QualitySpecification
				this.item.qualitySpecifications.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more inventoryItemsIds as a InventoryItems
	// to a Item
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInventoryItems( itemId, inventoryItemsIds ): Observable<any> {

		// get the Item
		this.loadHelper( itemId );

	// split on a comma with no spaces
	var idList = inventoryItemsIds.split(',')

	// iterate over array of inventoryItems ids
	idList.forEach(function (id) {
		// read the InventoryItem
		var inventoryItem = new InventoryItemService(this.http).getInventoryItem(id);
		// add the InventoryItem if not already assigned
		if ( this.item.inventoryItems.indexOf(inventoryItem) == -1 )
		this.item.inventoryItems.push(inventoryItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more inventoryItemsIds as a InventoryItems
	// from a Item
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInventoryItems( itemId, inventoryItemsIds ): Observable<any> {

		// get the Item
		this.loadHelper( itemId );


	// split on a comma with no spaces
	var idList 					= inventoryItemsIds.split(',');
	var inventoryItems 	= this.item.inventoryItems;

	if ( inventoryItems != null && inventoryItemsIds != null ) {

		// iterate over array of inventoryItems ids
		inventoryItems.forEach(function (obj) {
			if ( inventoryItemsIds.indexOf(obj._id) > -1 ) {
				// remove the InventoryItem
				this.item.inventoryItems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Item
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Item/update/' + this.item;

	return  this.http.post(uri_, this.item );
}

	//********************************************************************
	// loadHelper - internal helper to load a Item
	//********************************************************************	
	loadHelper( id ) {
		this.getItem(id)
			.subscribe((res : Item) => {
				this.item = res;
			});
	}
}