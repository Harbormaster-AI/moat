import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Application} from '../models/Application';
import {CustomerService} from '../services/Customer.service';
import {InsuranceProductService} from '../services/InsuranceProduct.service';
import {DistributorService} from '../services/Distributor.service';
import {QuoteService} from '../services/Quote.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ApplicationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	application : Application;

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
	// add a Application
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addApplication(applicationNumber, submissionDate, Customer, Product, Distributor, Quotes, SelectedQuote, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Application/create';
		const obj = {
			      		applicationNumber: applicationNumber,
      		submissionDate: submissionDate,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Product: Product != null && Product.length > 0 ? Product : null,
      		Distributor: Distributor != null && Distributor.length > 0 ? Distributor : null,
      		Quotes: Quotes != null && Quotes.length > 0 ? Quotes : null,
      		SelectedQuote: SelectedQuote != null && SelectedQuote.length > 0 ? SelectedQuote : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Application
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateApplication(applicationNumber, submissionDate, Customer, Product, Distributor, Quotes, SelectedQuote, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Application/update/' + id;
		const obj = {
				      		applicationNumber: applicationNumber,
      		submissionDate: submissionDate,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Product: Product != null && Product.length > 0 ? Product : null,
      		Distributor: Distributor != null && Distributor.length > 0 ? Distributor : null,
      		Quotes: Quotes != null && Quotes.length > 0 ? Quotes : null,
      		SelectedQuote: SelectedQuote != null && SelectedQuote.length > 0 ? SelectedQuote : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Application
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteApplication(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Application/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Application
	// returns the results untouched as an Observable Application
	// Application model
	// delegates via URI
	//********************************************************************
	getApplication(id) : Observable<Application> {
		const uri_ = this.apiUrl + '/Application/load/' + id;

		return this.http.get<Application>(uri_);
	}
	
	//********************************************************************
	// gets all Application
	// returns the results untouched as JSON representation of an
	// Observable array of Application models
	// delegates via URI
	//********************************************************************
	getApplications() : Observable<Application[]> {
		const uri_ = this.apiUrl + '/Application/';

		return this
			.http.get<Application[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Customer on a Application
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCustomer( applicationId, _customerId ): Observable<any> {

		// get the Application from storage
		this.loadHelper( applicationId );

	// get the Customer from storage
	var tmp 	= new CustomerService(this.http).getCustomer(_customerId);

	// assign the Customer
	this.application.customer = tmp;

	// save the Application
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Customer on a Application
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCustomer( applicationId ): Observable<any> {

		// get the Application from storage
		this.loadHelper( applicationId );

	// assign Customer to null
	this.application.customer = null;

	// save the Application
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Product on a Application
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignProduct( applicationId, _productId ): Observable<any> {

		// get the Application from storage
		this.loadHelper( applicationId );

	// get the InsuranceProduct from storage
	var tmp 	= new InsuranceProductService(this.http).getInsuranceProduct(_productId);

	// assign the Product
	this.application.product = tmp;

	// save the Application
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Product on a Application
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignProduct( applicationId ): Observable<any> {

		// get the Application from storage
		this.loadHelper( applicationId );

	// assign Product to null
	this.application.product = null;

	// save the Application
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Distributor on a Application
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDistributor( applicationId, _distributorId ): Observable<any> {

		// get the Application from storage
		this.loadHelper( applicationId );

	// get the Distributor from storage
	var tmp 	= new DistributorService(this.http).getDistributor(_distributorId);

	// assign the Distributor
	this.application.distributor = tmp;

	// save the Application
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Distributor on a Application
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDistributor( applicationId ): Observable<any> {

		// get the Application from storage
		this.loadHelper( applicationId );

	// assign Distributor to null
	this.application.distributor = null;

	// save the Application
	return this.saveHelper();
}

		//********************************************************************
	// assigns a SelectedQuote on a Application
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSelectedQuote( applicationId, _selectedQuoteId ): Observable<any> {

		// get the Application from storage
		this.loadHelper( applicationId );

	// get the Quote from storage
	var tmp 	= new QuoteService(this.http).getQuote(_selectedQuoteId);

	// assign the SelectedQuote
	this.application.selectedQuote = tmp;

	// save the Application
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a SelectedQuote on a Application
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSelectedQuote( applicationId ): Observable<any> {

		// get the Application from storage
		this.loadHelper( applicationId );

	// assign SelectedQuote to null
	this.application.selectedQuote = null;

	// save the Application
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more quotesIds as a Quotes
	// to a Application
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addQuotes( applicationId, quotesIds ): Observable<any> {

		// get the Application
		this.loadHelper( applicationId );

	// split on a comma with no spaces
	var idList = quotesIds.split(',')

	// iterate over array of quotes ids
	idList.forEach(function (id) {
		// read the Quote
		var quote = new QuoteService(this.http).getQuote(id);
		// add the Quote if not already assigned
		if ( this.application.quotes.indexOf(quote) == -1 )
		this.application.quotes.push(quote);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more quotesIds as a Quotes
	// from a Application
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeQuotes( applicationId, quotesIds ): Observable<any> {

		// get the Application
		this.loadHelper( applicationId );


	// split on a comma with no spaces
	var idList 					= quotesIds.split(',');
	var quotes 	= this.application.quotes;

	if ( quotes != null && quotesIds != null ) {

		// iterate over array of quotes ids
		quotes.forEach(function (obj) {
			if ( quotesIds.indexOf(obj._id) > -1 ) {
				// remove the Quote
				this.application.quotes.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Application
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Application/update/' + this.application;

	return  this.http.post(uri_, this.application );
}

	//********************************************************************
	// loadHelper - internal helper to load a Application
	//********************************************************************	
	loadHelper( id ) {
		this.getApplication(id)
			.subscribe((res : Application) => {
				this.application = res;
			});
	}
}