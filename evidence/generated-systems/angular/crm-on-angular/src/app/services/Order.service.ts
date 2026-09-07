import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Order} from '../models/Order';
import {OrganizationService} from '../services/Organization.service';
import {AccountService} from '../services/Account.service';
import {OpportunityService} from '../services/Opportunity.service';
import {QuoteService} from '../services/Quote.service';
import {UserService} from '../services/User.service';
import {OrderItemService} from '../services/OrderItem.service';
import {ContractService} from '../services/Contract.service';
import {PriceBookService} from '../services/PriceBook.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class OrderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	order : Order;

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
	// add a Order
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addOrder(orderNumber, orderDate, totalAmount, taxAmount, shippingAmount, Organization, Account, Opportunity, Quote, Owner, Items, Contract, PriceBook, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Order/create';
		const obj = {
			      		orderNumber: orderNumber,
      		orderDate: orderDate,
      		totalAmount: totalAmount,
      		taxAmount: taxAmount,
      		shippingAmount: shippingAmount,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Opportunity: Opportunity != null && Opportunity.length > 0 ? Opportunity : null,
      		Quote: Quote != null && Quote.length > 0 ? Quote : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Items: Items != null && Items.length > 0 ? Items : null,
      		Contract: Contract != null && Contract.length > 0 ? Contract : null,
      		PriceBook: PriceBook != null && PriceBook.length > 0 ? PriceBook : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Order
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateOrder(orderNumber, orderDate, totalAmount, taxAmount, shippingAmount, Organization, Account, Opportunity, Quote, Owner, Items, Contract, PriceBook, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Order/update/' + id;
		const obj = {
				      		orderNumber: orderNumber,
      		orderDate: orderDate,
      		totalAmount: totalAmount,
      		taxAmount: taxAmount,
      		shippingAmount: shippingAmount,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Opportunity: Opportunity != null && Opportunity.length > 0 ? Opportunity : null,
      		Quote: Quote != null && Quote.length > 0 ? Quote : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Items: Items != null && Items.length > 0 ? Items : null,
      		Contract: Contract != null && Contract.length > 0 ? Contract : null,
      		PriceBook: PriceBook != null && PriceBook.length > 0 ? PriceBook : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Order
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteOrder(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Order/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Order
	// returns the results untouched as an Observable Order
	// Order model
	// delegates via URI
	//********************************************************************
	getOrder(id) : Observable<Order> {
		const uri_ = this.apiUrl + '/Order/load/' + id;

		return this.http.get<Order>(uri_);
	}
	
	//********************************************************************
	// gets all Order
	// returns the results untouched as JSON representation of an
	// Observable array of Order models
	// delegates via URI
	//********************************************************************
	getOrders() : Observable<Order[]> {
		const uri_ = this.apiUrl + '/Order/';

		return this
			.http.get<Order[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Order
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( orderId, _organizationId ): Observable<any> {

		// get the Order from storage
		this.loadHelper( orderId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.order.organization = tmp;

	// save the Order
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Order
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( orderId ): Observable<any> {

		// get the Order from storage
		this.loadHelper( orderId );

	// assign Organization to null
	this.order.organization = null;

	// save the Order
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Account on a Order
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAccount( orderId, _accountId ): Observable<any> {

		// get the Order from storage
		this.loadHelper( orderId );

	// get the Account from storage
	var tmp 	= new AccountService(this.http).getAccount(_accountId);

	// assign the Account
	this.order.account = tmp;

	// save the Order
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Account on a Order
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAccount( orderId ): Observable<any> {

		// get the Order from storage
		this.loadHelper( orderId );

	// assign Account to null
	this.order.account = null;

	// save the Order
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Opportunity on a Order
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOpportunity( orderId, _opportunityId ): Observable<any> {

		// get the Order from storage
		this.loadHelper( orderId );

	// get the Opportunity from storage
	var tmp 	= new OpportunityService(this.http).getOpportunity(_opportunityId);

	// assign the Opportunity
	this.order.opportunity = tmp;

	// save the Order
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Opportunity on a Order
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOpportunity( orderId ): Observable<any> {

		// get the Order from storage
		this.loadHelper( orderId );

	// assign Opportunity to null
	this.order.opportunity = null;

	// save the Order
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Quote on a Order
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignQuote( orderId, _quoteId ): Observable<any> {

		// get the Order from storage
		this.loadHelper( orderId );

	// get the Quote from storage
	var tmp 	= new QuoteService(this.http).getQuote(_quoteId);

	// assign the Quote
	this.order.quote = tmp;

	// save the Order
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Quote on a Order
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignQuote( orderId ): Observable<any> {

		// get the Order from storage
		this.loadHelper( orderId );

	// assign Quote to null
	this.order.quote = null;

	// save the Order
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Owner on a Order
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOwner( orderId, _ownerId ): Observable<any> {

		// get the Order from storage
		this.loadHelper( orderId );

	// get the User from storage
	var tmp 	= new UserService(this.http).getUser(_ownerId);

	// assign the Owner
	this.order.owner = tmp;

	// save the Order
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Owner on a Order
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOwner( orderId ): Observable<any> {

		// get the Order from storage
		this.loadHelper( orderId );

	// assign Owner to null
	this.order.owner = null;

	// save the Order
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Contract on a Order
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignContract( orderId, _contractId ): Observable<any> {

		// get the Order from storage
		this.loadHelper( orderId );

	// get the Contract from storage
	var tmp 	= new ContractService(this.http).getContract(_contractId);

	// assign the Contract
	this.order.contract = tmp;

	// save the Order
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Contract on a Order
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignContract( orderId ): Observable<any> {

		// get the Order from storage
		this.loadHelper( orderId );

	// assign Contract to null
	this.order.contract = null;

	// save the Order
	return this.saveHelper();
}

		//********************************************************************
	// assigns a PriceBook on a Order
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPriceBook( orderId, _priceBookId ): Observable<any> {

		// get the Order from storage
		this.loadHelper( orderId );

	// get the PriceBook from storage
	var tmp 	= new PriceBookService(this.http).getPriceBook(_priceBookId);

	// assign the PriceBook
	this.order.priceBook = tmp;

	// save the Order
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a PriceBook on a Order
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPriceBook( orderId ): Observable<any> {

		// get the Order from storage
		this.loadHelper( orderId );

	// assign PriceBook to null
	this.order.priceBook = null;

	// save the Order
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more itemsIds as a Items
	// to a Order
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addItems( orderId, itemsIds ): Observable<any> {

		// get the Order
		this.loadHelper( orderId );

	// split on a comma with no spaces
	var idList = itemsIds.split(',')

	// iterate over array of items ids
	idList.forEach(function (id) {
		// read the OrderItem
		var orderItem = new OrderItemService(this.http).getOrderItem(id);
		// add the OrderItem if not already assigned
		if ( this.order.items.indexOf(orderItem) == -1 )
		this.order.items.push(orderItem);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more itemsIds as a Items
	// from a Order
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeItems( orderId, itemsIds ): Observable<any> {

		// get the Order
		this.loadHelper( orderId );


	// split on a comma with no spaces
	var idList 					= itemsIds.split(',');
	var items 	= this.order.items;

	if ( items != null && itemsIds != null ) {

		// iterate over array of items ids
		items.forEach(function (obj) {
			if ( itemsIds.indexOf(obj._id) > -1 ) {
				// remove the OrderItem
				this.order.items.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Order
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Order/update/' + this.order;

	return  this.http.post(uri_, this.order );
}

	//********************************************************************
	// loadHelper - internal helper to load a Order
	//********************************************************************	
	loadHelper( id ) {
		this.getOrder(id)
			.subscribe((res : Order) => {
				this.order = res;
			});
	}
}