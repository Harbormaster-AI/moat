import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {FraudScenario} from '../models/FraudScenario';
import {Model_Service} from '../services/Model_.service';
import {DataSetService} from '../services/DataSet.service';
import {AlertService} from '../services/Alert.service';
import {FraudSignalService} from '../services/FraudSignal.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class FraudScenarioService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	fraudScenario : FraudScenario;

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
	// add a FraudScenario
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addFraudScenario(name, riskAppetite, Models, Datasets, Alerts, Signals, DetectionType) : Observable<any> {
		const uri_ = this.apiUrl + '/FraudScenario/create';
		const obj = {
			      		name: name,
      		riskAppetite: riskAppetite,
      		Models: Models != null && Models.length > 0 ? Models : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Alerts: Alerts != null && Alerts.length > 0 ? Alerts : null,
      		Signals: Signals != null && Signals.length > 0 ? Signals : null,
			DetectionType: DetectionType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a FraudScenario
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateFraudScenario(name, riskAppetite, Models, Datasets, Alerts, Signals, DetectionType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/FraudScenario/update/' + id;
		const obj = {
				      		name: name,
      		riskAppetite: riskAppetite,
      		Models: Models != null && Models.length > 0 ? Models : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Alerts: Alerts != null && Alerts.length > 0 ? Alerts : null,
      		Signals: Signals != null && Signals.length > 0 ? Signals : null,
			DetectionType: DetectionType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a FraudScenario
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteFraudScenario(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/FraudScenario/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a FraudScenario
	// returns the results untouched as an Observable FraudScenario
	// FraudScenario model
	// delegates via URI
	//********************************************************************
	getFraudScenario(id) : Observable<FraudScenario> {
		const uri_ = this.apiUrl + '/FraudScenario/load/' + id;

		return this.http.get<FraudScenario>(uri_);
	}
	
	//********************************************************************
	// gets all FraudScenario
	// returns the results untouched as JSON representation of an
	// Observable array of FraudScenario models
	// delegates via URI
	//********************************************************************
	getFraudScenarios() : Observable<FraudScenario[]> {
		const uri_ = this.apiUrl + '/FraudScenario/';

		return this
			.http.get<FraudScenario[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more modelsIds as a Models
	// to a FraudScenario
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addModels( fraudScenarioId, modelsIds ): Observable<any> {

		// get the FraudScenario
		this.loadHelper( fraudScenarioId );

	// split on a comma with no spaces
	var idList = modelsIds.split(',')

	// iterate over array of models ids
	idList.forEach(function (id) {
		// read the Model_
		var model_ = new Model_Service(this.http).getModel_(id);
		// add the Model_ if not already assigned
		if ( this.fraudScenario.models.indexOf(model_) == -1 )
		this.fraudScenario.models.push(model_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more modelsIds as a Models
	// from a FraudScenario
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeModels( fraudScenarioId, modelsIds ): Observable<any> {

		// get the FraudScenario
		this.loadHelper( fraudScenarioId );


	// split on a comma with no spaces
	var idList 					= modelsIds.split(',');
	var models 	= this.fraudScenario.models;

	if ( models != null && modelsIds != null ) {

		// iterate over array of models ids
		models.forEach(function (obj) {
			if ( modelsIds.indexOf(obj._id) > -1 ) {
				// remove the Model_
				this.fraudScenario.models.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a FraudScenario
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( fraudScenarioId, datasetsIds ): Observable<any> {

		// get the FraudScenario
		this.loadHelper( fraudScenarioId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.fraudScenario.datasets.indexOf(dataSet) == -1 )
		this.fraudScenario.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a FraudScenario
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( fraudScenarioId, datasetsIds ): Observable<any> {

		// get the FraudScenario
		this.loadHelper( fraudScenarioId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.fraudScenario.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.fraudScenario.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more alertsIds as a Alerts
	// to a FraudScenario
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAlerts( fraudScenarioId, alertsIds ): Observable<any> {

		// get the FraudScenario
		this.loadHelper( fraudScenarioId );

	// split on a comma with no spaces
	var idList = alertsIds.split(',')

	// iterate over array of alerts ids
	idList.forEach(function (id) {
		// read the Alert
		var alert = new AlertService(this.http).getAlert(id);
		// add the Alert if not already assigned
		if ( this.fraudScenario.alerts.indexOf(alert) == -1 )
		this.fraudScenario.alerts.push(alert);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more alertsIds as a Alerts
	// from a FraudScenario
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAlerts( fraudScenarioId, alertsIds ): Observable<any> {

		// get the FraudScenario
		this.loadHelper( fraudScenarioId );


	// split on a comma with no spaces
	var idList 					= alertsIds.split(',');
	var alerts 	= this.fraudScenario.alerts;

	if ( alerts != null && alertsIds != null ) {

		// iterate over array of alerts ids
		alerts.forEach(function (obj) {
			if ( alertsIds.indexOf(obj._id) > -1 ) {
				// remove the Alert
				this.fraudScenario.alerts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more signalsIds as a Signals
	// to a FraudScenario
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSignals( fraudScenarioId, signalsIds ): Observable<any> {

		// get the FraudScenario
		this.loadHelper( fraudScenarioId );

	// split on a comma with no spaces
	var idList = signalsIds.split(',')

	// iterate over array of signals ids
	idList.forEach(function (id) {
		// read the FraudSignal
		var fraudSignal = new FraudSignalService(this.http).getFraudSignal(id);
		// add the FraudSignal if not already assigned
		if ( this.fraudScenario.signals.indexOf(fraudSignal) == -1 )
		this.fraudScenario.signals.push(fraudSignal);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more signalsIds as a Signals
	// from a FraudScenario
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSignals( fraudScenarioId, signalsIds ): Observable<any> {

		// get the FraudScenario
		this.loadHelper( fraudScenarioId );


	// split on a comma with no spaces
	var idList 					= signalsIds.split(',');
	var signals 	= this.fraudScenario.signals;

	if ( signals != null && signalsIds != null ) {

		// iterate over array of signals ids
		signals.forEach(function (obj) {
			if ( signalsIds.indexOf(obj._id) > -1 ) {
				// remove the FraudSignal
				this.fraudScenario.signals.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a FraudScenario
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/FraudScenario/update/' + this.fraudScenario;

	return  this.http.post(uri_, this.fraudScenario );
}

	//********************************************************************
	// loadHelper - internal helper to load a FraudScenario
	//********************************************************************	
	loadHelper( id ) {
		this.getFraudScenario(id)
			.subscribe((res : FraudScenario) => {
				this.fraudScenario = res;
			});
	}
}