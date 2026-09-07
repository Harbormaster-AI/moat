import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {SalaryComponent} from '../models/SalaryComponent';
import {CompensationPackageService} from '../services/CompensationPackage.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class SalaryComponentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	salaryComponent : SalaryComponent;

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
	// add a SalaryComponent
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addSalaryComponent(amount, recurring, CompensationPackage, ComponentType) : Observable<any> {
		const uri_ = this.apiUrl + '/SalaryComponent/create';
		const obj = {
			      		amount: amount,
      		recurring: recurring,
      		CompensationPackage: CompensationPackage != null && CompensationPackage.length > 0 ? CompensationPackage : null,
			ComponentType: ComponentType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a SalaryComponent
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateSalaryComponent(amount, recurring, CompensationPackage, ComponentType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/SalaryComponent/update/' + id;
		const obj = {
				      		amount: amount,
      		recurring: recurring,
      		CompensationPackage: CompensationPackage != null && CompensationPackage.length > 0 ? CompensationPackage : null,
			ComponentType: ComponentType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a SalaryComponent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteSalaryComponent(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/SalaryComponent/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a SalaryComponent
	// returns the results untouched as an Observable SalaryComponent
	// SalaryComponent model
	// delegates via URI
	//********************************************************************
	getSalaryComponent(id) : Observable<SalaryComponent> {
		const uri_ = this.apiUrl + '/SalaryComponent/load/' + id;

		return this.http.get<SalaryComponent>(uri_);
	}
	
	//********************************************************************
	// gets all SalaryComponent
	// returns the results untouched as JSON representation of an
	// Observable array of SalaryComponent models
	// delegates via URI
	//********************************************************************
	getSalaryComponents() : Observable<SalaryComponent[]> {
		const uri_ = this.apiUrl + '/SalaryComponent/';

		return this
			.http.get<SalaryComponent[]>(uri_);
	}
	
			//********************************************************************
	// assigns a CompensationPackage on a SalaryComponent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCompensationPackage( salaryComponentId, _compensationPackageId ): Observable<any> {

		// get the SalaryComponent from storage
		this.loadHelper( salaryComponentId );

	// get the CompensationPackage from storage
	var tmp 	= new CompensationPackageService(this.http).getCompensationPackage(_compensationPackageId);

	// assign the CompensationPackage
	this.salaryComponent.compensationPackage = tmp;

	// save the SalaryComponent
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CompensationPackage on a SalaryComponent
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCompensationPackage( salaryComponentId ): Observable<any> {

		// get the SalaryComponent from storage
		this.loadHelper( salaryComponentId );

	// assign CompensationPackage to null
	this.salaryComponent.compensationPackage = null;

	// save the SalaryComponent
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a SalaryComponent
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/SalaryComponent/update/' + this.salaryComponent;

	return  this.http.post(uri_, this.salaryComponent );
}

	//********************************************************************
	// loadHelper - internal helper to load a SalaryComponent
	//********************************************************************	
	loadHelper( id ) {
		this.getSalaryComponent(id)
			.subscribe((res : SalaryComponent) => {
				this.salaryComponent = res;
			});
	}
}