import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Routing} from '../models/Routing';
import {ItemService} from '../services/Item.service';
import {OperationService} from '../services/Operation.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class RoutingService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	routing : Routing;

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
	// add a Routing
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addRouting(routingNumber, revision, effectivityStart, effectivityEnd, Item, Operations, RoutingType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Routing/create';
		const obj = {
			      		routingNumber: routingNumber,
      		revision: revision,
      		effectivityStart: effectivityStart,
      		effectivityEnd: effectivityEnd,
      		Item: Item != null && Item.length > 0 ? Item : null,
      		Operations: Operations != null && Operations.length > 0 ? Operations : null,
      		RoutingType: RoutingType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Routing
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateRouting(routingNumber, revision, effectivityStart, effectivityEnd, Item, Operations, RoutingType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Routing/update/' + id;
		const obj = {
				      		routingNumber: routingNumber,
      		revision: revision,
      		effectivityStart: effectivityStart,
      		effectivityEnd: effectivityEnd,
      		Item: Item != null && Item.length > 0 ? Item : null,
      		Operations: Operations != null && Operations.length > 0 ? Operations : null,
      		RoutingType: RoutingType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Routing
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteRouting(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Routing/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Routing
	// returns the results untouched as an Observable Routing
	// Routing model
	// delegates via URI
	//********************************************************************
	getRouting(id) : Observable<Routing> {
		const uri_ = this.apiUrl + '/Routing/load/' + id;

		return this.http.get<Routing>(uri_);
	}
	
	//********************************************************************
	// gets all Routing
	// returns the results untouched as JSON representation of an
	// Observable array of Routing models
	// delegates via URI
	//********************************************************************
	getRoutings() : Observable<Routing[]> {
		const uri_ = this.apiUrl + '/Routing/';

		return this
			.http.get<Routing[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Item on a Routing
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignItem( routingId, _itemId ): Observable<any> {

		// get the Routing from storage
		this.loadHelper( routingId );

	// get the Item from storage
	var tmp 	= new ItemService(this.http).getItem(_itemId);

	// assign the Item
	this.routing.item = tmp;

	// save the Routing
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Item on a Routing
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignItem( routingId ): Observable<any> {

		// get the Routing from storage
		this.loadHelper( routingId );

	// assign Item to null
	this.routing.item = null;

	// save the Routing
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more operationsIds as a Operations
	// to a Routing
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOperations( routingId, operationsIds ): Observable<any> {

		// get the Routing
		this.loadHelper( routingId );

	// split on a comma with no spaces
	var idList = operationsIds.split(',')

	// iterate over array of operations ids
	idList.forEach(function (id) {
		// read the Operation
		var operation = new OperationService(this.http).getOperation(id);
		// add the Operation if not already assigned
		if ( this.routing.operations.indexOf(operation) == -1 )
		this.routing.operations.push(operation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more operationsIds as a Operations
	// from a Routing
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOperations( routingId, operationsIds ): Observable<any> {

		// get the Routing
		this.loadHelper( routingId );


	// split on a comma with no spaces
	var idList 					= operationsIds.split(',');
	var operations 	= this.routing.operations;

	if ( operations != null && operationsIds != null ) {

		// iterate over array of operations ids
		operations.forEach(function (obj) {
			if ( operationsIds.indexOf(obj._id) > -1 ) {
				// remove the Operation
				this.routing.operations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Routing
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Routing/update/' + this.routing;

	return  this.http.post(uri_, this.routing );
}

	//********************************************************************
	// loadHelper - internal helper to load a Routing
	//********************************************************************	
	loadHelper( id ) {
		this.getRouting(id)
			.subscribe((res : Routing) => {
				this.routing = res;
			});
	}
}