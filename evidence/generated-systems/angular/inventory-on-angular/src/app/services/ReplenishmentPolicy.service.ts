import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ReplenishmentPolicy} from '../models/ReplenishmentPolicy';
import {StockKeepingUnitService} from '../services/StockKeepingUnit.service';
import {WarehouseService} from '../services/Warehouse.service';
import {StorageLocationService} from '../services/StorageLocation.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ReplenishmentPolicyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	replenishmentPolicy : ReplenishmentPolicy;

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
	// add a ReplenishmentPolicy
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addReplenishmentPolicy(minLevel, maxLevel, reorderPoint, reorderQuantity, leadTimeDays, reviewPeriodDays, Sku, Warehouse, Location, PolicyType) : Observable<any> {
		const uri_ = this.apiUrl + '/ReplenishmentPolicy/create';
		const obj = {
			      		minLevel: minLevel,
      		maxLevel: maxLevel,
      		reorderPoint: reorderPoint,
      		reorderQuantity: reorderQuantity,
      		leadTimeDays: leadTimeDays,
      		reviewPeriodDays: reviewPeriodDays,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
			PolicyType: PolicyType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ReplenishmentPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateReplenishmentPolicy(minLevel, maxLevel, reorderPoint, reorderQuantity, leadTimeDays, reviewPeriodDays, Sku, Warehouse, Location, PolicyType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ReplenishmentPolicy/update/' + id;
		const obj = {
				      		minLevel: minLevel,
      		maxLevel: maxLevel,
      		reorderPoint: reorderPoint,
      		reorderQuantity: reorderQuantity,
      		leadTimeDays: leadTimeDays,
      		reviewPeriodDays: reviewPeriodDays,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
      		Location: Location != null && Location.length > 0 ? Location : null,
			PolicyType: PolicyType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ReplenishmentPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteReplenishmentPolicy(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ReplenishmentPolicy/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ReplenishmentPolicy
	// returns the results untouched as an Observable ReplenishmentPolicy
	// ReplenishmentPolicy model
	// delegates via URI
	//********************************************************************
	getReplenishmentPolicy(id) : Observable<ReplenishmentPolicy> {
		const uri_ = this.apiUrl + '/ReplenishmentPolicy/load/' + id;

		return this.http.get<ReplenishmentPolicy>(uri_);
	}
	
	//********************************************************************
	// gets all ReplenishmentPolicy
	// returns the results untouched as JSON representation of an
	// Observable array of ReplenishmentPolicy models
	// delegates via URI
	//********************************************************************
	getReplenishmentPolicys() : Observable<ReplenishmentPolicy[]> {
		const uri_ = this.apiUrl + '/ReplenishmentPolicy/';

		return this
			.http.get<ReplenishmentPolicy[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Sku on a ReplenishmentPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSku( replenishmentPolicyId, _skuId ): Observable<any> {

		// get the ReplenishmentPolicy from storage
		this.loadHelper( replenishmentPolicyId );

	// get the StockKeepingUnit from storage
	var tmp 	= new StockKeepingUnitService(this.http).getStockKeepingUnit(_skuId);

	// assign the Sku
	this.replenishmentPolicy.sku = tmp;

	// save the ReplenishmentPolicy
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Sku on a ReplenishmentPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSku( replenishmentPolicyId ): Observable<any> {

		// get the ReplenishmentPolicy from storage
		this.loadHelper( replenishmentPolicyId );

	// assign Sku to null
	this.replenishmentPolicy.sku = null;

	// save the ReplenishmentPolicy
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Warehouse on a ReplenishmentPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWarehouse( replenishmentPolicyId, _warehouseId ): Observable<any> {

		// get the ReplenishmentPolicy from storage
		this.loadHelper( replenishmentPolicyId );

	// get the Warehouse from storage
	var tmp 	= new WarehouseService(this.http).getWarehouse(_warehouseId);

	// assign the Warehouse
	this.replenishmentPolicy.warehouse = tmp;

	// save the ReplenishmentPolicy
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Warehouse on a ReplenishmentPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWarehouse( replenishmentPolicyId ): Observable<any> {

		// get the ReplenishmentPolicy from storage
		this.loadHelper( replenishmentPolicyId );

	// assign Warehouse to null
	this.replenishmentPolicy.warehouse = null;

	// save the ReplenishmentPolicy
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Location on a ReplenishmentPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLocation( replenishmentPolicyId, _locationId ): Observable<any> {

		// get the ReplenishmentPolicy from storage
		this.loadHelper( replenishmentPolicyId );

	// get the StorageLocation from storage
	var tmp 	= new StorageLocationService(this.http).getStorageLocation(_locationId);

	// assign the Location
	this.replenishmentPolicy.location = tmp;

	// save the ReplenishmentPolicy
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Location on a ReplenishmentPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLocation( replenishmentPolicyId ): Observable<any> {

		// get the ReplenishmentPolicy from storage
		this.loadHelper( replenishmentPolicyId );

	// assign Location to null
	this.replenishmentPolicy.location = null;

	// save the ReplenishmentPolicy
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a ReplenishmentPolicy
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ReplenishmentPolicy/update/' + this.replenishmentPolicy;

	return  this.http.post(uri_, this.replenishmentPolicy );
}

	//********************************************************************
	// loadHelper - internal helper to load a ReplenishmentPolicy
	//********************************************************************	
	loadHelper( id ) {
		this.getReplenishmentPolicy(id)
			.subscribe((res : ReplenishmentPolicy) => {
				this.replenishmentPolicy = res;
			});
	}
}