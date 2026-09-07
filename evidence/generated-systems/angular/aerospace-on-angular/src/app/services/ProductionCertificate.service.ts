import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ProductionCertificate} from '../models/ProductionCertificate';
import {AerospaceManufacturerService} from '../services/AerospaceManufacturer.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ProductionCertificateService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	productionCertificate : ProductionCertificate;

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
	// add a ProductionCertificate
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addProductionCertificate(certificateNumber, authority, Manufacturer) : Observable<any> {
		const uri_ = this.apiUrl + '/ProductionCertificate/create';
		const obj = {
			      		certificateNumber: certificateNumber,
      		authority: authority,
			Manufacturer: Manufacturer != null && Manufacturer.length > 0 ? Manufacturer : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ProductionCertificate
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateProductionCertificate(certificateNumber, authority, Manufacturer, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ProductionCertificate/update/' + id;
		const obj = {
				      		certificateNumber: certificateNumber,
      		authority: authority,
			Manufacturer: Manufacturer != null && Manufacturer.length > 0 ? Manufacturer : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ProductionCertificate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteProductionCertificate(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ProductionCertificate/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ProductionCertificate
	// returns the results untouched as an Observable ProductionCertificate
	// ProductionCertificate model
	// delegates via URI
	//********************************************************************
	getProductionCertificate(id) : Observable<ProductionCertificate> {
		const uri_ = this.apiUrl + '/ProductionCertificate/load/' + id;

		return this.http.get<ProductionCertificate>(uri_);
	}
	
	//********************************************************************
	// gets all ProductionCertificate
	// returns the results untouched as JSON representation of an
	// Observable array of ProductionCertificate models
	// delegates via URI
	//********************************************************************
	getProductionCertificates() : Observable<ProductionCertificate[]> {
		const uri_ = this.apiUrl + '/ProductionCertificate/';

		return this
			.http.get<ProductionCertificate[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Manufacturer on a ProductionCertificate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignManufacturer( productionCertificateId, _manufacturerId ): Observable<any> {

		// get the ProductionCertificate from storage
		this.loadHelper( productionCertificateId );

	// get the AerospaceManufacturer from storage
	var tmp 	= new AerospaceManufacturerService(this.http).getAerospaceManufacturer(_manufacturerId);

	// assign the Manufacturer
	this.productionCertificate.manufacturer = tmp;

	// save the ProductionCertificate
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Manufacturer on a ProductionCertificate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignManufacturer( productionCertificateId ): Observable<any> {

		// get the ProductionCertificate from storage
		this.loadHelper( productionCertificateId );

	// assign Manufacturer to null
	this.productionCertificate.manufacturer = null;

	// save the ProductionCertificate
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a ProductionCertificate
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ProductionCertificate/update/' + this.productionCertificate;

	return  this.http.post(uri_, this.productionCertificate );
}

	//********************************************************************
	// loadHelper - internal helper to load a ProductionCertificate
	//********************************************************************	
	loadHelper( id ) {
		this.getProductionCertificate(id)
			.subscribe((res : ProductionCertificate) => {
				this.productionCertificate = res;
			});
	}
}