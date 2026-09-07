import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {FraudSignal} from '../models/FraudSignal';
import {FraudScenarioService} from '../services/FraudScenario.service';
import {DataSetService} from '../services/DataSet.service';
import {ModelVersionService} from '../services/ModelVersion.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class FraudSignalService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	fraudSignal : FraudSignal;

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
	// add a FraudSignal
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addFraudSignal(name, ruleLogic, Scenario, Dataset, ModelVersion, SignalType) : Observable<any> {
		const uri_ = this.apiUrl + '/FraudSignal/create';
		const obj = {
			      		name: name,
      		ruleLogic: ruleLogic,
      		Scenario: Scenario != null && Scenario.length > 0 ? Scenario : null,
      		Dataset: Dataset != null && Dataset.length > 0 ? Dataset : null,
      		ModelVersion: ModelVersion != null && ModelVersion.length > 0 ? ModelVersion : null,
			SignalType: SignalType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a FraudSignal
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateFraudSignal(name, ruleLogic, Scenario, Dataset, ModelVersion, SignalType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/FraudSignal/update/' + id;
		const obj = {
				      		name: name,
      		ruleLogic: ruleLogic,
      		Scenario: Scenario != null && Scenario.length > 0 ? Scenario : null,
      		Dataset: Dataset != null && Dataset.length > 0 ? Dataset : null,
      		ModelVersion: ModelVersion != null && ModelVersion.length > 0 ? ModelVersion : null,
			SignalType: SignalType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a FraudSignal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteFraudSignal(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/FraudSignal/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a FraudSignal
	// returns the results untouched as an Observable FraudSignal
	// FraudSignal model
	// delegates via URI
	//********************************************************************
	getFraudSignal(id) : Observable<FraudSignal> {
		const uri_ = this.apiUrl + '/FraudSignal/load/' + id;

		return this.http.get<FraudSignal>(uri_);
	}
	
	//********************************************************************
	// gets all FraudSignal
	// returns the results untouched as JSON representation of an
	// Observable array of FraudSignal models
	// delegates via URI
	//********************************************************************
	getFraudSignals() : Observable<FraudSignal[]> {
		const uri_ = this.apiUrl + '/FraudSignal/';

		return this
			.http.get<FraudSignal[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Scenario on a FraudSignal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignScenario( fraudSignalId, _scenarioId ): Observable<any> {

		// get the FraudSignal from storage
		this.loadHelper( fraudSignalId );

	// get the FraudScenario from storage
	var tmp 	= new FraudScenarioService(this.http).getFraudScenario(_scenarioId);

	// assign the Scenario
	this.fraudSignal.scenario = tmp;

	// save the FraudSignal
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Scenario on a FraudSignal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignScenario( fraudSignalId ): Observable<any> {

		// get the FraudSignal from storage
		this.loadHelper( fraudSignalId );

	// assign Scenario to null
	this.fraudSignal.scenario = null;

	// save the FraudSignal
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Dataset on a FraudSignal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDataset( fraudSignalId, _datasetId ): Observable<any> {

		// get the FraudSignal from storage
		this.loadHelper( fraudSignalId );

	// get the DataSet from storage
	var tmp 	= new DataSetService(this.http).getDataSet(_datasetId);

	// assign the Dataset
	this.fraudSignal.dataset = tmp;

	// save the FraudSignal
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Dataset on a FraudSignal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDataset( fraudSignalId ): Observable<any> {

		// get the FraudSignal from storage
		this.loadHelper( fraudSignalId );

	// assign Dataset to null
	this.fraudSignal.dataset = null;

	// save the FraudSignal
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ModelVersion on a FraudSignal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignModelVersion( fraudSignalId, _modelVersionId ): Observable<any> {

		// get the FraudSignal from storage
		this.loadHelper( fraudSignalId );

	// get the ModelVersion from storage
	var tmp 	= new ModelVersionService(this.http).getModelVersion(_modelVersionId);

	// assign the ModelVersion
	this.fraudSignal.modelVersion = tmp;

	// save the FraudSignal
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ModelVersion on a FraudSignal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignModelVersion( fraudSignalId ): Observable<any> {

		// get the FraudSignal from storage
		this.loadHelper( fraudSignalId );

	// assign ModelVersion to null
	this.fraudSignal.modelVersion = null;

	// save the FraudSignal
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a FraudSignal
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/FraudSignal/update/' + this.fraudSignal;

	return  this.http.post(uri_, this.fraudSignal );
}

	//********************************************************************
	// loadHelper - internal helper to load a FraudSignal
	//********************************************************************	
	loadHelper( id ) {
		this.getFraudSignal(id)
			.subscribe((res : FraudSignal) => {
				this.fraudSignal = res;
			});
	}
}