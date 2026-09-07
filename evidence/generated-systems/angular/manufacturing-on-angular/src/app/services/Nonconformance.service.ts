import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Nonconformance} from '../models/Nonconformance';
import {ItemService} from '../services/Item.service';
import {WorkOrderService} from '../services/WorkOrder.service';
import {InspectionLotService} from '../services/InspectionLot.service';
import {CorrectiveActionService} from '../services/CorrectiveAction.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class NonconformanceService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	nonconformance : Nonconformance;

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
	// add a Nonconformance
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addNonconformance(ncNumber, description, containmentAction, Item, WorkOrder, InspectionLot, CorrectiveAction, NcType, Severity, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Nonconformance/create';
		const obj = {
			      		ncNumber: ncNumber,
      		description: description,
      		containmentAction: containmentAction,
      		Item: Item != null && Item.length > 0 ? Item : null,
      		WorkOrder: WorkOrder != null && WorkOrder.length > 0 ? WorkOrder : null,
      		InspectionLot: InspectionLot != null && InspectionLot.length > 0 ? InspectionLot : null,
      		CorrectiveAction: CorrectiveAction != null && CorrectiveAction.length > 0 ? CorrectiveAction : null,
      		NcType: NcType,
      		Severity: Severity,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Nonconformance
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateNonconformance(ncNumber, description, containmentAction, Item, WorkOrder, InspectionLot, CorrectiveAction, NcType, Severity, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Nonconformance/update/' + id;
		const obj = {
				      		ncNumber: ncNumber,
      		description: description,
      		containmentAction: containmentAction,
      		Item: Item != null && Item.length > 0 ? Item : null,
      		WorkOrder: WorkOrder != null && WorkOrder.length > 0 ? WorkOrder : null,
      		InspectionLot: InspectionLot != null && InspectionLot.length > 0 ? InspectionLot : null,
      		CorrectiveAction: CorrectiveAction != null && CorrectiveAction.length > 0 ? CorrectiveAction : null,
      		NcType: NcType,
      		Severity: Severity,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Nonconformance
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteNonconformance(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Nonconformance/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Nonconformance
	// returns the results untouched as an Observable Nonconformance
	// Nonconformance model
	// delegates via URI
	//********************************************************************
	getNonconformance(id) : Observable<Nonconformance> {
		const uri_ = this.apiUrl + '/Nonconformance/load/' + id;

		return this.http.get<Nonconformance>(uri_);
	}
	
	//********************************************************************
	// gets all Nonconformance
	// returns the results untouched as JSON representation of an
	// Observable array of Nonconformance models
	// delegates via URI
	//********************************************************************
	getNonconformances() : Observable<Nonconformance[]> {
		const uri_ = this.apiUrl + '/Nonconformance/';

		return this
			.http.get<Nonconformance[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Item on a Nonconformance
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignItem( nonconformanceId, _itemId ): Observable<any> {

		// get the Nonconformance from storage
		this.loadHelper( nonconformanceId );

	// get the Item from storage
	var tmp 	= new ItemService(this.http).getItem(_itemId);

	// assign the Item
	this.nonconformance.item = tmp;

	// save the Nonconformance
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Item on a Nonconformance
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignItem( nonconformanceId ): Observable<any> {

		// get the Nonconformance from storage
		this.loadHelper( nonconformanceId );

	// assign Item to null
	this.nonconformance.item = null;

	// save the Nonconformance
	return this.saveHelper();
}

		//********************************************************************
	// assigns a WorkOrder on a Nonconformance
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkOrder( nonconformanceId, _workOrderId ): Observable<any> {

		// get the Nonconformance from storage
		this.loadHelper( nonconformanceId );

	// get the WorkOrder from storage
	var tmp 	= new WorkOrderService(this.http).getWorkOrder(_workOrderId);

	// assign the WorkOrder
	this.nonconformance.workOrder = tmp;

	// save the Nonconformance
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a WorkOrder on a Nonconformance
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkOrder( nonconformanceId ): Observable<any> {

		// get the Nonconformance from storage
		this.loadHelper( nonconformanceId );

	// assign WorkOrder to null
	this.nonconformance.workOrder = null;

	// save the Nonconformance
	return this.saveHelper();
}

		//********************************************************************
	// assigns a InspectionLot on a Nonconformance
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInspectionLot( nonconformanceId, _inspectionLotId ): Observable<any> {

		// get the Nonconformance from storage
		this.loadHelper( nonconformanceId );

	// get the InspectionLot from storage
	var tmp 	= new InspectionLotService(this.http).getInspectionLot(_inspectionLotId);

	// assign the InspectionLot
	this.nonconformance.inspectionLot = tmp;

	// save the Nonconformance
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a InspectionLot on a Nonconformance
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInspectionLot( nonconformanceId ): Observable<any> {

		// get the Nonconformance from storage
		this.loadHelper( nonconformanceId );

	// assign InspectionLot to null
	this.nonconformance.inspectionLot = null;

	// save the Nonconformance
	return this.saveHelper();
}

		//********************************************************************
	// assigns a CorrectiveAction on a Nonconformance
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCorrectiveAction( nonconformanceId, _correctiveActionId ): Observable<any> {

		// get the Nonconformance from storage
		this.loadHelper( nonconformanceId );

	// get the CorrectiveAction from storage
	var tmp 	= new CorrectiveActionService(this.http).getCorrectiveAction(_correctiveActionId);

	// assign the CorrectiveAction
	this.nonconformance.correctiveAction = tmp;

	// save the Nonconformance
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CorrectiveAction on a Nonconformance
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCorrectiveAction( nonconformanceId ): Observable<any> {

		// get the Nonconformance from storage
		this.loadHelper( nonconformanceId );

	// assign CorrectiveAction to null
	this.nonconformance.correctiveAction = null;

	// save the Nonconformance
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Nonconformance
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Nonconformance/update/' + this.nonconformance;

	return  this.http.post(uri_, this.nonconformance );
}

	//********************************************************************
	// loadHelper - internal helper to load a Nonconformance
	//********************************************************************	
	loadHelper( id ) {
		this.getNonconformance(id)
			.subscribe((res : Nonconformance) => {
				this.nonconformance = res;
			});
	}
}