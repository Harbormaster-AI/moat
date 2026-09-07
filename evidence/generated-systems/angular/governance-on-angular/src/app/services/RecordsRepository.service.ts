import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {RecordsRepository} from '../models/RecordsRepository';
import {OrganizationService} from '../services/Organization.service';
import {Record_Service} from '../services/Record_.service';
import {System_Service} from '../services/System_.service';
import {RetentionScheduleService} from '../services/RetentionSchedule.service';
import {LegalHoldService} from '../services/LegalHold.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class RecordsRepositoryService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	recordsRepository : RecordsRepository;

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
	// add a RecordsRepository
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addRecordsRepository(name, location, ownerDepartment, Organization, Records, Systems, RetentionSchedules, LegalHolds, RepositoryType) : Observable<any> {
		const uri_ = this.apiUrl + '/RecordsRepository/create';
		const obj = {
			      		name: name,
      		location: location,
      		ownerDepartment: ownerDepartment,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Records: Records != null && Records.length > 0 ? Records : null,
      		Systems: Systems != null && Systems.length > 0 ? Systems : null,
      		RetentionSchedules: RetentionSchedules != null && RetentionSchedules.length > 0 ? RetentionSchedules : null,
      		LegalHolds: LegalHolds != null && LegalHolds.length > 0 ? LegalHolds : null,
			RepositoryType: RepositoryType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a RecordsRepository
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateRecordsRepository(name, location, ownerDepartment, Organization, Records, Systems, RetentionSchedules, LegalHolds, RepositoryType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/RecordsRepository/update/' + id;
		const obj = {
				      		name: name,
      		location: location,
      		ownerDepartment: ownerDepartment,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Records: Records != null && Records.length > 0 ? Records : null,
      		Systems: Systems != null && Systems.length > 0 ? Systems : null,
      		RetentionSchedules: RetentionSchedules != null && RetentionSchedules.length > 0 ? RetentionSchedules : null,
      		LegalHolds: LegalHolds != null && LegalHolds.length > 0 ? LegalHolds : null,
			RepositoryType: RepositoryType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a RecordsRepository
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteRecordsRepository(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/RecordsRepository/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a RecordsRepository
	// returns the results untouched as an Observable RecordsRepository
	// RecordsRepository model
	// delegates via URI
	//********************************************************************
	getRecordsRepository(id) : Observable<RecordsRepository> {
		const uri_ = this.apiUrl + '/RecordsRepository/load/' + id;

		return this.http.get<RecordsRepository>(uri_);
	}
	
	//********************************************************************
	// gets all RecordsRepository
	// returns the results untouched as JSON representation of an
	// Observable array of RecordsRepository models
	// delegates via URI
	//********************************************************************
	getRecordsRepositorys() : Observable<RecordsRepository[]> {
		const uri_ = this.apiUrl + '/RecordsRepository/';

		return this
			.http.get<RecordsRepository[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a RecordsRepository
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( recordsRepositoryId, _organizationId ): Observable<any> {

		// get the RecordsRepository from storage
		this.loadHelper( recordsRepositoryId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.recordsRepository.organization = tmp;

	// save the RecordsRepository
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a RecordsRepository
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( recordsRepositoryId ): Observable<any> {

		// get the RecordsRepository from storage
		this.loadHelper( recordsRepositoryId );

	// assign Organization to null
	this.recordsRepository.organization = null;

	// save the RecordsRepository
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more recordsIds as a Records
	// to a RecordsRepository
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRecords( recordsRepositoryId, recordsIds ): Observable<any> {

		// get the RecordsRepository
		this.loadHelper( recordsRepositoryId );

	// split on a comma with no spaces
	var idList = recordsIds.split(',')

	// iterate over array of records ids
	idList.forEach(function (id) {
		// read the Record_
		var record_ = new Record_Service(this.http).getRecord_(id);
		// add the Record_ if not already assigned
		if ( this.recordsRepository.records.indexOf(record_) == -1 )
		this.recordsRepository.records.push(record_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more recordsIds as a Records
	// from a RecordsRepository
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRecords( recordsRepositoryId, recordsIds ): Observable<any> {

		// get the RecordsRepository
		this.loadHelper( recordsRepositoryId );


	// split on a comma with no spaces
	var idList 					= recordsIds.split(',');
	var records 	= this.recordsRepository.records;

	if ( records != null && recordsIds != null ) {

		// iterate over array of records ids
		records.forEach(function (obj) {
			if ( recordsIds.indexOf(obj._id) > -1 ) {
				// remove the Record_
				this.recordsRepository.records.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more systemsIds as a Systems
	// to a RecordsRepository
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSystems( recordsRepositoryId, systemsIds ): Observable<any> {

		// get the RecordsRepository
		this.loadHelper( recordsRepositoryId );

	// split on a comma with no spaces
	var idList = systemsIds.split(',')

	// iterate over array of systems ids
	idList.forEach(function (id) {
		// read the System_
		var system_ = new System_Service(this.http).getSystem_(id);
		// add the System_ if not already assigned
		if ( this.recordsRepository.systems.indexOf(system_) == -1 )
		this.recordsRepository.systems.push(system_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more systemsIds as a Systems
	// from a RecordsRepository
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSystems( recordsRepositoryId, systemsIds ): Observable<any> {

		// get the RecordsRepository
		this.loadHelper( recordsRepositoryId );


	// split on a comma with no spaces
	var idList 					= systemsIds.split(',');
	var systems 	= this.recordsRepository.systems;

	if ( systems != null && systemsIds != null ) {

		// iterate over array of systems ids
		systems.forEach(function (obj) {
			if ( systemsIds.indexOf(obj._id) > -1 ) {
				// remove the System_
				this.recordsRepository.systems.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more retentionSchedulesIds as a RetentionSchedules
	// to a RecordsRepository
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRetentionSchedules( recordsRepositoryId, retentionSchedulesIds ): Observable<any> {

		// get the RecordsRepository
		this.loadHelper( recordsRepositoryId );

	// split on a comma with no spaces
	var idList = retentionSchedulesIds.split(',')

	// iterate over array of retentionSchedules ids
	idList.forEach(function (id) {
		// read the RetentionSchedule
		var retentionSchedule = new RetentionScheduleService(this.http).getRetentionSchedule(id);
		// add the RetentionSchedule if not already assigned
		if ( this.recordsRepository.retentionSchedules.indexOf(retentionSchedule) == -1 )
		this.recordsRepository.retentionSchedules.push(retentionSchedule);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more retentionSchedulesIds as a RetentionSchedules
	// from a RecordsRepository
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRetentionSchedules( recordsRepositoryId, retentionSchedulesIds ): Observable<any> {

		// get the RecordsRepository
		this.loadHelper( recordsRepositoryId );


	// split on a comma with no spaces
	var idList 					= retentionSchedulesIds.split(',');
	var retentionSchedules 	= this.recordsRepository.retentionSchedules;

	if ( retentionSchedules != null && retentionSchedulesIds != null ) {

		// iterate over array of retentionSchedules ids
		retentionSchedules.forEach(function (obj) {
			if ( retentionSchedulesIds.indexOf(obj._id) > -1 ) {
				// remove the RetentionSchedule
				this.recordsRepository.retentionSchedules.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more legalHoldsIds as a LegalHolds
	// to a RecordsRepository
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLegalHolds( recordsRepositoryId, legalHoldsIds ): Observable<any> {

		// get the RecordsRepository
		this.loadHelper( recordsRepositoryId );

	// split on a comma with no spaces
	var idList = legalHoldsIds.split(',')

	// iterate over array of legalHolds ids
	idList.forEach(function (id) {
		// read the LegalHold
		var legalHold = new LegalHoldService(this.http).getLegalHold(id);
		// add the LegalHold if not already assigned
		if ( this.recordsRepository.legalHolds.indexOf(legalHold) == -1 )
		this.recordsRepository.legalHolds.push(legalHold);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more legalHoldsIds as a LegalHolds
	// from a RecordsRepository
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLegalHolds( recordsRepositoryId, legalHoldsIds ): Observable<any> {

		// get the RecordsRepository
		this.loadHelper( recordsRepositoryId );


	// split on a comma with no spaces
	var idList 					= legalHoldsIds.split(',');
	var legalHolds 	= this.recordsRepository.legalHolds;

	if ( legalHolds != null && legalHoldsIds != null ) {

		// iterate over array of legalHolds ids
		legalHolds.forEach(function (obj) {
			if ( legalHoldsIds.indexOf(obj._id) > -1 ) {
				// remove the LegalHold
				this.recordsRepository.legalHolds.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a RecordsRepository
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/RecordsRepository/update/' + this.recordsRepository;

	return  this.http.post(uri_, this.recordsRepository );
}

	//********************************************************************
	// loadHelper - internal helper to load a RecordsRepository
	//********************************************************************	
	loadHelper( id ) {
		this.getRecordsRepository(id)
			.subscribe((res : RecordsRepository) => {
				this.recordsRepository = res;
			});
	}
}