import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Position} from '../models/Position';
import {InvestmentPortfolioService} from '../services/InvestmentPortfolio.service';
import {SecurityService} from '../services/Security.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PositionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	position : Position;

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
	// add a Position
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPosition(quantity, averageCost, marketValue, Portfolio, Security) : Observable<any> {
		const uri_ = this.apiUrl + '/Position/create';
		const obj = {
			      		quantity: quantity,
      		averageCost: averageCost,
      		marketValue: marketValue,
      		Portfolio: Portfolio != null && Portfolio.length > 0 ? Portfolio : null,
			Security: Security != null && Security.length > 0 ? Security : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Position
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePosition(quantity, averageCost, marketValue, Portfolio, Security, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Position/update/' + id;
		const obj = {
				      		quantity: quantity,
      		averageCost: averageCost,
      		marketValue: marketValue,
      		Portfolio: Portfolio != null && Portfolio.length > 0 ? Portfolio : null,
			Security: Security != null && Security.length > 0 ? Security : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Position
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePosition(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Position/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Position
	// returns the results untouched as an Observable Position
	// Position model
	// delegates via URI
	//********************************************************************
	getPosition(id) : Observable<Position> {
		const uri_ = this.apiUrl + '/Position/load/' + id;

		return this.http.get<Position>(uri_);
	}
	
	//********************************************************************
	// gets all Position
	// returns the results untouched as JSON representation of an
	// Observable array of Position models
	// delegates via URI
	//********************************************************************
	getPositions() : Observable<Position[]> {
		const uri_ = this.apiUrl + '/Position/';

		return this
			.http.get<Position[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Portfolio on a Position
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPortfolio( positionId, _portfolioId ): Observable<any> {

		// get the Position from storage
		this.loadHelper( positionId );

	// get the InvestmentPortfolio from storage
	var tmp 	= new InvestmentPortfolioService(this.http).getInvestmentPortfolio(_portfolioId);

	// assign the Portfolio
	this.position.portfolio = tmp;

	// save the Position
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Portfolio on a Position
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPortfolio( positionId ): Observable<any> {

		// get the Position from storage
		this.loadHelper( positionId );

	// assign Portfolio to null
	this.position.portfolio = null;

	// save the Position
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Security on a Position
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSecurity( positionId, _securityId ): Observable<any> {

		// get the Position from storage
		this.loadHelper( positionId );

	// get the Security from storage
	var tmp 	= new SecurityService(this.http).getSecurity(_securityId);

	// assign the Security
	this.position.security = tmp;

	// save the Position
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Security on a Position
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSecurity( positionId ): Observable<any> {

		// get the Position from storage
		this.loadHelper( positionId );

	// assign Security to null
	this.position.security = null;

	// save the Position
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Position
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Position/update/' + this.position;

	return  this.http.post(uri_, this.position );
}

	//********************************************************************
	// loadHelper - internal helper to load a Position
	//********************************************************************	
	loadHelper( id ) {
		this.getPosition(id)
			.subscribe((res : Position) => {
				this.position = res;
			});
	}
}