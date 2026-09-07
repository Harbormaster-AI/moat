import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AvionicsSuite} from '../models/AvionicsSuite';
import {SupplierService} from '../services/Supplier.service';
import {AircraftVariantService} from '../services/AircraftVariant.service';
import {SoftwareLoadService} from '../services/SoftwareLoad.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AvionicsSuiteService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	avionicsSuite : AvionicsSuite;

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
	// add a AvionicsSuite
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAvionicsSuite(suiteName, softwareBaseline, Supplier, Variants, SoftwareLoads) : Observable<any> {
		const uri_ = this.apiUrl + '/AvionicsSuite/create';
		const obj = {
			      		suiteName: suiteName,
      		softwareBaseline: softwareBaseline,
      		Supplier: Supplier != null && Supplier.length > 0 ? Supplier : null,
      		Variants: Variants != null && Variants.length > 0 ? Variants : null,
			SoftwareLoads: SoftwareLoads != null && SoftwareLoads.length > 0 ? SoftwareLoads : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AvionicsSuite
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAvionicsSuite(suiteName, softwareBaseline, Supplier, Variants, SoftwareLoads, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AvionicsSuite/update/' + id;
		const obj = {
				      		suiteName: suiteName,
      		softwareBaseline: softwareBaseline,
      		Supplier: Supplier != null && Supplier.length > 0 ? Supplier : null,
      		Variants: Variants != null && Variants.length > 0 ? Variants : null,
			SoftwareLoads: SoftwareLoads != null && SoftwareLoads.length > 0 ? SoftwareLoads : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AvionicsSuite
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAvionicsSuite(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AvionicsSuite/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AvionicsSuite
	// returns the results untouched as an Observable AvionicsSuite
	// AvionicsSuite model
	// delegates via URI
	//********************************************************************
	getAvionicsSuite(id) : Observable<AvionicsSuite> {
		const uri_ = this.apiUrl + '/AvionicsSuite/load/' + id;

		return this.http.get<AvionicsSuite>(uri_);
	}
	
	//********************************************************************
	// gets all AvionicsSuite
	// returns the results untouched as JSON representation of an
	// Observable array of AvionicsSuite models
	// delegates via URI
	//********************************************************************
	getAvionicsSuites() : Observable<AvionicsSuite[]> {
		const uri_ = this.apiUrl + '/AvionicsSuite/';

		return this
			.http.get<AvionicsSuite[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Supplier on a AvionicsSuite
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSupplier( avionicsSuiteId, _supplierId ): Observable<any> {

		// get the AvionicsSuite from storage
		this.loadHelper( avionicsSuiteId );

	// get the Supplier from storage
	var tmp 	= new SupplierService(this.http).getSupplier(_supplierId);

	// assign the Supplier
	this.avionicsSuite.supplier = tmp;

	// save the AvionicsSuite
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Supplier on a AvionicsSuite
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSupplier( avionicsSuiteId ): Observable<any> {

		// get the AvionicsSuite from storage
		this.loadHelper( avionicsSuiteId );

	// assign Supplier to null
	this.avionicsSuite.supplier = null;

	// save the AvionicsSuite
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more variantsIds as a Variants
	// to a AvionicsSuite
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addVariants( avionicsSuiteId, variantsIds ): Observable<any> {

		// get the AvionicsSuite
		this.loadHelper( avionicsSuiteId );

	// split on a comma with no spaces
	var idList = variantsIds.split(',')

	// iterate over array of variants ids
	idList.forEach(function (id) {
		// read the AircraftVariant
		var aircraftVariant = new AircraftVariantService(this.http).getAircraftVariant(id);
		// add the AircraftVariant if not already assigned
		if ( this.avionicsSuite.variants.indexOf(aircraftVariant) == -1 )
		this.avionicsSuite.variants.push(aircraftVariant);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more variantsIds as a Variants
	// from a AvionicsSuite
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeVariants( avionicsSuiteId, variantsIds ): Observable<any> {

		// get the AvionicsSuite
		this.loadHelper( avionicsSuiteId );


	// split on a comma with no spaces
	var idList 					= variantsIds.split(',');
	var variants 	= this.avionicsSuite.variants;

	if ( variants != null && variantsIds != null ) {

		// iterate over array of variants ids
		variants.forEach(function (obj) {
			if ( variantsIds.indexOf(obj._id) > -1 ) {
				// remove the AircraftVariant
				this.avionicsSuite.variants.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more softwareLoadsIds as a SoftwareLoads
	// to a AvionicsSuite
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSoftwareLoads( avionicsSuiteId, softwareLoadsIds ): Observable<any> {

		// get the AvionicsSuite
		this.loadHelper( avionicsSuiteId );

	// split on a comma with no spaces
	var idList = softwareLoadsIds.split(',')

	// iterate over array of softwareLoads ids
	idList.forEach(function (id) {
		// read the SoftwareLoad
		var softwareLoad = new SoftwareLoadService(this.http).getSoftwareLoad(id);
		// add the SoftwareLoad if not already assigned
		if ( this.avionicsSuite.softwareLoads.indexOf(softwareLoad) == -1 )
		this.avionicsSuite.softwareLoads.push(softwareLoad);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more softwareLoadsIds as a SoftwareLoads
	// from a AvionicsSuite
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSoftwareLoads( avionicsSuiteId, softwareLoadsIds ): Observable<any> {

		// get the AvionicsSuite
		this.loadHelper( avionicsSuiteId );


	// split on a comma with no spaces
	var idList 					= softwareLoadsIds.split(',');
	var softwareLoads 	= this.avionicsSuite.softwareLoads;

	if ( softwareLoads != null && softwareLoadsIds != null ) {

		// iterate over array of softwareLoads ids
		softwareLoads.forEach(function (obj) {
			if ( softwareLoadsIds.indexOf(obj._id) > -1 ) {
				// remove the SoftwareLoad
				this.avionicsSuite.softwareLoads.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a AvionicsSuite
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AvionicsSuite/update/' + this.avionicsSuite;

	return  this.http.post(uri_, this.avionicsSuite );
}

	//********************************************************************
	// loadHelper - internal helper to load a AvionicsSuite
	//********************************************************************	
	loadHelper( id ) {
		this.getAvionicsSuite(id)
			.subscribe((res : AvionicsSuite) => {
				this.avionicsSuite = res;
			});
	}
}