import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {BuildSchedule} from '../models/BuildSchedule';
import {ProductionOrderService} from '../services/ProductionOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class BuildScheduleService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	buildSchedule : BuildSchedule;

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
	// add a BuildSchedule
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addBuildSchedule(scheduleNumber, ProductionOrders, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/BuildSchedule/create';
		const obj = {
			      		scheduleNumber: scheduleNumber,
      		ProductionOrders: ProductionOrders != null && ProductionOrders.length > 0 ? ProductionOrders : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a BuildSchedule
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateBuildSchedule(scheduleNumber, ProductionOrders, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/BuildSchedule/update/' + id;
		const obj = {
				      		scheduleNumber: scheduleNumber,
      		ProductionOrders: ProductionOrders != null && ProductionOrders.length > 0 ? ProductionOrders : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a BuildSchedule
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteBuildSchedule(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/BuildSchedule/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a BuildSchedule
	// returns the results untouched as an Observable BuildSchedule
	// BuildSchedule model
	// delegates via URI
	//********************************************************************
	getBuildSchedule(id) : Observable<BuildSchedule> {
		const uri_ = this.apiUrl + '/BuildSchedule/load/' + id;

		return this.http.get<BuildSchedule>(uri_);
	}
	
	//********************************************************************
	// gets all BuildSchedule
	// returns the results untouched as JSON representation of an
	// Observable array of BuildSchedule models
	// delegates via URI
	//********************************************************************
	getBuildSchedules() : Observable<BuildSchedule[]> {
		const uri_ = this.apiUrl + '/BuildSchedule/';

		return this
			.http.get<BuildSchedule[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more productionOrdersIds as a ProductionOrders
	// to a BuildSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProductionOrders( buildScheduleId, productionOrdersIds ): Observable<any> {

		// get the BuildSchedule
		this.loadHelper( buildScheduleId );

	// split on a comma with no spaces
	var idList = productionOrdersIds.split(',')

	// iterate over array of productionOrders ids
	idList.forEach(function (id) {
		// read the ProductionOrder
		var productionOrder = new ProductionOrderService(this.http).getProductionOrder(id);
		// add the ProductionOrder if not already assigned
		if ( this.buildSchedule.productionOrders.indexOf(productionOrder) == -1 )
		this.buildSchedule.productionOrders.push(productionOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more productionOrdersIds as a ProductionOrders
	// from a BuildSchedule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProductionOrders( buildScheduleId, productionOrdersIds ): Observable<any> {

		// get the BuildSchedule
		this.loadHelper( buildScheduleId );


	// split on a comma with no spaces
	var idList 					= productionOrdersIds.split(',');
	var productionOrders 	= this.buildSchedule.productionOrders;

	if ( productionOrders != null && productionOrdersIds != null ) {

		// iterate over array of productionOrders ids
		productionOrders.forEach(function (obj) {
			if ( productionOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the ProductionOrder
				this.buildSchedule.productionOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a BuildSchedule
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/BuildSchedule/update/' + this.buildSchedule;

	return  this.http.post(uri_, this.buildSchedule );
}

	//********************************************************************
	// loadHelper - internal helper to load a BuildSchedule
	//********************************************************************	
	loadHelper( id ) {
		this.getBuildSchedule(id)
			.subscribe((res : BuildSchedule) => {
				this.buildSchedule = res;
			});
	}
}