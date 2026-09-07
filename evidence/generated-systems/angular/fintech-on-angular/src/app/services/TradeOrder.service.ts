import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {TradeOrder} from '../models/TradeOrder';
import {InvestmentPortfolioService} from '../services/InvestmentPortfolio.service';
import {SecurityService} from '../services/Security.service';
import {TradeService} from '../services/Trade.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TradeOrderService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	tradeOrder : TradeOrder;

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
	// add a TradeOrder
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTradeOrder(orderId, quantity, limitPrice, placedAt, Portfolio, Security, Trades, Side, Type, Status, TimeInForce) : Observable<any> {
		const uri_ = this.apiUrl + '/TradeOrder/create';
		const obj = {
			      		orderId: orderId,
      		quantity: quantity,
      		limitPrice: limitPrice,
      		placedAt: placedAt,
      		Portfolio: Portfolio != null && Portfolio.length > 0 ? Portfolio : null,
      		Security: Security != null && Security.length > 0 ? Security : null,
      		Trades: Trades != null && Trades.length > 0 ? Trades : null,
      		Side: Side,
      		Type: Type,
      		Status: Status,
			TimeInForce: TimeInForce
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a TradeOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTradeOrder(orderId, quantity, limitPrice, placedAt, Portfolio, Security, Trades, Side, Type, Status, TimeInForce, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/TradeOrder/update/' + id;
		const obj = {
				      		orderId: orderId,
      		quantity: quantity,
      		limitPrice: limitPrice,
      		placedAt: placedAt,
      		Portfolio: Portfolio != null && Portfolio.length > 0 ? Portfolio : null,
      		Security: Security != null && Security.length > 0 ? Security : null,
      		Trades: Trades != null && Trades.length > 0 ? Trades : null,
      		Side: Side,
      		Type: Type,
      		Status: Status,
			TimeInForce: TimeInForce
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a TradeOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTradeOrder(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/TradeOrder/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a TradeOrder
	// returns the results untouched as an Observable TradeOrder
	// TradeOrder model
	// delegates via URI
	//********************************************************************
	getTradeOrder(id) : Observable<TradeOrder> {
		const uri_ = this.apiUrl + '/TradeOrder/load/' + id;

		return this.http.get<TradeOrder>(uri_);
	}
	
	//********************************************************************
	// gets all TradeOrder
	// returns the results untouched as JSON representation of an
	// Observable array of TradeOrder models
	// delegates via URI
	//********************************************************************
	getTradeOrders() : Observable<TradeOrder[]> {
		const uri_ = this.apiUrl + '/TradeOrder/';

		return this
			.http.get<TradeOrder[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Portfolio on a TradeOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPortfolio( tradeOrderId, _portfolioId ): Observable<any> {

		// get the TradeOrder from storage
		this.loadHelper( tradeOrderId );

	// get the InvestmentPortfolio from storage
	var tmp 	= new InvestmentPortfolioService(this.http).getInvestmentPortfolio(_portfolioId);

	// assign the Portfolio
	this.tradeOrder.portfolio = tmp;

	// save the TradeOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Portfolio on a TradeOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPortfolio( tradeOrderId ): Observable<any> {

		// get the TradeOrder from storage
		this.loadHelper( tradeOrderId );

	// assign Portfolio to null
	this.tradeOrder.portfolio = null;

	// save the TradeOrder
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Security on a TradeOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSecurity( tradeOrderId, _securityId ): Observable<any> {

		// get the TradeOrder from storage
		this.loadHelper( tradeOrderId );

	// get the Security from storage
	var tmp 	= new SecurityService(this.http).getSecurity(_securityId);

	// assign the Security
	this.tradeOrder.security = tmp;

	// save the TradeOrder
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Security on a TradeOrder
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSecurity( tradeOrderId ): Observable<any> {

		// get the TradeOrder from storage
		this.loadHelper( tradeOrderId );

	// assign Security to null
	this.tradeOrder.security = null;

	// save the TradeOrder
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more tradesIds as a Trades
	// to a TradeOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTrades( tradeOrderId, tradesIds ): Observable<any> {

		// get the TradeOrder
		this.loadHelper( tradeOrderId );

	// split on a comma with no spaces
	var idList = tradesIds.split(',')

	// iterate over array of trades ids
	idList.forEach(function (id) {
		// read the Trade
		var trade = new TradeService(this.http).getTrade(id);
		// add the Trade if not already assigned
		if ( this.tradeOrder.trades.indexOf(trade) == -1 )
		this.tradeOrder.trades.push(trade);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more tradesIds as a Trades
	// from a TradeOrder
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTrades( tradeOrderId, tradesIds ): Observable<any> {

		// get the TradeOrder
		this.loadHelper( tradeOrderId );


	// split on a comma with no spaces
	var idList 					= tradesIds.split(',');
	var trades 	= this.tradeOrder.trades;

	if ( trades != null && tradesIds != null ) {

		// iterate over array of trades ids
		trades.forEach(function (obj) {
			if ( tradesIds.indexOf(obj._id) > -1 ) {
				// remove the Trade
				this.tradeOrder.trades.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a TradeOrder
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/TradeOrder/update/' + this.tradeOrder;

	return  this.http.post(uri_, this.tradeOrder );
}

	//********************************************************************
	// loadHelper - internal helper to load a TradeOrder
	//********************************************************************	
	loadHelper( id ) {
		this.getTradeOrder(id)
			.subscribe((res : TradeOrder) => {
				this.tradeOrder = res;
			});
	}
}