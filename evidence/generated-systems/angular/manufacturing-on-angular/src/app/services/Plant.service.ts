import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Plant} from '../models/Plant';
import {EnterpriseService} from '../services/Enterprise.service';
import {ProductionLineService} from '../services/ProductionLine.service';
import {WorkCenterService} from '../services/WorkCenter.service';
import {WarehouseService} from '../services/Warehouse.service';
import {AssetService} from '../services/Asset.service';
import {ProductionScheduleService} from '../services/ProductionSchedule.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PlantService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	plant : Plant;

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
	// add a Plant
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPlant(name, plantCode, address, timeZone, Enterprise, ProductionLines, WorkCenters, Warehouses, Assets, ProductionSchedules) : Observable<any> {
		const uri_ = this.apiUrl + '/Plant/create';
		const obj = {
			      		name: name,
      		plantCode: plantCode,
      		address: address,
      		timeZone: timeZone,
      		Enterprise: Enterprise != null && Enterprise.length > 0 ? Enterprise : null,
      		ProductionLines: ProductionLines != null && ProductionLines.length > 0 ? ProductionLines : null,
      		WorkCenters: WorkCenters != null && WorkCenters.length > 0 ? WorkCenters : null,
      		Warehouses: Warehouses != null && Warehouses.length > 0 ? Warehouses : null,
      		Assets: Assets != null && Assets.length > 0 ? Assets : null,
			ProductionSchedules: ProductionSchedules != null && ProductionSchedules.length > 0 ? ProductionSchedules : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Plant
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePlant(name, plantCode, address, timeZone, Enterprise, ProductionLines, WorkCenters, Warehouses, Assets, ProductionSchedules, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Plant/update/' + id;
		const obj = {
				      		name: name,
      		plantCode: plantCode,
      		address: address,
      		timeZone: timeZone,
      		Enterprise: Enterprise != null && Enterprise.length > 0 ? Enterprise : null,
      		ProductionLines: ProductionLines != null && ProductionLines.length > 0 ? ProductionLines : null,
      		WorkCenters: WorkCenters != null && WorkCenters.length > 0 ? WorkCenters : null,
      		Warehouses: Warehouses != null && Warehouses.length > 0 ? Warehouses : null,
      		Assets: Assets != null && Assets.length > 0 ? Assets : null,
			ProductionSchedules: ProductionSchedules != null && ProductionSchedules.length > 0 ? ProductionSchedules : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Plant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePlant(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Plant/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Plant
	// returns the results untouched as an Observable Plant
	// Plant model
	// delegates via URI
	//********************************************************************
	getPlant(id) : Observable<Plant> {
		const uri_ = this.apiUrl + '/Plant/load/' + id;

		return this.http.get<Plant>(uri_);
	}
	
	//********************************************************************
	// gets all Plant
	// returns the results untouched as JSON representation of an
	// Observable array of Plant models
	// delegates via URI
	//********************************************************************
	getPlants() : Observable<Plant[]> {
		const uri_ = this.apiUrl + '/Plant/';

		return this
			.http.get<Plant[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Enterprise on a Plant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEnterprise( plantId, _enterpriseId ): Observable<any> {

		// get the Plant from storage
		this.loadHelper( plantId );

	// get the Enterprise from storage
	var tmp 	= new EnterpriseService(this.http).getEnterprise(_enterpriseId);

	// assign the Enterprise
	this.plant.enterprise = tmp;

	// save the Plant
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Enterprise on a Plant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEnterprise( plantId ): Observable<any> {

		// get the Plant from storage
		this.loadHelper( plantId );

	// assign Enterprise to null
	this.plant.enterprise = null;

	// save the Plant
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more productionLinesIds as a ProductionLines
	// to a Plant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProductionLines( plantId, productionLinesIds ): Observable<any> {

		// get the Plant
		this.loadHelper( plantId );

	// split on a comma with no spaces
	var idList = productionLinesIds.split(',')

	// iterate over array of productionLines ids
	idList.forEach(function (id) {
		// read the ProductionLine
		var productionLine = new ProductionLineService(this.http).getProductionLine(id);
		// add the ProductionLine if not already assigned
		if ( this.plant.productionLines.indexOf(productionLine) == -1 )
		this.plant.productionLines.push(productionLine);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more productionLinesIds as a ProductionLines
	// from a Plant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProductionLines( plantId, productionLinesIds ): Observable<any> {

		// get the Plant
		this.loadHelper( plantId );


	// split on a comma with no spaces
	var idList 					= productionLinesIds.split(',');
	var productionLines 	= this.plant.productionLines;

	if ( productionLines != null && productionLinesIds != null ) {

		// iterate over array of productionLines ids
		productionLines.forEach(function (obj) {
			if ( productionLinesIds.indexOf(obj._id) > -1 ) {
				// remove the ProductionLine
				this.plant.productionLines.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more workCentersIds as a WorkCenters
	// to a Plant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addWorkCenters( plantId, workCentersIds ): Observable<any> {

		// get the Plant
		this.loadHelper( plantId );

	// split on a comma with no spaces
	var idList = workCentersIds.split(',')

	// iterate over array of workCenters ids
	idList.forEach(function (id) {
		// read the WorkCenter
		var workCenter = new WorkCenterService(this.http).getWorkCenter(id);
		// add the WorkCenter if not already assigned
		if ( this.plant.workCenters.indexOf(workCenter) == -1 )
		this.plant.workCenters.push(workCenter);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more workCentersIds as a WorkCenters
	// from a Plant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeWorkCenters( plantId, workCentersIds ): Observable<any> {

		// get the Plant
		this.loadHelper( plantId );


	// split on a comma with no spaces
	var idList 					= workCentersIds.split(',');
	var workCenters 	= this.plant.workCenters;

	if ( workCenters != null && workCentersIds != null ) {

		// iterate over array of workCenters ids
		workCenters.forEach(function (obj) {
			if ( workCentersIds.indexOf(obj._id) > -1 ) {
				// remove the WorkCenter
				this.plant.workCenters.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more warehousesIds as a Warehouses
	// to a Plant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addWarehouses( plantId, warehousesIds ): Observable<any> {

		// get the Plant
		this.loadHelper( plantId );

	// split on a comma with no spaces
	var idList = warehousesIds.split(',')

	// iterate over array of warehouses ids
	idList.forEach(function (id) {
		// read the Warehouse
		var warehouse = new WarehouseService(this.http).getWarehouse(id);
		// add the Warehouse if not already assigned
		if ( this.plant.warehouses.indexOf(warehouse) == -1 )
		this.plant.warehouses.push(warehouse);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more warehousesIds as a Warehouses
	// from a Plant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeWarehouses( plantId, warehousesIds ): Observable<any> {

		// get the Plant
		this.loadHelper( plantId );


	// split on a comma with no spaces
	var idList 					= warehousesIds.split(',');
	var warehouses 	= this.plant.warehouses;

	if ( warehouses != null && warehousesIds != null ) {

		// iterate over array of warehouses ids
		warehouses.forEach(function (obj) {
			if ( warehousesIds.indexOf(obj._id) > -1 ) {
				// remove the Warehouse
				this.plant.warehouses.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more assetsIds as a Assets
	// to a Plant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAssets( plantId, assetsIds ): Observable<any> {

		// get the Plant
		this.loadHelper( plantId );

	// split on a comma with no spaces
	var idList = assetsIds.split(',')

	// iterate over array of assets ids
	idList.forEach(function (id) {
		// read the Asset
		var asset = new AssetService(this.http).getAsset(id);
		// add the Asset if not already assigned
		if ( this.plant.assets.indexOf(asset) == -1 )
		this.plant.assets.push(asset);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more assetsIds as a Assets
	// from a Plant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAssets( plantId, assetsIds ): Observable<any> {

		// get the Plant
		this.loadHelper( plantId );


	// split on a comma with no spaces
	var idList 					= assetsIds.split(',');
	var assets 	= this.plant.assets;

	if ( assets != null && assetsIds != null ) {

		// iterate over array of assets ids
		assets.forEach(function (obj) {
			if ( assetsIds.indexOf(obj._id) > -1 ) {
				// remove the Asset
				this.plant.assets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more productionSchedulesIds as a ProductionSchedules
	// to a Plant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProductionSchedules( plantId, productionSchedulesIds ): Observable<any> {

		// get the Plant
		this.loadHelper( plantId );

	// split on a comma with no spaces
	var idList = productionSchedulesIds.split(',')

	// iterate over array of productionSchedules ids
	idList.forEach(function (id) {
		// read the ProductionSchedule
		var productionSchedule = new ProductionScheduleService(this.http).getProductionSchedule(id);
		// add the ProductionSchedule if not already assigned
		if ( this.plant.productionSchedules.indexOf(productionSchedule) == -1 )
		this.plant.productionSchedules.push(productionSchedule);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more productionSchedulesIds as a ProductionSchedules
	// from a Plant
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProductionSchedules( plantId, productionSchedulesIds ): Observable<any> {

		// get the Plant
		this.loadHelper( plantId );


	// split on a comma with no spaces
	var idList 					= productionSchedulesIds.split(',');
	var productionSchedules 	= this.plant.productionSchedules;

	if ( productionSchedules != null && productionSchedulesIds != null ) {

		// iterate over array of productionSchedules ids
		productionSchedules.forEach(function (obj) {
			if ( productionSchedulesIds.indexOf(obj._id) > -1 ) {
				// remove the ProductionSchedule
				this.plant.productionSchedules.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Plant
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Plant/update/' + this.plant;

	return  this.http.post(uri_, this.plant );
}

	//********************************************************************
	// loadHelper - internal helper to load a Plant
	//********************************************************************	
	loadHelper( id ) {
		this.getPlant(id)
			.subscribe((res : Plant) => {
				this.plant = res;
			});
	}
}