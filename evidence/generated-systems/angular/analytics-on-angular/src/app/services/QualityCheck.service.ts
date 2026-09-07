import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {QualityCheck} from '../models/QualityCheck';
import {QualityRuleService} from '../services/QualityRule.service';
import {DataSetService} from '../services/DataSet.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class QualityCheckService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	qualityCheck : QualityCheck;

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
	// add a QualityCheck
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addQualityCheck(checkedAt, observedValue, sampleSize, Rule, Dataset, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/QualityCheck/create';
		const obj = {
			      		checkedAt: checkedAt,
      		observedValue: observedValue,
      		sampleSize: sampleSize,
      		Rule: Rule != null && Rule.length > 0 ? Rule : null,
      		Dataset: Dataset != null && Dataset.length > 0 ? Dataset : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a QualityCheck
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateQualityCheck(checkedAt, observedValue, sampleSize, Rule, Dataset, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/QualityCheck/update/' + id;
		const obj = {
				      		checkedAt: checkedAt,
      		observedValue: observedValue,
      		sampleSize: sampleSize,
      		Rule: Rule != null && Rule.length > 0 ? Rule : null,
      		Dataset: Dataset != null && Dataset.length > 0 ? Dataset : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a QualityCheck
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteQualityCheck(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/QualityCheck/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a QualityCheck
	// returns the results untouched as an Observable QualityCheck
	// QualityCheck model
	// delegates via URI
	//********************************************************************
	getQualityCheck(id) : Observable<QualityCheck> {
		const uri_ = this.apiUrl + '/QualityCheck/load/' + id;

		return this.http.get<QualityCheck>(uri_);
	}
	
	//********************************************************************
	// gets all QualityCheck
	// returns the results untouched as JSON representation of an
	// Observable array of QualityCheck models
	// delegates via URI
	//********************************************************************
	getQualityChecks() : Observable<QualityCheck[]> {
		const uri_ = this.apiUrl + '/QualityCheck/';

		return this
			.http.get<QualityCheck[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Rule on a QualityCheck
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRule( qualityCheckId, _ruleId ): Observable<any> {

		// get the QualityCheck from storage
		this.loadHelper( qualityCheckId );

	// get the QualityRule from storage
	var tmp 	= new QualityRuleService(this.http).getQualityRule(_ruleId);

	// assign the Rule
	this.qualityCheck.rule = tmp;

	// save the QualityCheck
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Rule on a QualityCheck
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRule( qualityCheckId ): Observable<any> {

		// get the QualityCheck from storage
		this.loadHelper( qualityCheckId );

	// assign Rule to null
	this.qualityCheck.rule = null;

	// save the QualityCheck
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Dataset on a QualityCheck
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDataset( qualityCheckId, _datasetId ): Observable<any> {

		// get the QualityCheck from storage
		this.loadHelper( qualityCheckId );

	// get the DataSet from storage
	var tmp 	= new DataSetService(this.http).getDataSet(_datasetId);

	// assign the Dataset
	this.qualityCheck.dataset = tmp;

	// save the QualityCheck
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Dataset on a QualityCheck
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDataset( qualityCheckId ): Observable<any> {

		// get the QualityCheck from storage
		this.loadHelper( qualityCheckId );

	// assign Dataset to null
	this.qualityCheck.dataset = null;

	// save the QualityCheck
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a QualityCheck
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/QualityCheck/update/' + this.qualityCheck;

	return  this.http.post(uri_, this.qualityCheck );
}

	//********************************************************************
	// loadHelper - internal helper to load a QualityCheck
	//********************************************************************	
	loadHelper( id ) {
		this.getQualityCheck(id)
			.subscribe((res : QualityCheck) => {
				this.qualityCheck = res;
			});
	}
}