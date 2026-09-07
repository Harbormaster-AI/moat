import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AerospaceManufacturer} from '../models/AerospaceManufacturer';
import {AircraftProgramService} from '../services/AircraftProgram.service';
import {PlantService} from '../services/Plant.service';
import {SupplierService} from '../services/Supplier.service';
import {ProductionCertificateService} from '../services/ProductionCertificate.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AerospaceManufacturerService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	aerospaceManufacturer : AerospaceManufacturer;

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
	// add a AerospaceManufacturer
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAerospaceManufacturer(name, legalName, headquartersCountry, website, Programs, Plants, Suppliers, ProductionCertificates) : Observable<any> {
		const uri_ = this.apiUrl + '/AerospaceManufacturer/create';
		const obj = {
			      		name: name,
      		legalName: legalName,
      		headquartersCountry: headquartersCountry,
      		website: website,
      		Programs: Programs != null && Programs.length > 0 ? Programs : null,
      		Plants: Plants != null && Plants.length > 0 ? Plants : null,
      		Suppliers: Suppliers != null && Suppliers.length > 0 ? Suppliers : null,
			ProductionCertificates: ProductionCertificates != null && ProductionCertificates.length > 0 ? ProductionCertificates : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AerospaceManufacturer
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAerospaceManufacturer(name, legalName, headquartersCountry, website, Programs, Plants, Suppliers, ProductionCertificates, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AerospaceManufacturer/update/' + id;
		const obj = {
				      		name: name,
      		legalName: legalName,
      		headquartersCountry: headquartersCountry,
      		website: website,
      		Programs: Programs != null && Programs.length > 0 ? Programs : null,
      		Plants: Plants != null && Plants.length > 0 ? Plants : null,
      		Suppliers: Suppliers != null && Suppliers.length > 0 ? Suppliers : null,
			ProductionCertificates: ProductionCertificates != null && ProductionCertificates.length > 0 ? ProductionCertificates : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AerospaceManufacturer
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAerospaceManufacturer(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AerospaceManufacturer/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AerospaceManufacturer
	// returns the results untouched as an Observable AerospaceManufacturer
	// AerospaceManufacturer model
	// delegates via URI
	//********************************************************************
	getAerospaceManufacturer(id) : Observable<AerospaceManufacturer> {
		const uri_ = this.apiUrl + '/AerospaceManufacturer/load/' + id;

		return this.http.get<AerospaceManufacturer>(uri_);
	}
	
	//********************************************************************
	// gets all AerospaceManufacturer
	// returns the results untouched as JSON representation of an
	// Observable array of AerospaceManufacturer models
	// delegates via URI
	//********************************************************************
	getAerospaceManufacturers() : Observable<AerospaceManufacturer[]> {
		const uri_ = this.apiUrl + '/AerospaceManufacturer/';

		return this
			.http.get<AerospaceManufacturer[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more programsIds as a Programs
	// to a AerospaceManufacturer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPrograms( aerospaceManufacturerId, programsIds ): Observable<any> {

		// get the AerospaceManufacturer
		this.loadHelper( aerospaceManufacturerId );

	// split on a comma with no spaces
	var idList = programsIds.split(',')

	// iterate over array of programs ids
	idList.forEach(function (id) {
		// read the AircraftProgram
		var aircraftProgram = new AircraftProgramService(this.http).getAircraftProgram(id);
		// add the AircraftProgram if not already assigned
		if ( this.aerospaceManufacturer.programs.indexOf(aircraftProgram) == -1 )
		this.aerospaceManufacturer.programs.push(aircraftProgram);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more programsIds as a Programs
	// from a AerospaceManufacturer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePrograms( aerospaceManufacturerId, programsIds ): Observable<any> {

		// get the AerospaceManufacturer
		this.loadHelper( aerospaceManufacturerId );


	// split on a comma with no spaces
	var idList 					= programsIds.split(',');
	var programs 	= this.aerospaceManufacturer.programs;

	if ( programs != null && programsIds != null ) {

		// iterate over array of programs ids
		programs.forEach(function (obj) {
			if ( programsIds.indexOf(obj._id) > -1 ) {
				// remove the AircraftProgram
				this.aerospaceManufacturer.programs.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more plantsIds as a Plants
	// to a AerospaceManufacturer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPlants( aerospaceManufacturerId, plantsIds ): Observable<any> {

		// get the AerospaceManufacturer
		this.loadHelper( aerospaceManufacturerId );

	// split on a comma with no spaces
	var idList = plantsIds.split(',')

	// iterate over array of plants ids
	idList.forEach(function (id) {
		// read the Plant
		var plant = new PlantService(this.http).getPlant(id);
		// add the Plant if not already assigned
		if ( this.aerospaceManufacturer.plants.indexOf(plant) == -1 )
		this.aerospaceManufacturer.plants.push(plant);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more plantsIds as a Plants
	// from a AerospaceManufacturer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePlants( aerospaceManufacturerId, plantsIds ): Observable<any> {

		// get the AerospaceManufacturer
		this.loadHelper( aerospaceManufacturerId );


	// split on a comma with no spaces
	var idList 					= plantsIds.split(',');
	var plants 	= this.aerospaceManufacturer.plants;

	if ( plants != null && plantsIds != null ) {

		// iterate over array of plants ids
		plants.forEach(function (obj) {
			if ( plantsIds.indexOf(obj._id) > -1 ) {
				// remove the Plant
				this.aerospaceManufacturer.plants.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more suppliersIds as a Suppliers
	// to a AerospaceManufacturer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSuppliers( aerospaceManufacturerId, suppliersIds ): Observable<any> {

		// get the AerospaceManufacturer
		this.loadHelper( aerospaceManufacturerId );

	// split on a comma with no spaces
	var idList = suppliersIds.split(',')

	// iterate over array of suppliers ids
	idList.forEach(function (id) {
		// read the Supplier
		var supplier = new SupplierService(this.http).getSupplier(id);
		// add the Supplier if not already assigned
		if ( this.aerospaceManufacturer.suppliers.indexOf(supplier) == -1 )
		this.aerospaceManufacturer.suppliers.push(supplier);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more suppliersIds as a Suppliers
	// from a AerospaceManufacturer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSuppliers( aerospaceManufacturerId, suppliersIds ): Observable<any> {

		// get the AerospaceManufacturer
		this.loadHelper( aerospaceManufacturerId );


	// split on a comma with no spaces
	var idList 					= suppliersIds.split(',');
	var suppliers 	= this.aerospaceManufacturer.suppliers;

	if ( suppliers != null && suppliersIds != null ) {

		// iterate over array of suppliers ids
		suppliers.forEach(function (obj) {
			if ( suppliersIds.indexOf(obj._id) > -1 ) {
				// remove the Supplier
				this.aerospaceManufacturer.suppliers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more productionCertificatesIds as a ProductionCertificates
	// to a AerospaceManufacturer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProductionCertificates( aerospaceManufacturerId, productionCertificatesIds ): Observable<any> {

		// get the AerospaceManufacturer
		this.loadHelper( aerospaceManufacturerId );

	// split on a comma with no spaces
	var idList = productionCertificatesIds.split(',')

	// iterate over array of productionCertificates ids
	idList.forEach(function (id) {
		// read the ProductionCertificate
		var productionCertificate = new ProductionCertificateService(this.http).getProductionCertificate(id);
		// add the ProductionCertificate if not already assigned
		if ( this.aerospaceManufacturer.productionCertificates.indexOf(productionCertificate) == -1 )
		this.aerospaceManufacturer.productionCertificates.push(productionCertificate);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more productionCertificatesIds as a ProductionCertificates
	// from a AerospaceManufacturer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProductionCertificates( aerospaceManufacturerId, productionCertificatesIds ): Observable<any> {

		// get the AerospaceManufacturer
		this.loadHelper( aerospaceManufacturerId );


	// split on a comma with no spaces
	var idList 					= productionCertificatesIds.split(',');
	var productionCertificates 	= this.aerospaceManufacturer.productionCertificates;

	if ( productionCertificates != null && productionCertificatesIds != null ) {

		// iterate over array of productionCertificates ids
		productionCertificates.forEach(function (obj) {
			if ( productionCertificatesIds.indexOf(obj._id) > -1 ) {
				// remove the ProductionCertificate
				this.aerospaceManufacturer.productionCertificates.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a AerospaceManufacturer
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AerospaceManufacturer/update/' + this.aerospaceManufacturer;

	return  this.http.post(uri_, this.aerospaceManufacturer );
}

	//********************************************************************
	// loadHelper - internal helper to load a AerospaceManufacturer
	//********************************************************************	
	loadHelper( id ) {
		this.getAerospaceManufacturer(id)
			.subscribe((res : AerospaceManufacturer) => {
				this.aerospaceManufacturer = res;
			});
	}
}