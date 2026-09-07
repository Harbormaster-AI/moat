import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {WorkCenter} from '../models/WorkCenter';
import {ProductionLineService} from '../services/ProductionLine.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class WorkCenterService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	workCenter : WorkCenter;

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
	// add a WorkCenter
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addWorkCenter(name, capability, ProductionLine) : Observable<any> {
		const uri_ = this.apiUrl + '/WorkCenter/create';
		const obj = {
			      		name: name,
      		capability: capability,
			ProductionLine: ProductionLine != null && ProductionLine.length > 0 ? ProductionLine : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a WorkCenter
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateWorkCenter(name, capability, ProductionLine, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/WorkCenter/update/' + id;
		const obj = {
				      		name: name,
      		capability: capability,
			ProductionLine: ProductionLine != null && ProductionLine.length > 0 ? ProductionLine : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a WorkCenter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteWorkCenter(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/WorkCenter/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a WorkCenter
	// returns the results untouched as an Observable WorkCenter
	// WorkCenter model
	// delegates via URI
	//********************************************************************
	getWorkCenter(id) : Observable<WorkCenter> {
		const uri_ = this.apiUrl + '/WorkCenter/load/' + id;

		return this.http.get<WorkCenter>(uri_);
	}
	
	//********************************************************************
	// gets all WorkCenter
	// returns the results untouched as JSON representation of an
	// Observable array of WorkCenter models
	// delegates via URI
	//********************************************************************
	getWorkCenters() : Observable<WorkCenter[]> {
		const uri_ = this.apiUrl + '/WorkCenter/';

		return this
			.http.get<WorkCenter[]>(uri_);
	}
	
			//********************************************************************
	// assigns a ProductionLine on a WorkCenter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignProductionLine( workCenterId, _productionLineId ): Observable<any> {

		// get the WorkCenter from storage
		this.loadHelper( workCenterId );

	// get the ProductionLine from storage
	var tmp 	= new ProductionLineService(this.http).getProductionLine(_productionLineId);

	// assign the ProductionLine
	this.workCenter.productionLine = tmp;

	// save the WorkCenter
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ProductionLine on a WorkCenter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignProductionLine( workCenterId ): Observable<any> {

		// get the WorkCenter from storage
		this.loadHelper( workCenterId );

	// assign ProductionLine to null
	this.workCenter.productionLine = null;

	// save the WorkCenter
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a WorkCenter
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/WorkCenter/update/' + this.workCenter;

	return  this.http.post(uri_, this.workCenter );
}

	//********************************************************************
	// loadHelper - internal helper to load a WorkCenter
	//********************************************************************	
	loadHelper( id ) {
		this.getWorkCenter(id)
			.subscribe((res : WorkCenter) => {
				this.workCenter = res;
			});
	}
}