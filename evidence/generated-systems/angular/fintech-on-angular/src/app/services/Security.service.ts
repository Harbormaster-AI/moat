import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Security} from '../models/Security';
import {PositionService} from '../services/Position.service';
import {TradeService} from '../services/Trade.service';
import {TradeOrderService} from '../services/TradeOrder.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class SecurityService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	security : Security;

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
	// add a Security
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addSecurity(symbol, isin, cusip, currency, Positions, Trades, Orders, SecurityType) : Observable<any> {
		const uri_ = this.apiUrl + '/Security/create';
		const obj = {
			      		symbol: symbol,
      		isin: isin,
      		cusip: cusip,
      		currency: currency,
      		Positions: Positions != null && Positions.length > 0 ? Positions : null,
      		Trades: Trades != null && Trades.length > 0 ? Trades : null,
      		Orders: Orders != null && Orders.length > 0 ? Orders : null,
			SecurityType: SecurityType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Security
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateSecurity(symbol, isin, cusip, currency, Positions, Trades, Orders, SecurityType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Security/update/' + id;
		const obj = {
				      		symbol: symbol,
      		isin: isin,
      		cusip: cusip,
      		currency: currency,
      		Positions: Positions != null && Positions.length > 0 ? Positions : null,
      		Trades: Trades != null && Trades.length > 0 ? Trades : null,
      		Orders: Orders != null && Orders.length > 0 ? Orders : null,
			SecurityType: SecurityType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Security
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteSecurity(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Security/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Security
	// returns the results untouched as an Observable Security
	// Security model
	// delegates via URI
	//********************************************************************
	getSecurity(id) : Observable<Security> {
		const uri_ = this.apiUrl + '/Security/load/' + id;

		return this.http.get<Security>(uri_);
	}
	
	//********************************************************************
	// gets all Security
	// returns the results untouched as JSON representation of an
	// Observable array of Security models
	// delegates via URI
	//********************************************************************
	getSecuritys() : Observable<Security[]> {
		const uri_ = this.apiUrl + '/Security/';

		return this
			.http.get<Security[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more positionsIds as a Positions
	// to a Security
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPositions( securityId, positionsIds ): Observable<any> {

		// get the Security
		this.loadHelper( securityId );

	// split on a comma with no spaces
	var idList = positionsIds.split(',')

	// iterate over array of positions ids
	idList.forEach(function (id) {
		// read the Position
		var position = new PositionService(this.http).getPosition(id);
		// add the Position if not already assigned
		if ( this.security.positions.indexOf(position) == -1 )
		this.security.positions.push(position);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more positionsIds as a Positions
	// from a Security
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePositions( securityId, positionsIds ): Observable<any> {

		// get the Security
		this.loadHelper( securityId );


	// split on a comma with no spaces
	var idList 					= positionsIds.split(',');
	var positions 	= this.security.positions;

	if ( positions != null && positionsIds != null ) {

		// iterate over array of positions ids
		positions.forEach(function (obj) {
			if ( positionsIds.indexOf(obj._id) > -1 ) {
				// remove the Position
				this.security.positions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more tradesIds as a Trades
	// to a Security
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTrades( securityId, tradesIds ): Observable<any> {

		// get the Security
		this.loadHelper( securityId );

	// split on a comma with no spaces
	var idList = tradesIds.split(',')

	// iterate over array of trades ids
	idList.forEach(function (id) {
		// read the Trade
		var trade = new TradeService(this.http).getTrade(id);
		// add the Trade if not already assigned
		if ( this.security.trades.indexOf(trade) == -1 )
		this.security.trades.push(trade);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more tradesIds as a Trades
	// from a Security
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTrades( securityId, tradesIds ): Observable<any> {

		// get the Security
		this.loadHelper( securityId );


	// split on a comma with no spaces
	var idList 					= tradesIds.split(',');
	var trades 	= this.security.trades;

	if ( trades != null && tradesIds != null ) {

		// iterate over array of trades ids
		trades.forEach(function (obj) {
			if ( tradesIds.indexOf(obj._id) > -1 ) {
				// remove the Trade
				this.security.trades.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more ordersIds as a Orders
	// to a Security
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addOrders( securityId, ordersIds ): Observable<any> {

		// get the Security
		this.loadHelper( securityId );

	// split on a comma with no spaces
	var idList = ordersIds.split(',')

	// iterate over array of orders ids
	idList.forEach(function (id) {
		// read the TradeOrder
		var tradeOrder = new TradeOrderService(this.http).getTradeOrder(id);
		// add the TradeOrder if not already assigned
		if ( this.security.orders.indexOf(tradeOrder) == -1 )
		this.security.orders.push(tradeOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more ordersIds as a Orders
	// from a Security
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeOrders( securityId, ordersIds ): Observable<any> {

		// get the Security
		this.loadHelper( securityId );


	// split on a comma with no spaces
	var idList 					= ordersIds.split(',');
	var orders 	= this.security.orders;

	if ( orders != null && ordersIds != null ) {

		// iterate over array of orders ids
		orders.forEach(function (obj) {
			if ( ordersIds.indexOf(obj._id) > -1 ) {
				// remove the TradeOrder
				this.security.orders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Security
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Security/update/' + this.security;

	return  this.http.post(uri_, this.security );
}

	//********************************************************************
	// loadHelper - internal helper to load a Security
	//********************************************************************	
	loadHelper( id ) {
		this.getSecurity(id)
			.subscribe((res : Security) => {
				this.security = res;
			});
	}
}