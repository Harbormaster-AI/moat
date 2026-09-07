import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Authorization} from '../models/Authorization';
import {CoverageService} from '../services/Coverage.service';
import {ClinicalOrderService} from '../services/ClinicalOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AuthorizationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	authorization : Authorization;

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
	// add a Authorization
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAuthorization(authNumber, requestedService, Coverage, Order, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Authorization/create';
		const obj = {
			      		authNumber: authNumber,
      		requestedService: requestedService,
      		Coverage: Coverage != null && Coverage.length > 0 ? Coverage : null,
      		Order: Order != null && Order.length > 0 ? Order : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Authorization
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAuthorization(authNumber, requestedService, Coverage, Order, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Authorization/update/' + id;
		const obj = {
				      		authNumber: authNumber,
      		requestedService: requestedService,
      		Coverage: Coverage != null && Coverage.length > 0 ? Coverage : null,
      		Order: Order != null && Order.length > 0 ? Order : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Authorization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAuthorization(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Authorization/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Authorization
	// returns the results untouched as an Observable Authorization
	// Authorization model
	// delegates via URI
	//********************************************************************
	getAuthorization(id) : Observable<Authorization> {
		const uri_ = this.apiUrl + '/Authorization/load/' + id;

		return this.http.get<Authorization>(uri_);
	}
	
	//********************************************************************
	// gets all Authorization
	// returns the results untouched as JSON representation of an
	// Observable array of Authorization models
	// delegates via URI
	//********************************************************************
	getAuthorizations() : Observable<Authorization[]> {
		const uri_ = this.apiUrl + '/Authorization/';

		return this
			.http.get<Authorization[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Coverage on a Authorization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCoverage( authorizationId, _coverageId ): Observable<any> {

		// get the Authorization from storage
		this.loadHelper( authorizationId );

	// get the Coverage from storage
	var tmp 	= new CoverageService(this.http).getCoverage(_coverageId);

	// assign the Coverage
	this.authorization.coverage = tmp;

	// save the Authorization
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Coverage on a Authorization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCoverage( authorizationId ): Observable<any> {

		// get the Authorization from storage
		this.loadHelper( authorizationId );

	// assign Coverage to null
	this.authorization.coverage = null;

	// save the Authorization
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Order on a Authorization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrder( authorizationId, _orderId ): Observable<any> {

		// get the Authorization from storage
		this.loadHelper( authorizationId );

	// get the ClinicalOrder from storage
	var tmp 	= new ClinicalOrderService(this.http).getClinicalOrder(_orderId);

	// assign the Order
	this.authorization.order = tmp;

	// save the Authorization
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Order on a Authorization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrder( authorizationId ): Observable<any> {

		// get the Authorization from storage
		this.loadHelper( authorizationId );

	// assign Order to null
	this.authorization.order = null;

	// save the Authorization
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Authorization
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Authorization/update/' + this.authorization;

	return  this.http.post(uri_, this.authorization );
}

	//********************************************************************
	// loadHelper - internal helper to load a Authorization
	//********************************************************************	
	loadHelper( id ) {
		this.getAuthorization(id)
			.subscribe((res : Authorization) => {
				this.authorization = res;
			});
	}
}