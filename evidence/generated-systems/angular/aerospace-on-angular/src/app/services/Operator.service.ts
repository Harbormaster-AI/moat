import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Operator} from '../models/Operator';
import {AircraftOrderService} from '../services/AircraftOrder.service';
import {AircraftService} from '../services/Aircraft.service';
import {SalesRegionService} from '../services/SalesRegion.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class OperatorService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	operator : Operator;

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
	// add a Operator
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addOperator(name, icaoDesignator, AircraftOrders, OperatedAircraft, SalesRegion, OperatorType) : Observable<any> {
		const uri_ = this.apiUrl + '/Operator/create';
		const obj = {
			      		name: name,
      		icaoDesignator: icaoDesignator,
      		AircraftOrders: AircraftOrders != null && AircraftOrders.length > 0 ? AircraftOrders : null,
      		OperatedAircraft: OperatedAircraft != null && OperatedAircraft.length > 0 ? OperatedAircraft : null,
      		SalesRegion: SalesRegion != null && SalesRegion.length > 0 ? SalesRegion : null,
			OperatorType: OperatorType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Operator
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateOperator(name, icaoDesignator, AircraftOrders, OperatedAircraft, SalesRegion, OperatorType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Operator/update/' + id;
		const obj = {
				      		name: name,
      		icaoDesignator: icaoDesignator,
      		AircraftOrders: AircraftOrders != null && AircraftOrders.length > 0 ? AircraftOrders : null,
      		OperatedAircraft: OperatedAircraft != null && OperatedAircraft.length > 0 ? OperatedAircraft : null,
      		SalesRegion: SalesRegion != null && SalesRegion.length > 0 ? SalesRegion : null,
			OperatorType: OperatorType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Operator
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteOperator(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Operator/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Operator
	// returns the results untouched as an Observable Operator
	// Operator model
	// delegates via URI
	//********************************************************************
	getOperator(id) : Observable<Operator> {
		const uri_ = this.apiUrl + '/Operator/load/' + id;

		return this.http.get<Operator>(uri_);
	}
	
	//********************************************************************
	// gets all Operator
	// returns the results untouched as JSON representation of an
	// Observable array of Operator models
	// delegates via URI
	//********************************************************************
	getOperators() : Observable<Operator[]> {
		const uri_ = this.apiUrl + '/Operator/';

		return this
			.http.get<Operator[]>(uri_);
	}
	
			//********************************************************************
	// assigns a SalesRegion on a Operator
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSalesRegion( operatorId, _salesRegionId ): Observable<any> {

		// get the Operator from storage
		this.loadHelper( operatorId );

	// get the SalesRegion from storage
	var tmp 	= new SalesRegionService(this.http).getSalesRegion(_salesRegionId);

	// assign the SalesRegion
	this.operator.salesRegion = tmp;

	// save the Operator
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a SalesRegion on a Operator
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSalesRegion( operatorId ): Observable<any> {

		// get the Operator from storage
		this.loadHelper( operatorId );

	// assign SalesRegion to null
	this.operator.salesRegion = null;

	// save the Operator
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more aircraftOrdersIds as a AircraftOrders
	// to a Operator
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAircraftOrders( operatorId, aircraftOrdersIds ): Observable<any> {

		// get the Operator
		this.loadHelper( operatorId );

	// split on a comma with no spaces
	var idList = aircraftOrdersIds.split(',')

	// iterate over array of aircraftOrders ids
	idList.forEach(function (id) {
		// read the AircraftOrder
		var aircraftOrder = new AircraftOrderService(this.http).getAircraftOrder(id);
		// add the AircraftOrder if not already assigned
		if ( this.operator.aircraftOrders.indexOf(aircraftOrder) == -1 )
		this.operator.aircraftOrders.push(aircraftOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more aircraftOrdersIds as a AircraftOrders
	// from a Operator
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAircraftOrders( operatorId, aircraftOrdersIds ): Observable<any> {

		// get the Operator
		this.loadHelper( operatorId );


	// split on a comma with no spaces
	var idList 					= aircraftOrdersIds.split(',');
	var aircraftOrders 	= this.operator.aircraftOrders;

	if ( aircraftOrders != null && aircraftOrdersIds != null ) {

		// iterate over array of aircraftOrders ids
		aircraftOrders.forEach(function (obj) {
			if ( aircraftOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the AircraftOrder
				this.operator.aircraftOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more operatedAircraftIds as a OperatedAircraft
	// to a Operator
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOperatedAircraft( operatorId, operatedAircraftIds ): Observable<any> {

		// get the Operator
		this.loadHelper( operatorId );

	// split on a comma with no spaces
	var idList = operatedAircraftIds.split(',')

	// iterate over array of operatedAircraft ids
	idList.forEach(function (id) {
		// read the Aircraft
		var aircraft = new AircraftService(this.http).getAircraft(id);
		// add the Aircraft if not already assigned
		if ( this.operator.operatedAircraft.indexOf(aircraft) == -1 )
		this.operator.operatedAircraft.push(aircraft);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more operatedAircraftIds as a OperatedAircraft
	// from a Operator
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOperatedAircraft( operatorId, operatedAircraftIds ): Observable<any> {

		// get the Operator
		this.loadHelper( operatorId );


	// split on a comma with no spaces
	var idList 					= operatedAircraftIds.split(',');
	var operatedAircraft 	= this.operator.operatedAircraft;

	if ( operatedAircraft != null && operatedAircraftIds != null ) {

		// iterate over array of operatedAircraft ids
		operatedAircraft.forEach(function (obj) {
			if ( operatedAircraftIds.indexOf(obj._id) > -1 ) {
				// remove the Aircraft
				this.operator.operatedAircraft.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Operator
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Operator/update/' + this.operator;

	return  this.http.post(uri_, this.operator );
}

	//********************************************************************
	// loadHelper - internal helper to load a Operator
	//********************************************************************	
	loadHelper( id ) {
		this.getOperator(id)
			.subscribe((res : Operator) => {
				this.operator = res;
			});
	}
}