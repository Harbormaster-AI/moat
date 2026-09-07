import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ProductionLine} from '../models/ProductionLine';
import {PlantService} from '../services/Plant.service';
import {WorkCenterService} from '../services/WorkCenter.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ProductionLineService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	productionLine : ProductionLine;

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
	// add a ProductionLine
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addProductionLine(name, Plant, WorkCenters, LineType) : Observable<any> {
		const uri_ = this.apiUrl + '/ProductionLine/create';
		const obj = {
			      		name: name,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		WorkCenters: WorkCenters != null && WorkCenters.length > 0 ? WorkCenters : null,
			LineType: LineType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ProductionLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateProductionLine(name, Plant, WorkCenters, LineType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ProductionLine/update/' + id;
		const obj = {
				      		name: name,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		WorkCenters: WorkCenters != null && WorkCenters.length > 0 ? WorkCenters : null,
			LineType: LineType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ProductionLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteProductionLine(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ProductionLine/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ProductionLine
	// returns the results untouched as an Observable ProductionLine
	// ProductionLine model
	// delegates via URI
	//********************************************************************
	getProductionLine(id) : Observable<ProductionLine> {
		const uri_ = this.apiUrl + '/ProductionLine/load/' + id;

		return this.http.get<ProductionLine>(uri_);
	}
	
	//********************************************************************
	// gets all ProductionLine
	// returns the results untouched as JSON representation of an
	// Observable array of ProductionLine models
	// delegates via URI
	//********************************************************************
	getProductionLines() : Observable<ProductionLine[]> {
		const uri_ = this.apiUrl + '/ProductionLine/';

		return this
			.http.get<ProductionLine[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Plant on a ProductionLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPlant( productionLineId, _plantId ): Observable<any> {

		// get the ProductionLine from storage
		this.loadHelper( productionLineId );

	// get the Plant from storage
	var tmp 	= new PlantService(this.http).getPlant(_plantId);

	// assign the Plant
	this.productionLine.plant = tmp;

	// save the ProductionLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Plant on a ProductionLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPlant( productionLineId ): Observable<any> {

		// get the ProductionLine from storage
		this.loadHelper( productionLineId );

	// assign Plant to null
	this.productionLine.plant = null;

	// save the ProductionLine
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more workCentersIds as a WorkCenters
	// to a ProductionLine
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addWorkCenters( productionLineId, workCentersIds ): Observable<any> {

		// get the ProductionLine
		this.loadHelper( productionLineId );

	// split on a comma with no spaces
	var idList = workCentersIds.split(',')

	// iterate over array of workCenters ids
	idList.forEach(function (id) {
		// read the WorkCenter
		var workCenter = new WorkCenterService(this.http).getWorkCenter(id);
		// add the WorkCenter if not already assigned
		if ( this.productionLine.workCenters.indexOf(workCenter) == -1 )
		this.productionLine.workCenters.push(workCenter);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more workCentersIds as a WorkCenters
	// from a ProductionLine
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeWorkCenters( productionLineId, workCentersIds ): Observable<any> {

		// get the ProductionLine
		this.loadHelper( productionLineId );


	// split on a comma with no spaces
	var idList 					= workCentersIds.split(',');
	var workCenters 	= this.productionLine.workCenters;

	if ( workCenters != null && workCentersIds != null ) {

		// iterate over array of workCenters ids
		workCenters.forEach(function (obj) {
			if ( workCentersIds.indexOf(obj._id) > -1 ) {
				// remove the WorkCenter
				this.productionLine.workCenters.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a ProductionLine
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ProductionLine/update/' + this.productionLine;

	return  this.http.post(uri_, this.productionLine );
}

	//********************************************************************
	// loadHelper - internal helper to load a ProductionLine
	//********************************************************************	
	loadHelper( id ) {
		this.getProductionLine(id)
			.subscribe((res : ProductionLine) => {
				this.productionLine = res;
			});
	}
}