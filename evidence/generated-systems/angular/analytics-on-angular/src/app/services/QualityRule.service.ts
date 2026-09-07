import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {QualityRule} from '../models/QualityRule';
import {DataSetService} from '../services/DataSet.service';
import {QualityCheckService} from '../services/QualityCheck.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class QualityRuleService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	qualityRule : QualityRule;

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
	// add a QualityRule
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addQualityRule(name, threshold, targetField, Dataset, Checks, Dimension, Operator) : Observable<any> {
		const uri_ = this.apiUrl + '/QualityRule/create';
		const obj = {
			      		name: name,
      		threshold: threshold,
      		targetField: targetField,
      		Dataset: Dataset != null && Dataset.length > 0 ? Dataset : null,
      		Checks: Checks != null && Checks.length > 0 ? Checks : null,
      		Dimension: Dimension,
			Operator: Operator
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a QualityRule
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateQualityRule(name, threshold, targetField, Dataset, Checks, Dimension, Operator, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/QualityRule/update/' + id;
		const obj = {
				      		name: name,
      		threshold: threshold,
      		targetField: targetField,
      		Dataset: Dataset != null && Dataset.length > 0 ? Dataset : null,
      		Checks: Checks != null && Checks.length > 0 ? Checks : null,
      		Dimension: Dimension,
			Operator: Operator
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a QualityRule
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteQualityRule(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/QualityRule/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a QualityRule
	// returns the results untouched as an Observable QualityRule
	// QualityRule model
	// delegates via URI
	//********************************************************************
	getQualityRule(id) : Observable<QualityRule> {
		const uri_ = this.apiUrl + '/QualityRule/load/' + id;

		return this.http.get<QualityRule>(uri_);
	}
	
	//********************************************************************
	// gets all QualityRule
	// returns the results untouched as JSON representation of an
	// Observable array of QualityRule models
	// delegates via URI
	//********************************************************************
	getQualityRules() : Observable<QualityRule[]> {
		const uri_ = this.apiUrl + '/QualityRule/';

		return this
			.http.get<QualityRule[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Dataset on a QualityRule
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDataset( qualityRuleId, _datasetId ): Observable<any> {

		// get the QualityRule from storage
		this.loadHelper( qualityRuleId );

	// get the DataSet from storage
	var tmp 	= new DataSetService(this.http).getDataSet(_datasetId);

	// assign the Dataset
	this.qualityRule.dataset = tmp;

	// save the QualityRule
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Dataset on a QualityRule
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDataset( qualityRuleId ): Observable<any> {

		// get the QualityRule from storage
		this.loadHelper( qualityRuleId );

	// assign Dataset to null
	this.qualityRule.dataset = null;

	// save the QualityRule
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more checksIds as a Checks
	// to a QualityRule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addChecks( qualityRuleId, checksIds ): Observable<any> {

		// get the QualityRule
		this.loadHelper( qualityRuleId );

	// split on a comma with no spaces
	var idList = checksIds.split(',')

	// iterate over array of checks ids
	idList.forEach(function (id) {
		// read the QualityCheck
		var qualityCheck = new QualityCheckService(this.http).getQualityCheck(id);
		// add the QualityCheck if not already assigned
		if ( this.qualityRule.checks.indexOf(qualityCheck) == -1 )
		this.qualityRule.checks.push(qualityCheck);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more checksIds as a Checks
	// from a QualityRule
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeChecks( qualityRuleId, checksIds ): Observable<any> {

		// get the QualityRule
		this.loadHelper( qualityRuleId );


	// split on a comma with no spaces
	var idList 					= checksIds.split(',');
	var checks 	= this.qualityRule.checks;

	if ( checks != null && checksIds != null ) {

		// iterate over array of checks ids
		checks.forEach(function (obj) {
			if ( checksIds.indexOf(obj._id) > -1 ) {
				// remove the QualityCheck
				this.qualityRule.checks.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a QualityRule
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/QualityRule/update/' + this.qualityRule;

	return  this.http.post(uri_, this.qualityRule );
}

	//********************************************************************
	// loadHelper - internal helper to load a QualityRule
	//********************************************************************	
	loadHelper( id ) {
		this.getQualityRule(id)
			.subscribe((res : QualityRule) => {
				this.qualityRule = res;
			});
	}
}