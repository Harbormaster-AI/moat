import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InvestmentAccount} from '../models/InvestmentAccount';
import {InvestmentPortfolioService} from '../services/InvestmentPortfolio.service';
import {TradeService} from '../services/Trade.service';
import {TradeOrderService} from '../services/TradeOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InvestmentAccountService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	investmentAccount : InvestmentAccount;

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
	// add a InvestmentAccount
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInvestmentAccount(accountNumber, baseCurrency, balance, Portfolio, Trades, Orders, AccountType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/InvestmentAccount/create';
		const obj = {
			      		accountNumber: accountNumber,
      		baseCurrency: baseCurrency,
      		balance: balance,
      		Portfolio: Portfolio != null && Portfolio.length > 0 ? Portfolio : null,
      		Trades: Trades != null && Trades.length > 0 ? Trades : null,
      		Orders: Orders != null && Orders.length > 0 ? Orders : null,
      		AccountType: AccountType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InvestmentAccount
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInvestmentAccount(accountNumber, baseCurrency, balance, Portfolio, Trades, Orders, AccountType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InvestmentAccount/update/' + id;
		const obj = {
				      		accountNumber: accountNumber,
      		baseCurrency: baseCurrency,
      		balance: balance,
      		Portfolio: Portfolio != null && Portfolio.length > 0 ? Portfolio : null,
      		Trades: Trades != null && Trades.length > 0 ? Trades : null,
      		Orders: Orders != null && Orders.length > 0 ? Orders : null,
      		AccountType: AccountType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InvestmentAccount
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInvestmentAccount(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InvestmentAccount/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InvestmentAccount
	// returns the results untouched as an Observable InvestmentAccount
	// InvestmentAccount model
	// delegates via URI
	//********************************************************************
	getInvestmentAccount(id) : Observable<InvestmentAccount> {
		const uri_ = this.apiUrl + '/InvestmentAccount/load/' + id;

		return this.http.get<InvestmentAccount>(uri_);
	}
	
	//********************************************************************
	// gets all InvestmentAccount
	// returns the results untouched as JSON representation of an
	// Observable array of InvestmentAccount models
	// delegates via URI
	//********************************************************************
	getInvestmentAccounts() : Observable<InvestmentAccount[]> {
		const uri_ = this.apiUrl + '/InvestmentAccount/';

		return this
			.http.get<InvestmentAccount[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Portfolio on a InvestmentAccount
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPortfolio( investmentAccountId, _portfolioId ): Observable<any> {

		// get the InvestmentAccount from storage
		this.loadHelper( investmentAccountId );

	// get the InvestmentPortfolio from storage
	var tmp 	= new InvestmentPortfolioService(this.http).getInvestmentPortfolio(_portfolioId);

	// assign the Portfolio
	this.investmentAccount.portfolio = tmp;

	// save the InvestmentAccount
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Portfolio on a InvestmentAccount
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPortfolio( investmentAccountId ): Observable<any> {

		// get the InvestmentAccount from storage
		this.loadHelper( investmentAccountId );

	// assign Portfolio to null
	this.investmentAccount.portfolio = null;

	// save the InvestmentAccount
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more tradesIds as a Trades
	// to a InvestmentAccount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTrades( investmentAccountId, tradesIds ): Observable<any> {

		// get the InvestmentAccount
		this.loadHelper( investmentAccountId );

	// split on a comma with no spaces
	var idList = tradesIds.split(',')

	// iterate over array of trades ids
	idList.forEach(function (id) {
		// read the Trade
		var trade = new TradeService(this.http).getTrade(id);
		// add the Trade if not already assigned
		if ( this.investmentAccount.trades.indexOf(trade) == -1 )
		this.investmentAccount.trades.push(trade);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more tradesIds as a Trades
	// from a InvestmentAccount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTrades( investmentAccountId, tradesIds ): Observable<any> {

		// get the InvestmentAccount
		this.loadHelper( investmentAccountId );


	// split on a comma with no spaces
	var idList 					= tradesIds.split(',');
	var trades 	= this.investmentAccount.trades;

	if ( trades != null && tradesIds != null ) {

		// iterate over array of trades ids
		trades.forEach(function (obj) {
			if ( tradesIds.indexOf(obj._id) > -1 ) {
				// remove the Trade
				this.investmentAccount.trades.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more ordersIds as a Orders
	// to a InvestmentAccount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOrders( investmentAccountId, ordersIds ): Observable<any> {

		// get the InvestmentAccount
		this.loadHelper( investmentAccountId );

	// split on a comma with no spaces
	var idList = ordersIds.split(',')

	// iterate over array of orders ids
	idList.forEach(function (id) {
		// read the TradeOrder
		var tradeOrder = new TradeOrderService(this.http).getTradeOrder(id);
		// add the TradeOrder if not already assigned
		if ( this.investmentAccount.orders.indexOf(tradeOrder) == -1 )
		this.investmentAccount.orders.push(tradeOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more ordersIds as a Orders
	// from a InvestmentAccount
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOrders( investmentAccountId, ordersIds ): Observable<any> {

		// get the InvestmentAccount
		this.loadHelper( investmentAccountId );


	// split on a comma with no spaces
	var idList 					= ordersIds.split(',');
	var orders 	= this.investmentAccount.orders;

	if ( orders != null && ordersIds != null ) {

		// iterate over array of orders ids
		orders.forEach(function (obj) {
			if ( ordersIds.indexOf(obj._id) > -1 ) {
				// remove the TradeOrder
				this.investmentAccount.orders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a InvestmentAccount
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InvestmentAccount/update/' + this.investmentAccount;

	return  this.http.post(uri_, this.investmentAccount );
}

	//********************************************************************
	// loadHelper - internal helper to load a InvestmentAccount
	//********************************************************************	
	loadHelper( id ) {
		this.getInvestmentAccount(id)
			.subscribe((res : InvestmentAccount) => {
				this.investmentAccount = res;
			});
	}
}