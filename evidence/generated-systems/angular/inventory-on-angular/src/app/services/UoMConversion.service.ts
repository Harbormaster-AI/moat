import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {UoMConversion} from '../models/UoMConversion';
import {StockKeepingUnitService} from '../services/StockKeepingUnit.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class UoMConversionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	uoMConversion : UoMConversion;

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
	// add a UoMConversion
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addUoMConversion(factor, precision, Sku, FromUnit, ToUnit) : Observable<any> {
		const uri_ = this.apiUrl + '/UoMConversion/create';
		const obj = {
			      		factor: factor,
      		precision: precision,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		FromUnit: FromUnit,
			ToUnit: ToUnit
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a UoMConversion
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateUoMConversion(factor, precision, Sku, FromUnit, ToUnit, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/UoMConversion/update/' + id;
		const obj = {
				      		factor: factor,
      		precision: precision,
      		Sku: Sku != null && Sku.length > 0 ? Sku : null,
      		FromUnit: FromUnit,
			ToUnit: ToUnit
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a UoMConversion
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteUoMConversion(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/UoMConversion/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a UoMConversion
	// returns the results untouched as an Observable UoMConversion
	// UoMConversion model
	// delegates via URI
	//********************************************************************
	getUoMConversion(id) : Observable<UoMConversion> {
		const uri_ = this.apiUrl + '/UoMConversion/load/' + id;

		return this.http.get<UoMConversion>(uri_);
	}
	
	//********************************************************************
	// gets all UoMConversion
	// returns the results untouched as JSON representation of an
	// Observable array of UoMConversion models
	// delegates via URI
	//********************************************************************
	getUoMConversions() : Observable<UoMConversion[]> {
		const uri_ = this.apiUrl + '/UoMConversion/';

		return this
			.http.get<UoMConversion[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Sku on a UoMConversion
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignSku( uoMConversionId, _skuId ): Observable<any> {

		// get the UoMConversion from storage
		this.loadHelper( uoMConversionId );

	// get the StockKeepingUnit from storage
	var tmp 	= new StockKeepingUnitService(this.http).getStockKeepingUnit(_skuId);

	// assign the Sku
	this.uoMConversion.sku = tmp;

	// save the UoMConversion
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Sku on a UoMConversion
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignSku( uoMConversionId ): Observable<any> {

		// get the UoMConversion from storage
		this.loadHelper( uoMConversionId );

	// assign Sku to null
	this.uoMConversion.sku = null;

	// save the UoMConversion
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a UoMConversion
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/UoMConversion/update/' + this.uoMConversion;

	return  this.http.post(uri_, this.uoMConversion );
}

	//********************************************************************
	// loadHelper - internal helper to load a UoMConversion
	//********************************************************************	
	loadHelper( id ) {
		this.getUoMConversion(id)
			.subscribe((res : UoMConversion) => {
				this.uoMConversion = res;
			});
	}
}