import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Operation} from '../models/Operation';
import {RoutingService} from '../services/Routing.service';
import {WorkCenterService} from '../services/WorkCenter.service';
import {InspectionPlanService} from '../services/InspectionPlan.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class OperationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	operation : Operation;

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
	// add a Operation
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addOperation(operationNumber, name, setupTime, standardCycleTime, Routing, WorkCenter, InspectionPlan, OperationType) : Observable<any> {
		const uri_ = this.apiUrl + '/Operation/create';
		const obj = {
			      		operationNumber: operationNumber,
      		name: name,
      		setupTime: setupTime,
      		standardCycleTime: standardCycleTime,
      		Routing: Routing != null && Routing.length > 0 ? Routing : null,
      		WorkCenter: WorkCenter != null && WorkCenter.length > 0 ? WorkCenter : null,
      		InspectionPlan: InspectionPlan != null && InspectionPlan.length > 0 ? InspectionPlan : null,
			OperationType: OperationType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Operation
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateOperation(operationNumber, name, setupTime, standardCycleTime, Routing, WorkCenter, InspectionPlan, OperationType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Operation/update/' + id;
		const obj = {
				      		operationNumber: operationNumber,
      		name: name,
      		setupTime: setupTime,
      		standardCycleTime: standardCycleTime,
      		Routing: Routing != null && Routing.length > 0 ? Routing : null,
      		WorkCenter: WorkCenter != null && WorkCenter.length > 0 ? WorkCenter : null,
      		InspectionPlan: InspectionPlan != null && InspectionPlan.length > 0 ? InspectionPlan : null,
			OperationType: OperationType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Operation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteOperation(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Operation/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Operation
	// returns the results untouched as an Observable Operation
	// Operation model
	// delegates via URI
	//********************************************************************
	getOperation(id) : Observable<Operation> {
		const uri_ = this.apiUrl + '/Operation/load/' + id;

		return this.http.get<Operation>(uri_);
	}
	
	//********************************************************************
	// gets all Operation
	// returns the results untouched as JSON representation of an
	// Observable array of Operation models
	// delegates via URI
	//********************************************************************
	getOperations() : Observable<Operation[]> {
		const uri_ = this.apiUrl + '/Operation/';

		return this
			.http.get<Operation[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Routing on a Operation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRouting( operationId, _routingId ): Observable<any> {

		// get the Operation from storage
		this.loadHelper( operationId );

	// get the Routing from storage
	var tmp 	= new RoutingService(this.http).getRouting(_routingId);

	// assign the Routing
	this.operation.routing = tmp;

	// save the Operation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Routing on a Operation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRouting( operationId ): Observable<any> {

		// get the Operation from storage
		this.loadHelper( operationId );

	// assign Routing to null
	this.operation.routing = null;

	// save the Operation
	return this.saveHelper();
}

		//********************************************************************
	// assigns a WorkCenter on a Operation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkCenter( operationId, _workCenterId ): Observable<any> {

		// get the Operation from storage
		this.loadHelper( operationId );

	// get the WorkCenter from storage
	var tmp 	= new WorkCenterService(this.http).getWorkCenter(_workCenterId);

	// assign the WorkCenter
	this.operation.workCenter = tmp;

	// save the Operation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a WorkCenter on a Operation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkCenter( operationId ): Observable<any> {

		// get the Operation from storage
		this.loadHelper( operationId );

	// assign WorkCenter to null
	this.operation.workCenter = null;

	// save the Operation
	return this.saveHelper();
}

		//********************************************************************
	// assigns a InspectionPlan on a Operation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInspectionPlan( operationId, _inspectionPlanId ): Observable<any> {

		// get the Operation from storage
		this.loadHelper( operationId );

	// get the InspectionPlan from storage
	var tmp 	= new InspectionPlanService(this.http).getInspectionPlan(_inspectionPlanId);

	// assign the InspectionPlan
	this.operation.inspectionPlan = tmp;

	// save the Operation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a InspectionPlan on a Operation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInspectionPlan( operationId ): Observable<any> {

		// get the Operation from storage
		this.loadHelper( operationId );

	// assign InspectionPlan to null
	this.operation.inspectionPlan = null;

	// save the Operation
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Operation
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Operation/update/' + this.operation;

	return  this.http.post(uri_, this.operation );
}

	//********************************************************************
	// loadHelper - internal helper to load a Operation
	//********************************************************************	
	loadHelper( id ) {
		this.getOperation(id)
			.subscribe((res : Operation) => {
				this.operation = res;
			});
	}
}