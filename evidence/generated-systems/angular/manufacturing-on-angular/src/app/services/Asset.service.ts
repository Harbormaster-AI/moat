import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Asset} from '../models/Asset';
import {PlantService} from '../services/Plant.service';
import {WorkCenterService} from '../services/WorkCenter.service';
import {MaintenanceOrderService} from '../services/MaintenanceOrder.service';
import {MaintenancePlanService} from '../services/MaintenancePlan.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AssetService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	asset : Asset;

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
	// add a Asset
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAsset(assetTag, assetName, commissioningDate, Plant, WorkCenter, MaintenanceOrders, MaintenancePlans, AssetStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/Asset/create';
		const obj = {
			      		assetTag: assetTag,
      		assetName: assetName,
      		commissioningDate: commissioningDate,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		WorkCenter: WorkCenter != null && WorkCenter.length > 0 ? WorkCenter : null,
      		MaintenanceOrders: MaintenanceOrders != null && MaintenanceOrders.length > 0 ? MaintenanceOrders : null,
      		MaintenancePlans: MaintenancePlans != null && MaintenancePlans.length > 0 ? MaintenancePlans : null,
			AssetStatus: AssetStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Asset
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAsset(assetTag, assetName, commissioningDate, Plant, WorkCenter, MaintenanceOrders, MaintenancePlans, AssetStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Asset/update/' + id;
		const obj = {
				      		assetTag: assetTag,
      		assetName: assetName,
      		commissioningDate: commissioningDate,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		WorkCenter: WorkCenter != null && WorkCenter.length > 0 ? WorkCenter : null,
      		MaintenanceOrders: MaintenanceOrders != null && MaintenanceOrders.length > 0 ? MaintenanceOrders : null,
      		MaintenancePlans: MaintenancePlans != null && MaintenancePlans.length > 0 ? MaintenancePlans : null,
			AssetStatus: AssetStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Asset
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAsset(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Asset/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Asset
	// returns the results untouched as an Observable Asset
	// Asset model
	// delegates via URI
	//********************************************************************
	getAsset(id) : Observable<Asset> {
		const uri_ = this.apiUrl + '/Asset/load/' + id;

		return this.http.get<Asset>(uri_);
	}
	
	//********************************************************************
	// gets all Asset
	// returns the results untouched as JSON representation of an
	// Observable array of Asset models
	// delegates via URI
	//********************************************************************
	getAssets() : Observable<Asset[]> {
		const uri_ = this.apiUrl + '/Asset/';

		return this
			.http.get<Asset[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Plant on a Asset
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPlant( assetId, _plantId ): Observable<any> {

		// get the Asset from storage
		this.loadHelper( assetId );

	// get the Plant from storage
	var tmp 	= new PlantService(this.http).getPlant(_plantId);

	// assign the Plant
	this.asset.plant = tmp;

	// save the Asset
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Plant on a Asset
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPlant( assetId ): Observable<any> {

		// get the Asset from storage
		this.loadHelper( assetId );

	// assign Plant to null
	this.asset.plant = null;

	// save the Asset
	return this.saveHelper();
}

		//********************************************************************
	// assigns a WorkCenter on a Asset
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkCenter( assetId, _workCenterId ): Observable<any> {

		// get the Asset from storage
		this.loadHelper( assetId );

	// get the WorkCenter from storage
	var tmp 	= new WorkCenterService(this.http).getWorkCenter(_workCenterId);

	// assign the WorkCenter
	this.asset.workCenter = tmp;

	// save the Asset
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a WorkCenter on a Asset
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkCenter( assetId ): Observable<any> {

		// get the Asset from storage
		this.loadHelper( assetId );

	// assign WorkCenter to null
	this.asset.workCenter = null;

	// save the Asset
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more maintenanceOrdersIds as a MaintenanceOrders
	// to a Asset
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMaintenanceOrders( assetId, maintenanceOrdersIds ): Observable<any> {

		// get the Asset
		this.loadHelper( assetId );

	// split on a comma with no spaces
	var idList = maintenanceOrdersIds.split(',')

	// iterate over array of maintenanceOrders ids
	idList.forEach(function (id) {
		// read the MaintenanceOrder
		var maintenanceOrder = new MaintenanceOrderService(this.http).getMaintenanceOrder(id);
		// add the MaintenanceOrder if not already assigned
		if ( this.asset.maintenanceOrders.indexOf(maintenanceOrder) == -1 )
		this.asset.maintenanceOrders.push(maintenanceOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more maintenanceOrdersIds as a MaintenanceOrders
	// from a Asset
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMaintenanceOrders( assetId, maintenanceOrdersIds ): Observable<any> {

		// get the Asset
		this.loadHelper( assetId );


	// split on a comma with no spaces
	var idList 					= maintenanceOrdersIds.split(',');
	var maintenanceOrders 	= this.asset.maintenanceOrders;

	if ( maintenanceOrders != null && maintenanceOrdersIds != null ) {

		// iterate over array of maintenanceOrders ids
		maintenanceOrders.forEach(function (obj) {
			if ( maintenanceOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the MaintenanceOrder
				this.asset.maintenanceOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more maintenancePlansIds as a MaintenancePlans
	// to a Asset
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMaintenancePlans( assetId, maintenancePlansIds ): Observable<any> {

		// get the Asset
		this.loadHelper( assetId );

	// split on a comma with no spaces
	var idList = maintenancePlansIds.split(',')

	// iterate over array of maintenancePlans ids
	idList.forEach(function (id) {
		// read the MaintenancePlan
		var maintenancePlan = new MaintenancePlanService(this.http).getMaintenancePlan(id);
		// add the MaintenancePlan if not already assigned
		if ( this.asset.maintenancePlans.indexOf(maintenancePlan) == -1 )
		this.asset.maintenancePlans.push(maintenancePlan);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more maintenancePlansIds as a MaintenancePlans
	// from a Asset
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMaintenancePlans( assetId, maintenancePlansIds ): Observable<any> {

		// get the Asset
		this.loadHelper( assetId );


	// split on a comma with no spaces
	var idList 					= maintenancePlansIds.split(',');
	var maintenancePlans 	= this.asset.maintenancePlans;

	if ( maintenancePlans != null && maintenancePlansIds != null ) {

		// iterate over array of maintenancePlans ids
		maintenancePlans.forEach(function (obj) {
			if ( maintenancePlansIds.indexOf(obj._id) > -1 ) {
				// remove the MaintenancePlan
				this.asset.maintenancePlans.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Asset
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Asset/update/' + this.asset;

	return  this.http.post(uri_, this.asset );
}

	//********************************************************************
	// loadHelper - internal helper to load a Asset
	//********************************************************************	
	loadHelper( id ) {
		this.getAsset(id)
			.subscribe((res : Asset) => {
				this.asset = res;
			});
	}
}