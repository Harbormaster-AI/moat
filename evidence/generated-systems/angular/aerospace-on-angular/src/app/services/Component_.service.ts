import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Component_} from '../models/Component_';
import {SupplierService} from '../services/Supplier.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class Component_Service extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	component_ : Component_;

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
	// add a Component_
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addComponent_(partNumber, name, Supplier, ComponentCategory, SerializationMethod) : Observable<any> {
		const uri_ = this.apiUrl + '/Component_/create';
		const obj = {
			      		partNumber: partNumber,
      		name: name,
      		Supplier: Supplier != null && Supplier.length > 0 ? Supplier : null,
      		ComponentCategory: ComponentCategory,
			SerializationMethod: SerializationMethod
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Component_
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateComponent_(partNumber, name, Supplier, ComponentCategory, SerializationMethod, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Component_/update/' + id;
		const obj = {
				      		partNumber: partNumber,
      		name: name,
      		Supplier: Supplier != null && Supplier.length > 0 ? Supplier : null,
      		ComponentCategory: ComponentCategory,
			SerializationMethod: SerializationMethod
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Component_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteComponent_(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Component_/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Component_
	// returns the results untouched as an Observable Component_
	// Component_ model
	// delegates via URI
	//********************************************************************
	getComponent_(id) : Observable<Component_> {
		const uri_ = this.apiUrl + '/Component_/load/' + id;

		return this.http.get<Component_>(uri_);
	}
	
	//********************************************************************
	// gets all Component_
	// returns the results untouched as JSON representation of an
	// Observable array of Component_ models
	// delegates via URI
	//********************************************************************
	getComponent_s() : Observable<Component_[]> {
		const uri_ = this.apiUrl + '/Component_/';

		return this
			.http.get<Component_[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Supplier on a Component_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSupplier( component_Id, _supplierId ): Observable<any> {

		// get the Component_ from storage
		this.loadHelper( component_Id );

	// get the Supplier from storage
	var tmp 	= new SupplierService(this.http).getSupplier(_supplierId);

	// assign the Supplier
	this.component_.supplier = tmp;

	// save the Component_
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Supplier on a Component_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSupplier( component_Id ): Observable<any> {

		// get the Component_ from storage
		this.loadHelper( component_Id );

	// assign Supplier to null
	this.component_.supplier = null;

	// save the Component_
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Component_
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Component_/update/' + this.component_;

	return  this.http.post(uri_, this.component_ );
}

	//********************************************************************
	// loadHelper - internal helper to load a Component_
	//********************************************************************	
	loadHelper( id ) {
		this.getComponent_(id)
			.subscribe((res : Component_) => {
				this.component_ = res;
			});
	}
}