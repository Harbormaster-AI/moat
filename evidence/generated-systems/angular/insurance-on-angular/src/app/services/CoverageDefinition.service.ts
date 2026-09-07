import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CoverageDefinition} from '../models/CoverageDefinition';
import {InsuranceProductService} from '../services/InsuranceProduct.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CoverageDefinitionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	coverageDefinition : CoverageDefinition;

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
	// add a CoverageDefinition
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCoverageDefinition(name, defaultLimit, defaultDeductible, asMandatory, Product, CoverageType) : Observable<any> {
		const uri_ = this.apiUrl + '/CoverageDefinition/create';
		const obj = {
			      		name: name,
      		defaultLimit: defaultLimit,
      		defaultDeductible: defaultDeductible,
      		asMandatory: asMandatory,
      		Product: Product != null && Product.length > 0 ? Product : null,
			CoverageType: CoverageType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CoverageDefinition
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCoverageDefinition(name, defaultLimit, defaultDeductible, asMandatory, Product, CoverageType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CoverageDefinition/update/' + id;
		const obj = {
				      		name: name,
      		defaultLimit: defaultLimit,
      		defaultDeductible: defaultDeductible,
      		asMandatory: asMandatory,
      		Product: Product != null && Product.length > 0 ? Product : null,
			CoverageType: CoverageType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CoverageDefinition
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCoverageDefinition(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CoverageDefinition/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CoverageDefinition
	// returns the results untouched as an Observable CoverageDefinition
	// CoverageDefinition model
	// delegates via URI
	//********************************************************************
	getCoverageDefinition(id) : Observable<CoverageDefinition> {
		const uri_ = this.apiUrl + '/CoverageDefinition/load/' + id;

		return this.http.get<CoverageDefinition>(uri_);
	}
	
	//********************************************************************
	// gets all CoverageDefinition
	// returns the results untouched as JSON representation of an
	// Observable array of CoverageDefinition models
	// delegates via URI
	//********************************************************************
	getCoverageDefinitions() : Observable<CoverageDefinition[]> {
		const uri_ = this.apiUrl + '/CoverageDefinition/';

		return this
			.http.get<CoverageDefinition[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Product on a CoverageDefinition
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignProduct( coverageDefinitionId, _productId ): Observable<any> {

		// get the CoverageDefinition from storage
		this.loadHelper( coverageDefinitionId );

	// get the InsuranceProduct from storage
	var tmp 	= new InsuranceProductService(this.http).getInsuranceProduct(_productId);

	// assign the Product
	this.coverageDefinition.product = tmp;

	// save the CoverageDefinition
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Product on a CoverageDefinition
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignProduct( coverageDefinitionId ): Observable<any> {

		// get the CoverageDefinition from storage
		this.loadHelper( coverageDefinitionId );

	// assign Product to null
	this.coverageDefinition.product = null;

	// save the CoverageDefinition
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a CoverageDefinition
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CoverageDefinition/update/' + this.coverageDefinition;

	return  this.http.post(uri_, this.coverageDefinition );
}

	//********************************************************************
	// loadHelper - internal helper to load a CoverageDefinition
	//********************************************************************	
	loadHelper( id ) {
		this.getCoverageDefinition(id)
			.subscribe((res : CoverageDefinition) => {
				this.coverageDefinition = res;
			});
	}
}