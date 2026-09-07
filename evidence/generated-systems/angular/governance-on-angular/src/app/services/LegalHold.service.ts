import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {LegalHold} from '../models/LegalHold';
import {RecordsRepositoryService} from '../services/RecordsRepository.service';
import {Record_Service} from '../services/Record_.service';
import {MatterService} from '../services/Matter.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class LegalHoldService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	legalHold : LegalHold;

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
	// add a LegalHold
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addLegalHold(name, reason, issuedDate, releaseDate, Repositories, Records, Matter, HoldStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/LegalHold/create';
		const obj = {
			      		name: name,
      		reason: reason,
      		issuedDate: issuedDate,
      		releaseDate: releaseDate,
      		Repositories: Repositories != null && Repositories.length > 0 ? Repositories : null,
      		Records: Records != null && Records.length > 0 ? Records : null,
      		Matter: Matter != null && Matter.length > 0 ? Matter : null,
			HoldStatus: HoldStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a LegalHold
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateLegalHold(name, reason, issuedDate, releaseDate, Repositories, Records, Matter, HoldStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/LegalHold/update/' + id;
		const obj = {
				      		name: name,
      		reason: reason,
      		issuedDate: issuedDate,
      		releaseDate: releaseDate,
      		Repositories: Repositories != null && Repositories.length > 0 ? Repositories : null,
      		Records: Records != null && Records.length > 0 ? Records : null,
      		Matter: Matter != null && Matter.length > 0 ? Matter : null,
			HoldStatus: HoldStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a LegalHold
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteLegalHold(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/LegalHold/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a LegalHold
	// returns the results untouched as an Observable LegalHold
	// LegalHold model
	// delegates via URI
	//********************************************************************
	getLegalHold(id) : Observable<LegalHold> {
		const uri_ = this.apiUrl + '/LegalHold/load/' + id;

		return this.http.get<LegalHold>(uri_);
	}
	
	//********************************************************************
	// gets all LegalHold
	// returns the results untouched as JSON representation of an
	// Observable array of LegalHold models
	// delegates via URI
	//********************************************************************
	getLegalHolds() : Observable<LegalHold[]> {
		const uri_ = this.apiUrl + '/LegalHold/';

		return this
			.http.get<LegalHold[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Matter on a LegalHold
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignMatter( legalHoldId, _matterId ): Observable<any> {

		// get the LegalHold from storage
		this.loadHelper( legalHoldId );

	// get the Matter from storage
	var tmp 	= new MatterService(this.http).getMatter(_matterId);

	// assign the Matter
	this.legalHold.matter = tmp;

	// save the LegalHold
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Matter on a LegalHold
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignMatter( legalHoldId ): Observable<any> {

		// get the LegalHold from storage
		this.loadHelper( legalHoldId );

	// assign Matter to null
	this.legalHold.matter = null;

	// save the LegalHold
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more repositoriesIds as a Repositories
	// to a LegalHold
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRepositories( legalHoldId, repositoriesIds ): Observable<any> {

		// get the LegalHold
		this.loadHelper( legalHoldId );

	// split on a comma with no spaces
	var idList = repositoriesIds.split(',')

	// iterate over array of repositories ids
	idList.forEach(function (id) {
		// read the RecordsRepository
		var recordsRepository = new RecordsRepositoryService(this.http).getRecordsRepository(id);
		// add the RecordsRepository if not already assigned
		if ( this.legalHold.repositories.indexOf(recordsRepository) == -1 )
		this.legalHold.repositories.push(recordsRepository);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more repositoriesIds as a Repositories
	// from a LegalHold
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRepositories( legalHoldId, repositoriesIds ): Observable<any> {

		// get the LegalHold
		this.loadHelper( legalHoldId );


	// split on a comma with no spaces
	var idList 					= repositoriesIds.split(',');
	var repositories 	= this.legalHold.repositories;

	if ( repositories != null && repositoriesIds != null ) {

		// iterate over array of repositories ids
		repositories.forEach(function (obj) {
			if ( repositoriesIds.indexOf(obj._id) > -1 ) {
				// remove the RecordsRepository
				this.legalHold.repositories.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more recordsIds as a Records
	// to a LegalHold
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRecords( legalHoldId, recordsIds ): Observable<any> {

		// get the LegalHold
		this.loadHelper( legalHoldId );

	// split on a comma with no spaces
	var idList = recordsIds.split(',')

	// iterate over array of records ids
	idList.forEach(function (id) {
		// read the Record_
		var record_ = new Record_Service(this.http).getRecord_(id);
		// add the Record_ if not already assigned
		if ( this.legalHold.records.indexOf(record_) == -1 )
		this.legalHold.records.push(record_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more recordsIds as a Records
	// from a LegalHold
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRecords( legalHoldId, recordsIds ): Observable<any> {

		// get the LegalHold
		this.loadHelper( legalHoldId );


	// split on a comma with no spaces
	var idList 					= recordsIds.split(',');
	var records 	= this.legalHold.records;

	if ( records != null && recordsIds != null ) {

		// iterate over array of records ids
		records.forEach(function (obj) {
			if ( recordsIds.indexOf(obj._id) > -1 ) {
				// remove the Record_
				this.legalHold.records.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a LegalHold
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/LegalHold/update/' + this.legalHold;

	return  this.http.post(uri_, this.legalHold );
}

	//********************************************************************
	// loadHelper - internal helper to load a LegalHold
	//********************************************************************	
	loadHelper( id ) {
		this.getLegalHold(id)
			.subscribe((res : LegalHold) => {
				this.legalHold = res;
			});
	}
}