import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Plant} from '../models/Plant';
import {AerospaceManufacturerService} from '../services/AerospaceManufacturer.service';
import {ProductionLineService} from '../services/ProductionLine.service';
import {WarehouseService} from '../services/Warehouse.service';
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
	addPlant(name, plantCode, address, Manufacturer, ProductionLines, Warehouses) : Observable<any> {
		const uri_ = this.apiUrl + '/Plant/create';
		const obj = {
			      		name: name,
      		plantCode: plantCode,
      		address: address,
      		Manufacturer: Manufacturer != null && Manufacturer.length > 0 ? Manufacturer : null,
      		ProductionLines: ProductionLines != null && ProductionLines.length > 0 ? ProductionLines : null,
			Warehouses: Warehouses != null && Warehouses.length > 0 ? Warehouses : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Plant
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePlant(name, plantCode, address, Manufacturer, ProductionLines, Warehouses, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Plant/update/' + id;
		const obj = {
				      		name: name,
      		plantCode: plantCode,
      		address: address,
      		Manufacturer: Manufacturer != null && Manufacturer.length > 0 ? Manufacturer : null,
      		ProductionLines: ProductionLines != null && ProductionLines.length > 0 ? ProductionLines : null,
			Warehouses: Warehouses != null && Warehouses.length > 0 ? Warehouses : null
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
	// assigns a Manufacturer on a Plant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignManufacturer( plantId, _manufacturerId ): Observable<any> {

		// get the Plant from storage
		this.loadHelper( plantId );

	// get the AerospaceManufacturer from storage
	var tmp 	= new AerospaceManufacturerService(this.http).getAerospaceManufacturer(_manufacturerId);

	// assign the Manufacturer
	this.plant.manufacturer = tmp;

	// save the Plant
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Manufacturer on a Plant
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignManufacturer( plantId ): Observable<any> {

		// get the Plant from storage
		this.loadHelper( plantId );

	// assign Manufacturer to null
	this.plant.manufacturer = null;

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