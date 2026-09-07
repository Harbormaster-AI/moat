import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {MRPRun} from '../models/MRPRun';
import {PlantService} from '../services/Plant.service';
import {PlannedOrderService} from '../services/PlannedOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class MRPRunService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	mRPRun : MRPRun;

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
	// add a MRPRun
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addMRPRun(runNumber, runDateTime, planningHorizonDays, Plant, PlannedOrders, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/MRPRun/create';
		const obj = {
			      		runNumber: runNumber,
      		runDateTime: runDateTime,
      		planningHorizonDays: planningHorizonDays,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		PlannedOrders: PlannedOrders != null && PlannedOrders.length > 0 ? PlannedOrders : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a MRPRun
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateMRPRun(runNumber, runDateTime, planningHorizonDays, Plant, PlannedOrders, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/MRPRun/update/' + id;
		const obj = {
				      		runNumber: runNumber,
      		runDateTime: runDateTime,
      		planningHorizonDays: planningHorizonDays,
      		Plant: Plant != null && Plant.length > 0 ? Plant : null,
      		PlannedOrders: PlannedOrders != null && PlannedOrders.length > 0 ? PlannedOrders : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a MRPRun
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteMRPRun(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/MRPRun/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a MRPRun
	// returns the results untouched as an Observable MRPRun
	// MRPRun model
	// delegates via URI
	//********************************************************************
	getMRPRun(id) : Observable<MRPRun> {
		const uri_ = this.apiUrl + '/MRPRun/load/' + id;

		return this.http.get<MRPRun>(uri_);
	}
	
	//********************************************************************
	// gets all MRPRun
	// returns the results untouched as JSON representation of an
	// Observable array of MRPRun models
	// delegates via URI
	//********************************************************************
	getMRPRuns() : Observable<MRPRun[]> {
		const uri_ = this.apiUrl + '/MRPRun/';

		return this
			.http.get<MRPRun[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Plant on a MRPRun
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPlant( mRPRunId, _plantId ): Observable<any> {

		// get the MRPRun from storage
		this.loadHelper( mRPRunId );

	// get the Plant from storage
	var tmp 	= new PlantService(this.http).getPlant(_plantId);

	// assign the Plant
	this.mRPRun.plant = tmp;

	// save the MRPRun
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Plant on a MRPRun
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPlant( mRPRunId ): Observable<any> {

		// get the MRPRun from storage
		this.loadHelper( mRPRunId );

	// assign Plant to null
	this.mRPRun.plant = null;

	// save the MRPRun
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more plannedOrdersIds as a PlannedOrders
	// to a MRPRun
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPlannedOrders( mRPRunId, plannedOrdersIds ): Observable<any> {

		// get the MRPRun
		this.loadHelper( mRPRunId );

	// split on a comma with no spaces
	var idList = plannedOrdersIds.split(',')

	// iterate over array of plannedOrders ids
	idList.forEach(function (id) {
		// read the PlannedOrder
		var plannedOrder = new PlannedOrderService(this.http).getPlannedOrder(id);
		// add the PlannedOrder if not already assigned
		if ( this.mRPRun.plannedOrders.indexOf(plannedOrder) == -1 )
		this.mRPRun.plannedOrders.push(plannedOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more plannedOrdersIds as a PlannedOrders
	// from a MRPRun
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePlannedOrders( mRPRunId, plannedOrdersIds ): Observable<any> {

		// get the MRPRun
		this.loadHelper( mRPRunId );


	// split on a comma with no spaces
	var idList 					= plannedOrdersIds.split(',');
	var plannedOrders 	= this.mRPRun.plannedOrders;

	if ( plannedOrders != null && plannedOrdersIds != null ) {

		// iterate over array of plannedOrders ids
		plannedOrders.forEach(function (obj) {
			if ( plannedOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the PlannedOrder
				this.mRPRun.plannedOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a MRPRun
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/MRPRun/update/' + this.mRPRun;

	return  this.http.post(uri_, this.mRPRun );
}

	//********************************************************************
	// loadHelper - internal helper to load a MRPRun
	//********************************************************************	
	loadHelper( id ) {
		this.getMRPRun(id)
			.subscribe((res : MRPRun) => {
				this.mRPRun = res;
			});
	}
}