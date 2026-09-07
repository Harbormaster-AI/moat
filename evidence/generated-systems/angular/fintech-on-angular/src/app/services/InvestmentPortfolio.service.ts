import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InvestmentPortfolio} from '../models/InvestmentPortfolio';
import {CustomerService} from '../services/Customer.service';
import {InvestmentAccountService} from '../services/InvestmentAccount.service';
import {TradeOrderService} from '../services/TradeOrder.service';
import {PositionService} from '../services/Position.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InvestmentPortfolioService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	investmentPortfolio : InvestmentPortfolio;

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
	// add a InvestmentPortfolio
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInvestmentPortfolio(portfolioCode, baseCurrency, createdAt, Customer, Accounts, Orders, Holdings, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/InvestmentPortfolio/create';
		const obj = {
			      		portfolioCode: portfolioCode,
      		baseCurrency: baseCurrency,
      		createdAt: createdAt,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Accounts: Accounts != null && Accounts.length > 0 ? Accounts : null,
      		Orders: Orders != null && Orders.length > 0 ? Orders : null,
      		Holdings: Holdings != null && Holdings.length > 0 ? Holdings : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InvestmentPortfolio
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInvestmentPortfolio(portfolioCode, baseCurrency, createdAt, Customer, Accounts, Orders, Holdings, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InvestmentPortfolio/update/' + id;
		const obj = {
				      		portfolioCode: portfolioCode,
      		baseCurrency: baseCurrency,
      		createdAt: createdAt,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Accounts: Accounts != null && Accounts.length > 0 ? Accounts : null,
      		Orders: Orders != null && Orders.length > 0 ? Orders : null,
      		Holdings: Holdings != null && Holdings.length > 0 ? Holdings : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InvestmentPortfolio
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInvestmentPortfolio(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InvestmentPortfolio/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InvestmentPortfolio
	// returns the results untouched as an Observable InvestmentPortfolio
	// InvestmentPortfolio model
	// delegates via URI
	//********************************************************************
	getInvestmentPortfolio(id) : Observable<InvestmentPortfolio> {
		const uri_ = this.apiUrl + '/InvestmentPortfolio/load/' + id;

		return this.http.get<InvestmentPortfolio>(uri_);
	}
	
	//********************************************************************
	// gets all InvestmentPortfolio
	// returns the results untouched as JSON representation of an
	// Observable array of InvestmentPortfolio models
	// delegates via URI
	//********************************************************************
	getInvestmentPortfolios() : Observable<InvestmentPortfolio[]> {
		const uri_ = this.apiUrl + '/InvestmentPortfolio/';

		return this
			.http.get<InvestmentPortfolio[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Customer on a InvestmentPortfolio
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCustomer( investmentPortfolioId, _customerId ): Observable<any> {

		// get the InvestmentPortfolio from storage
		this.loadHelper( investmentPortfolioId );

	// get the Customer from storage
	var tmp 	= new CustomerService(this.http).getCustomer(_customerId);

	// assign the Customer
	this.investmentPortfolio.customer = tmp;

	// save the InvestmentPortfolio
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Customer on a InvestmentPortfolio
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCustomer( investmentPortfolioId ): Observable<any> {

		// get the InvestmentPortfolio from storage
		this.loadHelper( investmentPortfolioId );

	// assign Customer to null
	this.investmentPortfolio.customer = null;

	// save the InvestmentPortfolio
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more accountsIds as a Accounts
	// to a InvestmentPortfolio
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAccounts( investmentPortfolioId, accountsIds ): Observable<any> {

		// get the InvestmentPortfolio
		this.loadHelper( investmentPortfolioId );

	// split on a comma with no spaces
	var idList = accountsIds.split(',')

	// iterate over array of accounts ids
	idList.forEach(function (id) {
		// read the InvestmentAccount
		var investmentAccount = new InvestmentAccountService(this.http).getInvestmentAccount(id);
		// add the InvestmentAccount if not already assigned
		if ( this.investmentPortfolio.accounts.indexOf(investmentAccount) == -1 )
		this.investmentPortfolio.accounts.push(investmentAccount);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more accountsIds as a Accounts
	// from a InvestmentPortfolio
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAccounts( investmentPortfolioId, accountsIds ): Observable<any> {

		// get the InvestmentPortfolio
		this.loadHelper( investmentPortfolioId );


	// split on a comma with no spaces
	var idList 					= accountsIds.split(',');
	var accounts 	= this.investmentPortfolio.accounts;

	if ( accounts != null && accountsIds != null ) {

		// iterate over array of accounts ids
		accounts.forEach(function (obj) {
			if ( accountsIds.indexOf(obj._id) > -1 ) {
				// remove the InvestmentAccount
				this.investmentPortfolio.accounts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more ordersIds as a Orders
	// to a InvestmentPortfolio
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOrders( investmentPortfolioId, ordersIds ): Observable<any> {

		// get the InvestmentPortfolio
		this.loadHelper( investmentPortfolioId );

	// split on a comma with no spaces
	var idList = ordersIds.split(',')

	// iterate over array of orders ids
	idList.forEach(function (id) {
		// read the TradeOrder
		var tradeOrder = new TradeOrderService(this.http).getTradeOrder(id);
		// add the TradeOrder if not already assigned
		if ( this.investmentPortfolio.orders.indexOf(tradeOrder) == -1 )
		this.investmentPortfolio.orders.push(tradeOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more ordersIds as a Orders
	// from a InvestmentPortfolio
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOrders( investmentPortfolioId, ordersIds ): Observable<any> {

		// get the InvestmentPortfolio
		this.loadHelper( investmentPortfolioId );


	// split on a comma with no spaces
	var idList 					= ordersIds.split(',');
	var orders 	= this.investmentPortfolio.orders;

	if ( orders != null && ordersIds != null ) {

		// iterate over array of orders ids
		orders.forEach(function (obj) {
			if ( ordersIds.indexOf(obj._id) > -1 ) {
				// remove the TradeOrder
				this.investmentPortfolio.orders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more holdingsIds as a Holdings
	// to a InvestmentPortfolio
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addHoldings( investmentPortfolioId, holdingsIds ): Observable<any> {

		// get the InvestmentPortfolio
		this.loadHelper( investmentPortfolioId );

	// split on a comma with no spaces
	var idList = holdingsIds.split(',')

	// iterate over array of holdings ids
	idList.forEach(function (id) {
		// read the Position
		var position = new PositionService(this.http).getPosition(id);
		// add the Position if not already assigned
		if ( this.investmentPortfolio.holdings.indexOf(position) == -1 )
		this.investmentPortfolio.holdings.push(position);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more holdingsIds as a Holdings
	// from a InvestmentPortfolio
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeHoldings( investmentPortfolioId, holdingsIds ): Observable<any> {

		// get the InvestmentPortfolio
		this.loadHelper( investmentPortfolioId );


	// split on a comma with no spaces
	var idList 					= holdingsIds.split(',');
	var holdings 	= this.investmentPortfolio.holdings;

	if ( holdings != null && holdingsIds != null ) {

		// iterate over array of holdings ids
		holdings.forEach(function (obj) {
			if ( holdingsIds.indexOf(obj._id) > -1 ) {
				// remove the Position
				this.investmentPortfolio.holdings.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a InvestmentPortfolio
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InvestmentPortfolio/update/' + this.investmentPortfolio;

	return  this.http.post(uri_, this.investmentPortfolio );
}

	//********************************************************************
	// loadHelper - internal helper to load a InvestmentPortfolio
	//********************************************************************	
	loadHelper( id ) {
		this.getInvestmentPortfolio(id)
			.subscribe((res : InvestmentPortfolio) => {
				this.investmentPortfolio = res;
			});
	}
}