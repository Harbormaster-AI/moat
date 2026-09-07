import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Trade} from '../models/Trade';
import {TradeOrderService} from '../services/TradeOrder.service';
import {SecurityService} from '../services/Security.service';
import {InvestmentAccountService} from '../services/InvestmentAccount.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TradeService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	trade : Trade;

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
	// add a Trade
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTrade(executedAt, quantity, price, fees, settlementDate, Order, Security, InvestmentAccount) : Observable<any> {
		const uri_ = this.apiUrl + '/Trade/create';
		const obj = {
			      		executedAt: executedAt,
      		quantity: quantity,
      		price: price,
      		fees: fees,
      		settlementDate: settlementDate,
      		Order: Order != null && Order.length > 0 ? Order : null,
      		Security: Security != null && Security.length > 0 ? Security : null,
			InvestmentAccount: InvestmentAccount != null && InvestmentAccount.length > 0 ? InvestmentAccount : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Trade
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTrade(executedAt, quantity, price, fees, settlementDate, Order, Security, InvestmentAccount, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Trade/update/' + id;
		const obj = {
				      		executedAt: executedAt,
      		quantity: quantity,
      		price: price,
      		fees: fees,
      		settlementDate: settlementDate,
      		Order: Order != null && Order.length > 0 ? Order : null,
      		Security: Security != null && Security.length > 0 ? Security : null,
			InvestmentAccount: InvestmentAccount != null && InvestmentAccount.length > 0 ? InvestmentAccount : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Trade
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTrade(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Trade/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Trade
	// returns the results untouched as an Observable Trade
	// Trade model
	// delegates via URI
	//********************************************************************
	getTrade(id) : Observable<Trade> {
		const uri_ = this.apiUrl + '/Trade/load/' + id;

		return this.http.get<Trade>(uri_);
	}
	
	//********************************************************************
	// gets all Trade
	// returns the results untouched as JSON representation of an
	// Observable array of Trade models
	// delegates via URI
	//********************************************************************
	getTrades() : Observable<Trade[]> {
		const uri_ = this.apiUrl + '/Trade/';

		return this
			.http.get<Trade[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Order on a Trade
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrder( tradeId, _orderId ): Observable<any> {

		// get the Trade from storage
		this.loadHelper( tradeId );

	// get the TradeOrder from storage
	var tmp 	= new TradeOrderService(this.http).getTradeOrder(_orderId);

	// assign the Order
	this.trade.order = tmp;

	// save the Trade
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Order on a Trade
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrder( tradeId ): Observable<any> {

		// get the Trade from storage
		this.loadHelper( tradeId );

	// assign Order to null
	this.trade.order = null;

	// save the Trade
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Security on a Trade
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSecurity( tradeId, _securityId ): Observable<any> {

		// get the Trade from storage
		this.loadHelper( tradeId );

	// get the Security from storage
	var tmp 	= new SecurityService(this.http).getSecurity(_securityId);

	// assign the Security
	this.trade.security = tmp;

	// save the Trade
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Security on a Trade
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSecurity( tradeId ): Observable<any> {

		// get the Trade from storage
		this.loadHelper( tradeId );

	// assign Security to null
	this.trade.security = null;

	// save the Trade
	return this.saveHelper();
}

		//********************************************************************
	// assigns a InvestmentAccount on a Trade
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInvestmentAccount( tradeId, _investmentAccountId ): Observable<any> {

		// get the Trade from storage
		this.loadHelper( tradeId );

	// get the InvestmentAccount from storage
	var tmp 	= new InvestmentAccountService(this.http).getInvestmentAccount(_investmentAccountId);

	// assign the InvestmentAccount
	this.trade.investmentAccount = tmp;

	// save the Trade
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a InvestmentAccount on a Trade
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInvestmentAccount( tradeId ): Observable<any> {

		// get the Trade from storage
		this.loadHelper( tradeId );

	// assign InvestmentAccount to null
	this.trade.investmentAccount = null;

	// save the Trade
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Trade
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Trade/update/' + this.trade;

	return  this.http.post(uri_, this.trade );
}

	//********************************************************************
	// loadHelper - internal helper to load a Trade
	//********************************************************************	
	loadHelper( id ) {
		this.getTrade(id)
			.subscribe((res : Trade) => {
				this.trade = res;
			});
	}
}