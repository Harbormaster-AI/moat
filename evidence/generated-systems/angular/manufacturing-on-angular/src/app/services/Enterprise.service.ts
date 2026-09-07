import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Enterprise} from '../models/Enterprise';
import {BusinessUnitService} from '../services/BusinessUnit.service';
import {PlantService} from '../services/Plant.service';
import {SupplierService} from '../services/Supplier.service';
import {CustomerService} from '../services/Customer.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class EnterpriseService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	enterprise : Enterprise;

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
	// add a Enterprise
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addEnterprise(name, legalName, registrationCountry, website, taxId, BusinessUnits, Plants, Suppliers, Customers) : Observable<any> {
		const uri_ = this.apiUrl + '/Enterprise/create';
		const obj = {
			      		name: name,
      		legalName: legalName,
      		registrationCountry: registrationCountry,
      		website: website,
      		taxId: taxId,
      		BusinessUnits: BusinessUnits != null && BusinessUnits.length > 0 ? BusinessUnits : null,
      		Plants: Plants != null && Plants.length > 0 ? Plants : null,
      		Suppliers: Suppliers != null && Suppliers.length > 0 ? Suppliers : null,
			Customers: Customers != null && Customers.length > 0 ? Customers : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Enterprise
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateEnterprise(name, legalName, registrationCountry, website, taxId, BusinessUnits, Plants, Suppliers, Customers, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Enterprise/update/' + id;
		const obj = {
				      		name: name,
      		legalName: legalName,
      		registrationCountry: registrationCountry,
      		website: website,
      		taxId: taxId,
      		BusinessUnits: BusinessUnits != null && BusinessUnits.length > 0 ? BusinessUnits : null,
      		Plants: Plants != null && Plants.length > 0 ? Plants : null,
      		Suppliers: Suppliers != null && Suppliers.length > 0 ? Suppliers : null,
			Customers: Customers != null && Customers.length > 0 ? Customers : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Enterprise
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteEnterprise(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Enterprise/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Enterprise
	// returns the results untouched as an Observable Enterprise
	// Enterprise model
	// delegates via URI
	//********************************************************************
	getEnterprise(id) : Observable<Enterprise> {
		const uri_ = this.apiUrl + '/Enterprise/load/' + id;

		return this.http.get<Enterprise>(uri_);
	}
	
	//********************************************************************
	// gets all Enterprise
	// returns the results untouched as JSON representation of an
	// Observable array of Enterprise models
	// delegates via URI
	//********************************************************************
	getEnterprises() : Observable<Enterprise[]> {
		const uri_ = this.apiUrl + '/Enterprise/';

		return this
			.http.get<Enterprise[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more businessUnitsIds as a BusinessUnits
	// to a Enterprise
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addBusinessUnits( enterpriseId, businessUnitsIds ): Observable<any> {

		// get the Enterprise
		this.loadHelper( enterpriseId );

	// split on a comma with no spaces
	var idList = businessUnitsIds.split(',')

	// iterate over array of businessUnits ids
	idList.forEach(function (id) {
		// read the BusinessUnit
		var businessUnit = new BusinessUnitService(this.http).getBusinessUnit(id);
		// add the BusinessUnit if not already assigned
		if ( this.enterprise.businessUnits.indexOf(businessUnit) == -1 )
		this.enterprise.businessUnits.push(businessUnit);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more businessUnitsIds as a BusinessUnits
	// from a Enterprise
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeBusinessUnits( enterpriseId, businessUnitsIds ): Observable<any> {

		// get the Enterprise
		this.loadHelper( enterpriseId );


	// split on a comma with no spaces
	var idList 					= businessUnitsIds.split(',');
	var businessUnits 	= this.enterprise.businessUnits;

	if ( businessUnits != null && businessUnitsIds != null ) {

		// iterate over array of businessUnits ids
		businessUnits.forEach(function (obj) {
			if ( businessUnitsIds.indexOf(obj._id) > -1 ) {
				// remove the BusinessUnit
				this.enterprise.businessUnits.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more plantsIds as a Plants
	// to a Enterprise
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPlants( enterpriseId, plantsIds ): Observable<any> {

		// get the Enterprise
		this.loadHelper( enterpriseId );

	// split on a comma with no spaces
	var idList = plantsIds.split(',')

	// iterate over array of plants ids
	idList.forEach(function (id) {
		// read the Plant
		var plant = new PlantService(this.http).getPlant(id);
		// add the Plant if not already assigned
		if ( this.enterprise.plants.indexOf(plant) == -1 )
		this.enterprise.plants.push(plant);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more plantsIds as a Plants
	// from a Enterprise
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePlants( enterpriseId, plantsIds ): Observable<any> {

		// get the Enterprise
		this.loadHelper( enterpriseId );


	// split on a comma with no spaces
	var idList 					= plantsIds.split(',');
	var plants 	= this.enterprise.plants;

	if ( plants != null && plantsIds != null ) {

		// iterate over array of plants ids
		plants.forEach(function (obj) {
			if ( plantsIds.indexOf(obj._id) > -1 ) {
				// remove the Plant
				this.enterprise.plants.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more suppliersIds as a Suppliers
	// to a Enterprise
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSuppliers( enterpriseId, suppliersIds ): Observable<any> {

		// get the Enterprise
		this.loadHelper( enterpriseId );

	// split on a comma with no spaces
	var idList = suppliersIds.split(',')

	// iterate over array of suppliers ids
	idList.forEach(function (id) {
		// read the Supplier
		var supplier = new SupplierService(this.http).getSupplier(id);
		// add the Supplier if not already assigned
		if ( this.enterprise.suppliers.indexOf(supplier) == -1 )
		this.enterprise.suppliers.push(supplier);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more suppliersIds as a Suppliers
	// from a Enterprise
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSuppliers( enterpriseId, suppliersIds ): Observable<any> {

		// get the Enterprise
		this.loadHelper( enterpriseId );


	// split on a comma with no spaces
	var idList 					= suppliersIds.split(',');
	var suppliers 	= this.enterprise.suppliers;

	if ( suppliers != null && suppliersIds != null ) {

		// iterate over array of suppliers ids
		suppliers.forEach(function (obj) {
			if ( suppliersIds.indexOf(obj._id) > -1 ) {
				// remove the Supplier
				this.enterprise.suppliers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more customersIds as a Customers
	// to a Enterprise
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCustomers( enterpriseId, customersIds ): Observable<any> {

		// get the Enterprise
		this.loadHelper( enterpriseId );

	// split on a comma with no spaces
	var idList = customersIds.split(',')

	// iterate over array of customers ids
	idList.forEach(function (id) {
		// read the Customer
		var customer = new CustomerService(this.http).getCustomer(id);
		// add the Customer if not already assigned
		if ( this.enterprise.customers.indexOf(customer) == -1 )
		this.enterprise.customers.push(customer);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more customersIds as a Customers
	// from a Enterprise
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCustomers( enterpriseId, customersIds ): Observable<any> {

		// get the Enterprise
		this.loadHelper( enterpriseId );


	// split on a comma with no spaces
	var idList 					= customersIds.split(',');
	var customers 	= this.enterprise.customers;

	if ( customers != null && customersIds != null ) {

		// iterate over array of customers ids
		customers.forEach(function (obj) {
			if ( customersIds.indexOf(obj._id) > -1 ) {
				// remove the Customer
				this.enterprise.customers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Enterprise
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Enterprise/update/' + this.enterprise;

	return  this.http.post(uri_, this.enterprise );
}

	//********************************************************************
	// loadHelper - internal helper to load a Enterprise
	//********************************************************************	
	loadHelper( id ) {
		this.getEnterprise(id)
			.subscribe((res : Enterprise) => {
				this.enterprise = res;
			});
	}
}