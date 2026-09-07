import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InsuranceProduct} from '../models/InsuranceProduct';
import {InsurerService} from '../services/Insurer.service';
import {CoverageDefinitionService} from '../services/CoverageDefinition.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InsuranceProductService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	insuranceProduct : InsuranceProduct;

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
	// add a InsuranceProduct
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInsuranceProduct(name, productCode, Insurer, CoverageDefinitions, LineOfBusiness) : Observable<any> {
		const uri_ = this.apiUrl + '/InsuranceProduct/create';
		const obj = {
			      		name: name,
      		productCode: productCode,
      		Insurer: Insurer != null && Insurer.length > 0 ? Insurer : null,
      		CoverageDefinitions: CoverageDefinitions != null && CoverageDefinitions.length > 0 ? CoverageDefinitions : null,
			LineOfBusiness: LineOfBusiness
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InsuranceProduct
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInsuranceProduct(name, productCode, Insurer, CoverageDefinitions, LineOfBusiness, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InsuranceProduct/update/' + id;
		const obj = {
				      		name: name,
      		productCode: productCode,
      		Insurer: Insurer != null && Insurer.length > 0 ? Insurer : null,
      		CoverageDefinitions: CoverageDefinitions != null && CoverageDefinitions.length > 0 ? CoverageDefinitions : null,
			LineOfBusiness: LineOfBusiness
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InsuranceProduct
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInsuranceProduct(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InsuranceProduct/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InsuranceProduct
	// returns the results untouched as an Observable InsuranceProduct
	// InsuranceProduct model
	// delegates via URI
	//********************************************************************
	getInsuranceProduct(id) : Observable<InsuranceProduct> {
		const uri_ = this.apiUrl + '/InsuranceProduct/load/' + id;

		return this.http.get<InsuranceProduct>(uri_);
	}
	
	//********************************************************************
	// gets all InsuranceProduct
	// returns the results untouched as JSON representation of an
	// Observable array of InsuranceProduct models
	// delegates via URI
	//********************************************************************
	getInsuranceProducts() : Observable<InsuranceProduct[]> {
		const uri_ = this.apiUrl + '/InsuranceProduct/';

		return this
			.http.get<InsuranceProduct[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Insurer on a InsuranceProduct
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInsurer( insuranceProductId, _insurerId ): Observable<any> {

		// get the InsuranceProduct from storage
		this.loadHelper( insuranceProductId );

	// get the Insurer from storage
	var tmp 	= new InsurerService(this.http).getInsurer(_insurerId);

	// assign the Insurer
	this.insuranceProduct.insurer = tmp;

	// save the InsuranceProduct
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Insurer on a InsuranceProduct
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInsurer( insuranceProductId ): Observable<any> {

		// get the InsuranceProduct from storage
		this.loadHelper( insuranceProductId );

	// assign Insurer to null
	this.insuranceProduct.insurer = null;

	// save the InsuranceProduct
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more coverageDefinitionsIds as a CoverageDefinitions
	// to a InsuranceProduct
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCoverageDefinitions( insuranceProductId, coverageDefinitionsIds ): Observable<any> {

		// get the InsuranceProduct
		this.loadHelper( insuranceProductId );

	// split on a comma with no spaces
	var idList = coverageDefinitionsIds.split(',')

	// iterate over array of coverageDefinitions ids
	idList.forEach(function (id) {
		// read the CoverageDefinition
		var coverageDefinition = new CoverageDefinitionService(this.http).getCoverageDefinition(id);
		// add the CoverageDefinition if not already assigned
		if ( this.insuranceProduct.coverageDefinitions.indexOf(coverageDefinition) == -1 )
		this.insuranceProduct.coverageDefinitions.push(coverageDefinition);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more coverageDefinitionsIds as a CoverageDefinitions
	// from a InsuranceProduct
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCoverageDefinitions( insuranceProductId, coverageDefinitionsIds ): Observable<any> {

		// get the InsuranceProduct
		this.loadHelper( insuranceProductId );


	// split on a comma with no spaces
	var idList 					= coverageDefinitionsIds.split(',');
	var coverageDefinitions 	= this.insuranceProduct.coverageDefinitions;

	if ( coverageDefinitions != null && coverageDefinitionsIds != null ) {

		// iterate over array of coverageDefinitions ids
		coverageDefinitions.forEach(function (obj) {
			if ( coverageDefinitionsIds.indexOf(obj._id) > -1 ) {
				// remove the CoverageDefinition
				this.insuranceProduct.coverageDefinitions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a InsuranceProduct
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InsuranceProduct/update/' + this.insuranceProduct;

	return  this.http.post(uri_, this.insuranceProduct );
}

	//********************************************************************
	// loadHelper - internal helper to load a InsuranceProduct
	//********************************************************************	
	loadHelper( id ) {
		this.getInsuranceProduct(id)
			.subscribe((res : InsuranceProduct) => {
				this.insuranceProduct = res;
			});
	}
}