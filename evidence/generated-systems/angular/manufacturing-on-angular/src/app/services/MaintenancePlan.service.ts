import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {MaintenancePlan} from '../models/MaintenancePlan';
import {AssetService} from '../services/Asset.service';
import {MaintenanceOrderService} from '../services/MaintenanceOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class MaintenancePlanService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	maintenancePlan : MaintenancePlan;

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
	// add a MaintenancePlan
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addMaintenancePlan(planNumber, interval, lastServiceDate, Asset, MaintenanceOrders, Strategy) : Observable<any> {
		const uri_ = this.apiUrl + '/MaintenancePlan/create';
		const obj = {
			      		planNumber: planNumber,
      		interval: interval,
      		lastServiceDate: lastServiceDate,
      		Asset: Asset != null && Asset.length > 0 ? Asset : null,
      		MaintenanceOrders: MaintenanceOrders != null && MaintenanceOrders.length > 0 ? MaintenanceOrders : null,
			Strategy: Strategy
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a MaintenancePlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateMaintenancePlan(planNumber, interval, lastServiceDate, Asset, MaintenanceOrders, Strategy, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/MaintenancePlan/update/' + id;
		const obj = {
				      		planNumber: planNumber,
      		interval: interval,
      		lastServiceDate: lastServiceDate,
      		Asset: Asset != null && Asset.length > 0 ? Asset : null,
      		MaintenanceOrders: MaintenanceOrders != null && MaintenanceOrders.length > 0 ? MaintenanceOrders : null,
			Strategy: Strategy
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a MaintenancePlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteMaintenancePlan(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/MaintenancePlan/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a MaintenancePlan
	// returns the results untouched as an Observable MaintenancePlan
	// MaintenancePlan model
	// delegates via URI
	//********************************************************************
	getMaintenancePlan(id) : Observable<MaintenancePlan> {
		const uri_ = this.apiUrl + '/MaintenancePlan/load/' + id;

		return this.http.get<MaintenancePlan>(uri_);
	}
	
	//********************************************************************
	// gets all MaintenancePlan
	// returns the results untouched as JSON representation of an
	// Observable array of MaintenancePlan models
	// delegates via URI
	//********************************************************************
	getMaintenancePlans() : Observable<MaintenancePlan[]> {
		const uri_ = this.apiUrl + '/MaintenancePlan/';

		return this
			.http.get<MaintenancePlan[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Asset on a MaintenancePlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAsset( maintenancePlanId, _assetId ): Observable<any> {

		// get the MaintenancePlan from storage
		this.loadHelper( maintenancePlanId );

	// get the Asset from storage
	var tmp 	= new AssetService(this.http).getAsset(_assetId);

	// assign the Asset
	this.maintenancePlan.asset = tmp;

	// save the MaintenancePlan
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Asset on a MaintenancePlan
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAsset( maintenancePlanId ): Observable<any> {

		// get the MaintenancePlan from storage
		this.loadHelper( maintenancePlanId );

	// assign Asset to null
	this.maintenancePlan.asset = null;

	// save the MaintenancePlan
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more maintenanceOrdersIds as a MaintenanceOrders
	// to a MaintenancePlan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMaintenanceOrders( maintenancePlanId, maintenanceOrdersIds ): Observable<any> {

		// get the MaintenancePlan
		this.loadHelper( maintenancePlanId );

	// split on a comma with no spaces
	var idList = maintenanceOrdersIds.split(',')

	// iterate over array of maintenanceOrders ids
	idList.forEach(function (id) {
		// read the MaintenanceOrder
		var maintenanceOrder = new MaintenanceOrderService(this.http).getMaintenanceOrder(id);
		// add the MaintenanceOrder if not already assigned
		if ( this.maintenancePlan.maintenanceOrders.indexOf(maintenanceOrder) == -1 )
		this.maintenancePlan.maintenanceOrders.push(maintenanceOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more maintenanceOrdersIds as a MaintenanceOrders
	// from a MaintenancePlan
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMaintenanceOrders( maintenancePlanId, maintenanceOrdersIds ): Observable<any> {

		// get the MaintenancePlan
		this.loadHelper( maintenancePlanId );


	// split on a comma with no spaces
	var idList 					= maintenanceOrdersIds.split(',');
	var maintenanceOrders 	= this.maintenancePlan.maintenanceOrders;

	if ( maintenanceOrders != null && maintenanceOrdersIds != null ) {

		// iterate over array of maintenanceOrders ids
		maintenanceOrders.forEach(function (obj) {
			if ( maintenanceOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the MaintenanceOrder
				this.maintenancePlan.maintenanceOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a MaintenancePlan
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/MaintenancePlan/update/' + this.maintenancePlan;

	return  this.http.post(uri_, this.maintenancePlan );
}

	//********************************************************************
	// loadHelper - internal helper to load a MaintenancePlan
	//********************************************************************	
	loadHelper( id ) {
		this.getMaintenancePlan(id)
			.subscribe((res : MaintenancePlan) => {
				this.maintenancePlan = res;
			});
	}
}