import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Supplier} from '../models/Supplier';
import {EnterpriseService} from '../services/Enterprise.service';
import {ItemService} from '../services/Item.service';
import {PurchaseOrderService} from '../services/PurchaseOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class SupplierService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	supplier : Supplier;

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
	// add a Supplier
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addSupplier(name, supplierCode, address, Enterprises, Items, PurchaseOrders, SupplierTier, PaymentTerms) : Observable<any> {
		const uri_ = this.apiUrl + '/Supplier/create';
		const obj = {
			      		name: name,
      		supplierCode: supplierCode,
      		address: address,
      		Enterprises: Enterprises != null && Enterprises.length > 0 ? Enterprises : null,
      		Items: Items != null && Items.length > 0 ? Items : null,
      		PurchaseOrders: PurchaseOrders != null && PurchaseOrders.length > 0 ? PurchaseOrders : null,
      		SupplierTier: SupplierTier,
			PaymentTerms: PaymentTerms
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Supplier
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateSupplier(name, supplierCode, address, Enterprises, Items, PurchaseOrders, SupplierTier, PaymentTerms, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Supplier/update/' + id;
		const obj = {
				      		name: name,
      		supplierCode: supplierCode,
      		address: address,
      		Enterprises: Enterprises != null && Enterprises.length > 0 ? Enterprises : null,
      		Items: Items != null && Items.length > 0 ? Items : null,
      		PurchaseOrders: PurchaseOrders != null && PurchaseOrders.length > 0 ? PurchaseOrders : null,
      		SupplierTier: SupplierTier,
			PaymentTerms: PaymentTerms
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Supplier
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteSupplier(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Supplier/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Supplier
	// returns the results untouched as an Observable Supplier
	// Supplier model
	// delegates via URI
	//********************************************************************
	getSupplier(id) : Observable<Supplier> {
		const uri_ = this.apiUrl + '/Supplier/load/' + id;

		return this.http.get<Supplier>(uri_);
	}
	
	//********************************************************************
	// gets all Supplier
	// returns the results untouched as JSON representation of an
	// Observable array of Supplier models
	// delegates via URI
	//********************************************************************
	getSuppliers() : Observable<Supplier[]> {
		const uri_ = this.apiUrl + '/Supplier/';

		return this
			.http.get<Supplier[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more enterprisesIds as a Enterprises
	// to a Supplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEnterprises( supplierId, enterprisesIds ): Observable<any> {

		// get the Supplier
		this.loadHelper( supplierId );

	// split on a comma with no spaces
	var idList = enterprisesIds.split(',')

	// iterate over array of enterprises ids
	idList.forEach(function (id) {
		// read the Enterprise
		var enterprise = new EnterpriseService(this.http).getEnterprise(id);
		// add the Enterprise if not already assigned
		if ( this.supplier.enterprises.indexOf(enterprise) == -1 )
		this.supplier.enterprises.push(enterprise);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more enterprisesIds as a Enterprises
	// from a Supplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEnterprises( supplierId, enterprisesIds ): Observable<any> {

		// get the Supplier
		this.loadHelper( supplierId );


	// split on a comma with no spaces
	var idList 					= enterprisesIds.split(',');
	var enterprises 	= this.supplier.enterprises;

	if ( enterprises != null && enterprisesIds != null ) {

		// iterate over array of enterprises ids
		enterprises.forEach(function (obj) {
			if ( enterprisesIds.indexOf(obj._id) > -1 ) {
				// remove the Enterprise
				this.supplier.enterprises.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more itemsIds as a Items
	// to a Supplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addItems( supplierId, itemsIds ): Observable<any> {

		// get the Supplier
		this.loadHelper( supplierId );

	// split on a comma with no spaces
	var idList = itemsIds.split(',')

	// iterate over array of items ids
	idList.forEach(function (id) {
		// read the Item
		var item = new ItemService(this.http).getItem(id);
		// add the Item if not already assigned
		if ( this.supplier.items.indexOf(item) == -1 )
		this.supplier.items.push(item);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more itemsIds as a Items
	// from a Supplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeItems( supplierId, itemsIds ): Observable<any> {

		// get the Supplier
		this.loadHelper( supplierId );


	// split on a comma with no spaces
	var idList 					= itemsIds.split(',');
	var items 	= this.supplier.items;

	if ( items != null && itemsIds != null ) {

		// iterate over array of items ids
		items.forEach(function (obj) {
			if ( itemsIds.indexOf(obj._id) > -1 ) {
				// remove the Item
				this.supplier.items.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more purchaseOrdersIds as a PurchaseOrders
	// to a Supplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPurchaseOrders( supplierId, purchaseOrdersIds ): Observable<any> {

		// get the Supplier
		this.loadHelper( supplierId );

	// split on a comma with no spaces
	var idList = purchaseOrdersIds.split(',')

	// iterate over array of purchaseOrders ids
	idList.forEach(function (id) {
		// read the PurchaseOrder
		var purchaseOrder = new PurchaseOrderService(this.http).getPurchaseOrder(id);
		// add the PurchaseOrder if not already assigned
		if ( this.supplier.purchaseOrders.indexOf(purchaseOrder) == -1 )
		this.supplier.purchaseOrders.push(purchaseOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more purchaseOrdersIds as a PurchaseOrders
	// from a Supplier
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePurchaseOrders( supplierId, purchaseOrdersIds ): Observable<any> {

		// get the Supplier
		this.loadHelper( supplierId );


	// split on a comma with no spaces
	var idList 					= purchaseOrdersIds.split(',');
	var purchaseOrders 	= this.supplier.purchaseOrders;

	if ( purchaseOrders != null && purchaseOrdersIds != null ) {

		// iterate over array of purchaseOrders ids
		purchaseOrders.forEach(function (obj) {
			if ( purchaseOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the PurchaseOrder
				this.supplier.purchaseOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Supplier
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Supplier/update/' + this.supplier;

	return  this.http.post(uri_, this.supplier );
}

	//********************************************************************
	// loadHelper - internal helper to load a Supplier
	//********************************************************************	
	loadHelper( id ) {
		this.getSupplier(id)
			.subscribe((res : Supplier) => {
				this.supplier = res;
			});
	}
}