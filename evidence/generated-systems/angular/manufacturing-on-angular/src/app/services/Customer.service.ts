import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Customer} from '../models/Customer';
import {EnterpriseService} from '../services/Enterprise.service';
import {SalesOrderService} from '../services/SalesOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CustomerService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	customer : Customer;

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
	// add a Customer
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCustomer(name, customerCode, address, Enterprises, SalesOrders, CustomerType) : Observable<any> {
		const uri_ = this.apiUrl + '/Customer/create';
		const obj = {
			      		name: name,
      		customerCode: customerCode,
      		address: address,
      		Enterprises: Enterprises != null && Enterprises.length > 0 ? Enterprises : null,
      		SalesOrders: SalesOrders != null && SalesOrders.length > 0 ? SalesOrders : null,
			CustomerType: CustomerType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Customer
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCustomer(name, customerCode, address, Enterprises, SalesOrders, CustomerType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Customer/update/' + id;
		const obj = {
				      		name: name,
      		customerCode: customerCode,
      		address: address,
      		Enterprises: Enterprises != null && Enterprises.length > 0 ? Enterprises : null,
      		SalesOrders: SalesOrders != null && SalesOrders.length > 0 ? SalesOrders : null,
			CustomerType: CustomerType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Customer
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCustomer(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Customer/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Customer
	// returns the results untouched as an Observable Customer
	// Customer model
	// delegates via URI
	//********************************************************************
	getCustomer(id) : Observable<Customer> {
		const uri_ = this.apiUrl + '/Customer/load/' + id;

		return this.http.get<Customer>(uri_);
	}
	
	//********************************************************************
	// gets all Customer
	// returns the results untouched as JSON representation of an
	// Observable array of Customer models
	// delegates via URI
	//********************************************************************
	getCustomers() : Observable<Customer[]> {
		const uri_ = this.apiUrl + '/Customer/';

		return this
			.http.get<Customer[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more enterprisesIds as a Enterprises
	// to a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEnterprises( customerId, enterprisesIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );

	// split on a comma with no spaces
	var idList = enterprisesIds.split(',')

	// iterate over array of enterprises ids
	idList.forEach(function (id) {
		// read the Enterprise
		var enterprise = new EnterpriseService(this.http).getEnterprise(id);
		// add the Enterprise if not already assigned
		if ( this.customer.enterprises.indexOf(enterprise) == -1 )
		this.customer.enterprises.push(enterprise);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more enterprisesIds as a Enterprises
	// from a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEnterprises( customerId, enterprisesIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );


	// split on a comma with no spaces
	var idList 					= enterprisesIds.split(',');
	var enterprises 	= this.customer.enterprises;

	if ( enterprises != null && enterprisesIds != null ) {

		// iterate over array of enterprises ids
		enterprises.forEach(function (obj) {
			if ( enterprisesIds.indexOf(obj._id) > -1 ) {
				// remove the Enterprise
				this.customer.enterprises.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more salesOrdersIds as a SalesOrders
	// to a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSalesOrders( customerId, salesOrdersIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );

	// split on a comma with no spaces
	var idList = salesOrdersIds.split(',')

	// iterate over array of salesOrders ids
	idList.forEach(function (id) {
		// read the SalesOrder
		var salesOrder = new SalesOrderService(this.http).getSalesOrder(id);
		// add the SalesOrder if not already assigned
		if ( this.customer.salesOrders.indexOf(salesOrder) == -1 )
		this.customer.salesOrders.push(salesOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more salesOrdersIds as a SalesOrders
	// from a Customer
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSalesOrders( customerId, salesOrdersIds ): Observable<any> {

		// get the Customer
		this.loadHelper( customerId );


	// split on a comma with no spaces
	var idList 					= salesOrdersIds.split(',');
	var salesOrders 	= this.customer.salesOrders;

	if ( salesOrders != null && salesOrdersIds != null ) {

		// iterate over array of salesOrders ids
		salesOrders.forEach(function (obj) {
			if ( salesOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the SalesOrder
				this.customer.salesOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Customer
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Customer/update/' + this.customer;

	return  this.http.post(uri_, this.customer );
}

	//********************************************************************
	// loadHelper - internal helper to load a Customer
	//********************************************************************	
	loadHelper( id ) {
		this.getCustomer(id)
			.subscribe((res : Customer) => {
				this.customer = res;
			});
	}
}