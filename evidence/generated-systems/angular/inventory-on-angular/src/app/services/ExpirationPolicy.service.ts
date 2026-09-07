import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ExpirationPolicy} from '../models/ExpirationPolicy';
import {StockKeepingUnitService} from '../services/StockKeepingUnit.service';
import {WarehouseService} from '../services/Warehouse.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ExpirationPolicyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	expirationPolicy : ExpirationPolicy;

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
	// add a ExpirationPolicy
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addExpirationPolicy(rejectIfDaysToExpireLessThan, autoQuarantineDaysToExpire, Sku, Warehouse, RotationMethod) : Observable<any> {
		const uri_ = this.apiUrl + '/ExpirationPolicy/create';
		const obj = {
			      		rejectIfDaysToExpireLessThan: rejectIfDaysToExpireLessThan,
      		autoQuarantineDaysToExpire: autoQuarantineDaysToExpire,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
			RotationMethod: RotationMethod
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ExpirationPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateExpirationPolicy(rejectIfDaysToExpireLessThan, autoQuarantineDaysToExpire, Sku, Warehouse, RotationMethod, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ExpirationPolicy/update/' + id;
		const obj = {
				      		rejectIfDaysToExpireLessThan: rejectIfDaysToExpireLessThan,
      		autoQuarantineDaysToExpire: autoQuarantineDaysToExpire,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		Warehouse: Warehouse != null && Warehouse.length > 0 ? Warehouse : null,
			RotationMethod: RotationMethod
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ExpirationPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteExpirationPolicy(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ExpirationPolicy/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ExpirationPolicy
	// returns the results untouched as an Observable ExpirationPolicy
	// ExpirationPolicy model
	// delegates via URI
	//********************************************************************
	getExpirationPolicy(id) : Observable<ExpirationPolicy> {
		const uri_ = this.apiUrl + '/ExpirationPolicy/load/' + id;

		return this.http.get<ExpirationPolicy>(uri_);
	}
	
	//********************************************************************
	// gets all ExpirationPolicy
	// returns the results untouched as JSON representation of an
	// Observable array of ExpirationPolicy models
	// delegates via URI
	//********************************************************************
	getExpirationPolicys() : Observable<ExpirationPolicy[]> {
		const uri_ = this.apiUrl + '/ExpirationPolicy/';

		return this
			.http.get<ExpirationPolicy[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Sku on a ExpirationPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSku( expirationPolicyId, _skuId ): Observable<any> {

		// get the ExpirationPolicy from storage
		this.loadHelper( expirationPolicyId );

	// get the StockKeepingUnit from storage
	var tmp 	= new StockKeepingUnitService(this.http).getStockKeepingUnit(_skuId);

	// assign the Sku
	this.expirationPolicy.sku = tmp;

	// save the ExpirationPolicy
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Sku on a ExpirationPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSku( expirationPolicyId ): Observable<any> {

		// get the ExpirationPolicy from storage
		this.loadHelper( expirationPolicyId );

	// assign Sku to null
	this.expirationPolicy.sku = null;

	// save the ExpirationPolicy
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Warehouse on a ExpirationPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWarehouse( expirationPolicyId, _warehouseId ): Observable<any> {

		// get the ExpirationPolicy from storage
		this.loadHelper( expirationPolicyId );

	// get the Warehouse from storage
	var tmp 	= new WarehouseService(this.http).getWarehouse(_warehouseId);

	// assign the Warehouse
	this.expirationPolicy.warehouse = tmp;

	// save the ExpirationPolicy
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Warehouse on a ExpirationPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWarehouse( expirationPolicyId ): Observable<any> {

		// get the ExpirationPolicy from storage
		this.loadHelper( expirationPolicyId );

	// assign Warehouse to null
	this.expirationPolicy.warehouse = null;

	// save the ExpirationPolicy
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a ExpirationPolicy
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ExpirationPolicy/update/' + this.expirationPolicy;

	return  this.http.post(uri_, this.expirationPolicy );
}

	//********************************************************************
	// loadHelper - internal helper to load a ExpirationPolicy
	//********************************************************************	
	loadHelper( id ) {
		this.getExpirationPolicy(id)
			.subscribe((res : ExpirationPolicy) => {
				this.expirationPolicy = res;
			});
	}
}