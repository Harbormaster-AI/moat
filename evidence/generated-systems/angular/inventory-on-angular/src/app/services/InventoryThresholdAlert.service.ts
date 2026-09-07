import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InventoryThresholdAlert} from '../models/InventoryThresholdAlert';
import {StockKeepingUnitService} from '../services/StockKeepingUnit.service';
import {WarehouseService} from '../services/Warehouse.service';
import {StorageLocationService} from '../services/StorageLocation.service';
import {ReplenishmentPolicyService} from '../services/ReplenishmentPolicy.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InventoryThresholdAlertService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	inventoryThresholdAlert : InventoryThresholdAlert;

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
	// add a InventoryThresholdAlert
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInventoryThresholdAlert(alertNumber, detectedAt, message, Sku, Warehouse, Location, RelatedPolicy, AlertType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/InventoryThresholdAlert/create';
		const obj = {
			      		alertNumber: alertNumber,
      		detectedAt: detectedAt,
      		message: message,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
      		RelatedPolicy: RelatedPolicy != null && RelatedPolicy.length > 0 ? RelatedPolicy : null,
      		AlertType: AlertType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InventoryThresholdAlert
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInventoryThresholdAlert(alertNumber, detectedAt, message, Sku, Warehouse, Location, RelatedPolicy, AlertType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InventoryThresholdAlert/update/' + id;
		const obj = {
				      		alertNumber: alertNumber,
      		detectedAt: detectedAt,
      		message: message,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
      		RelatedPolicy: RelatedPolicy != null && RelatedPolicy.length > 0 ? RelatedPolicy : null,
      		AlertType: AlertType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InventoryThresholdAlert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInventoryThresholdAlert(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InventoryThresholdAlert/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InventoryThresholdAlert
	// returns the results untouched as an Observable InventoryThresholdAlert
	// InventoryThresholdAlert model
	// delegates via URI
	//********************************************************************
	getInventoryThresholdAlert(id) : Observable<InventoryThresholdAlert> {
		const uri_ = this.apiUrl + '/InventoryThresholdAlert/load/' + id;

		return this.http.get<InventoryThresholdAlert>(uri_);
	}
	
	//********************************************************************
	// gets all InventoryThresholdAlert
	// returns the results untouched as JSON representation of an
	// Observable array of InventoryThresholdAlert models
	// delegates via URI
	//********************************************************************
	getInventoryThresholdAlerts() : Observable<InventoryThresholdAlert[]> {
		const uri_ = this.apiUrl + '/InventoryThresholdAlert/';

		return this
			.http.get<InventoryThresholdAlert[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Sku on a InventoryThresholdAlert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSku( inventoryThresholdAlertId, _skuId ): Observable<any> {

		// get the InventoryThresholdAlert from storage
		this.loadHelper( inventoryThresholdAlertId );

	// get the StockKeepingUnit from storage
	var tmp 	= new StockKeepingUnitService(this.http).getStockKeepingUnit(_skuId);

	// assign the Sku
	this.inventoryThresholdAlert.sku = tmp;

	// save the InventoryThresholdAlert
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Sku on a InventoryThresholdAlert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSku( inventoryThresholdAlertId ): Observable<any> {

		// get the InventoryThresholdAlert from storage
		this.loadHelper( inventoryThresholdAlertId );

	// assign Sku to null
	this.inventoryThresholdAlert.sku = null;

	// save the InventoryThresholdAlert
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Warehouse on a InventoryThresholdAlert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWarehouse( inventoryThresholdAlertId, _warehouseId ): Observable<any> {

		// get the InventoryThresholdAlert from storage
		this.loadHelper( inventoryThresholdAlertId );

	// get the Warehouse from storage
	var tmp 	= new WarehouseService(this.http).getWarehouse(_warehouseId);

	// assign the Warehouse
	this.inventoryThresholdAlert.warehouse = tmp;

	// save the InventoryThresholdAlert
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Warehouse on a InventoryThresholdAlert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWarehouse( inventoryThresholdAlertId ): Observable<any> {

		// get the InventoryThresholdAlert from storage
		this.loadHelper( inventoryThresholdAlertId );

	// assign Warehouse to null
	this.inventoryThresholdAlert.warehouse = null;

	// save the InventoryThresholdAlert
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Location on a InventoryThresholdAlert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLocation( inventoryThresholdAlertId, _locationId ): Observable<any> {

		// get the InventoryThresholdAlert from storage
		this.loadHelper( inventoryThresholdAlertId );

	// get the StorageLocation from storage
	var tmp 	= new StorageLocationService(this.http).getStorageLocation(_locationId);

	// assign the Location
	this.inventoryThresholdAlert.location = tmp;

	// save the InventoryThresholdAlert
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Location on a InventoryThresholdAlert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLocation( inventoryThresholdAlertId ): Observable<any> {

		// get the InventoryThresholdAlert from storage
		this.loadHelper( inventoryThresholdAlertId );

	// assign Location to null
	this.inventoryThresholdAlert.location = null;

	// save the InventoryThresholdAlert
	return this.saveHelper();
}

		//********************************************************************
	// assigns a RelatedPolicy on a InventoryThresholdAlert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRelatedPolicy( inventoryThresholdAlertId, _relatedPolicyId ): Observable<any> {

		// get the InventoryThresholdAlert from storage
		this.loadHelper( inventoryThresholdAlertId );

	// get the ReplenishmentPolicy from storage
	var tmp 	= new ReplenishmentPolicyService(this.http).getReplenishmentPolicy(_relatedPolicyId);

	// assign the RelatedPolicy
	this.inventoryThresholdAlert.relatedPolicy = tmp;

	// save the InventoryThresholdAlert
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a RelatedPolicy on a InventoryThresholdAlert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRelatedPolicy( inventoryThresholdAlertId ): Observable<any> {

		// get the InventoryThresholdAlert from storage
		this.loadHelper( inventoryThresholdAlertId );

	// assign RelatedPolicy to null
	this.inventoryThresholdAlert.relatedPolicy = null;

	// save the InventoryThresholdAlert
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a InventoryThresholdAlert
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InventoryThresholdAlert/update/' + this.inventoryThresholdAlert;

	return  this.http.post(uri_, this.inventoryThresholdAlert );
}

	//********************************************************************
	// loadHelper - internal helper to load a InventoryThresholdAlert
	//********************************************************************	
	loadHelper( id ) {
		this.getInventoryThresholdAlert(id)
			.subscribe((res : InventoryThresholdAlert) => {
				this.inventoryThresholdAlert = res;
			});
	}
}