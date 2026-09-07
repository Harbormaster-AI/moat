import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {APU} from '../models/APU';
import {SupplierService} from '../services/Supplier.service';
import {AircraftVariantService} from '../services/AircraftVariant.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class APUService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	aPU : APU;

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
	// add a APU
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAPU(model_, Supplier, Variants) : Observable<any> {
		const uri_ = this.apiUrl + '/APU/create';
		const obj = {
			      		model_: model_,
      		Supplier: Supplier != null && Supplier.length > 0 ? Supplier : null,
			Variants: Variants != null && Variants.length > 0 ? Variants : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a APU
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAPU(model_, Supplier, Variants, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/APU/update/' + id;
		const obj = {
				      		model_: model_,
      		Supplier: Supplier != null && Supplier.length > 0 ? Supplier : null,
			Variants: Variants != null && Variants.length > 0 ? Variants : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a APU
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAPU(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/APU/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a APU
	// returns the results untouched as an Observable APU
	// APU model
	// delegates via URI
	//********************************************************************
	getAPU(id) : Observable<APU> {
		const uri_ = this.apiUrl + '/APU/load/' + id;

		return this.http.get<APU>(uri_);
	}
	
	//********************************************************************
	// gets all APU
	// returns the results untouched as JSON representation of an
	// Observable array of APU models
	// delegates via URI
	//********************************************************************
	getAPUs() : Observable<APU[]> {
		const uri_ = this.apiUrl + '/APU/';

		return this
			.http.get<APU[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Supplier on a APU
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSupplier( aPUId, _supplierId ): Observable<any> {

		// get the APU from storage
		this.loadHelper( aPUId );

	// get the Supplier from storage
	var tmp 	= new SupplierService(this.http).getSupplier(_supplierId);

	// assign the Supplier
	this.aPU.supplier = tmp;

	// save the APU
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Supplier on a APU
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSupplier( aPUId ): Observable<any> {

		// get the APU from storage
		this.loadHelper( aPUId );

	// assign Supplier to null
	this.aPU.supplier = null;

	// save the APU
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more variantsIds as a Variants
	// to a APU
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addVariants( aPUId, variantsIds ): Observable<any> {

		// get the APU
		this.loadHelper( aPUId );

	// split on a comma with no spaces
	var idList = variantsIds.split(',')

	// iterate over array of variants ids
	idList.forEach(function (id) {
		// read the AircraftVariant
		var aircraftVariant = new AircraftVariantService(this.http).getAircraftVariant(id);
		// add the AircraftVariant if not already assigned
		if ( this.aPU.variants.indexOf(aircraftVariant) == -1 )
		this.aPU.variants.push(aircraftVariant);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more variantsIds as a Variants
	// from a APU
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeVariants( aPUId, variantsIds ): Observable<any> {

		// get the APU
		this.loadHelper( aPUId );


	// split on a comma with no spaces
	var idList 					= variantsIds.split(',');
	var variants 	= this.aPU.variants;

	if ( variants != null && variantsIds != null ) {

		// iterate over array of variants ids
		variants.forEach(function (obj) {
			if ( variantsIds.indexOf(obj._id) > -1 ) {
				// remove the AircraftVariant
				this.aPU.variants.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a APU
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/APU/update/' + this.aPU;

	return  this.http.post(uri_, this.aPU );
}

	//********************************************************************
	// loadHelper - internal helper to load a APU
	//********************************************************************	
	loadHelper( id ) {
		this.getAPU(id)
			.subscribe((res : APU) => {
				this.aPU = res;
			});
	}
}