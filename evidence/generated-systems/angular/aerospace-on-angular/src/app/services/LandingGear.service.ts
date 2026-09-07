import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {LandingGear} from '../models/LandingGear';
import {SupplierService} from '../services/Supplier.service';
import {AircraftVariantService} from '../services/AircraftVariant.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class LandingGearService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	landingGear : LandingGear;

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
	// add a LandingGear
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addLandingGear(supplierPartNumber, Supplier, Variants, GearType) : Observable<any> {
		const uri_ = this.apiUrl + '/LandingGear/create';
		const obj = {
			      		supplierPartNumber: supplierPartNumber,
      		Supplier: Supplier != null && Supplier.length > 0 ? Supplier : null,
      		Variants: Variants != null && Variants.length > 0 ? Variants : null,
			GearType: GearType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a LandingGear
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateLandingGear(supplierPartNumber, Supplier, Variants, GearType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/LandingGear/update/' + id;
		const obj = {
				      		supplierPartNumber: supplierPartNumber,
      		Supplier: Supplier != null && Supplier.length > 0 ? Supplier : null,
      		Variants: Variants != null && Variants.length > 0 ? Variants : null,
			GearType: GearType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a LandingGear
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteLandingGear(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/LandingGear/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a LandingGear
	// returns the results untouched as an Observable LandingGear
	// LandingGear model
	// delegates via URI
	//********************************************************************
	getLandingGear(id) : Observable<LandingGear> {
		const uri_ = this.apiUrl + '/LandingGear/load/' + id;

		return this.http.get<LandingGear>(uri_);
	}
	
	//********************************************************************
	// gets all LandingGear
	// returns the results untouched as JSON representation of an
	// Observable array of LandingGear models
	// delegates via URI
	//********************************************************************
	getLandingGears() : Observable<LandingGear[]> {
		const uri_ = this.apiUrl + '/LandingGear/';

		return this
			.http.get<LandingGear[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Supplier on a LandingGear
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSupplier( landingGearId, _supplierId ): Observable<any> {

		// get the LandingGear from storage
		this.loadHelper( landingGearId );

	// get the Supplier from storage
	var tmp 	= new SupplierService(this.http).getSupplier(_supplierId);

	// assign the Supplier
	this.landingGear.supplier = tmp;

	// save the LandingGear
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Supplier on a LandingGear
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSupplier( landingGearId ): Observable<any> {

		// get the LandingGear from storage
		this.loadHelper( landingGearId );

	// assign Supplier to null
	this.landingGear.supplier = null;

	// save the LandingGear
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more variantsIds as a Variants
	// to a LandingGear
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addVariants( landingGearId, variantsIds ): Observable<any> {

		// get the LandingGear
		this.loadHelper( landingGearId );

	// split on a comma with no spaces
	var idList = variantsIds.split(',')

	// iterate over array of variants ids
	idList.forEach(function (id) {
		// read the AircraftVariant
		var aircraftVariant = new AircraftVariantService(this.http).getAircraftVariant(id);
		// add the AircraftVariant if not already assigned
		if ( this.landingGear.variants.indexOf(aircraftVariant) == -1 )
		this.landingGear.variants.push(aircraftVariant);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more variantsIds as a Variants
	// from a LandingGear
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeVariants( landingGearId, variantsIds ): Observable<any> {

		// get the LandingGear
		this.loadHelper( landingGearId );


	// split on a comma with no spaces
	var idList 					= variantsIds.split(',');
	var variants 	= this.landingGear.variants;

	if ( variants != null && variantsIds != null ) {

		// iterate over array of variants ids
		variants.forEach(function (obj) {
			if ( variantsIds.indexOf(obj._id) > -1 ) {
				// remove the AircraftVariant
				this.landingGear.variants.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a LandingGear
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/LandingGear/update/' + this.landingGear;

	return  this.http.post(uri_, this.landingGear );
}

	//********************************************************************
	// loadHelper - internal helper to load a LandingGear
	//********************************************************************	
	loadHelper( id ) {
		this.getLandingGear(id)
			.subscribe((res : LandingGear) => {
				this.landingGear = res;
			});
	}
}