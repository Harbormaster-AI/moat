import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ContentCategory} from '../models/ContentCategory';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ContentCategoryService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	contentCategory : ContentCategory;

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
	// add a ContentCategory
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addContentCategory(code, name) : Observable<any> {
		const uri_ = this.apiUrl + '/ContentCategory/create';
		const obj = {
			      		code: code,
			name: name
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ContentCategory
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateContentCategory(code, name, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ContentCategory/update/' + id;
		const obj = {
				      		code: code,
			name: name
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ContentCategory
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteContentCategory(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ContentCategory/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ContentCategory
	// returns the results untouched as an Observable ContentCategory
	// ContentCategory model
	// delegates via URI
	//********************************************************************
	getContentCategory(id) : Observable<ContentCategory> {
		const uri_ = this.apiUrl + '/ContentCategory/load/' + id;

		return this.http.get<ContentCategory>(uri_);
	}
	
	//********************************************************************
	// gets all ContentCategory
	// returns the results untouched as JSON representation of an
	// Observable array of ContentCategory models
	// delegates via URI
	//********************************************************************
	getContentCategorys() : Observable<ContentCategory[]> {
		const uri_ = this.apiUrl + '/ContentCategory/';

		return this
			.http.get<ContentCategory[]>(uri_);
	}
	
		
	
	//********************************************************************
	// saveHelper - internal helper to save a ContentCategory
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ContentCategory/update/' + this.contentCategory;

	return  this.http.post(uri_, this.contentCategory );
}

	//********************************************************************
	// loadHelper - internal helper to load a ContentCategory
	//********************************************************************	
	loadHelper( id ) {
		this.getContentCategory(id)
			.subscribe((res : ContentCategory) => {
				this.contentCategory = res;
			});
	}
}