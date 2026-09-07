import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {RecommendationScenario} from '../models/RecommendationScenario';
import {Model_Service} from '../services/Model_.service';
import {DataSetService} from '../services/DataSet.service';
import {ExperimentService} from '../services/Experiment.service';
import {AlertService} from '../services/Alert.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class RecommendationScenarioService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	recommendationScenario : RecommendationScenario;

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
	// add a RecommendationScenario
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addRecommendationScenario(name, objective, Models, Datasets, Experiments, Alerts, RecommendationType) : Observable<any> {
		const uri_ = this.apiUrl + '/RecommendationScenario/create';
		const obj = {
			      		name: name,
      		objective: objective,
      		Models: Models != null && Models.length > 0 ? Models : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Experiments: Experiments != null && Experiments.length > 0 ? Experiments : null,
      		Alerts: Alerts != null && Alerts.length > 0 ? Alerts : null,
			RecommendationType: RecommendationType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a RecommendationScenario
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateRecommendationScenario(name, objective, Models, Datasets, Experiments, Alerts, RecommendationType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/RecommendationScenario/update/' + id;
		const obj = {
				      		name: name,
      		objective: objective,
      		Models: Models != null && Models.length > 0 ? Models : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Experiments: Experiments != null && Experiments.length > 0 ? Experiments : null,
      		Alerts: Alerts != null && Alerts.length > 0 ? Alerts : null,
			RecommendationType: RecommendationType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a RecommendationScenario
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteRecommendationScenario(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/RecommendationScenario/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a RecommendationScenario
	// returns the results untouched as an Observable RecommendationScenario
	// RecommendationScenario model
	// delegates via URI
	//********************************************************************
	getRecommendationScenario(id) : Observable<RecommendationScenario> {
		const uri_ = this.apiUrl + '/RecommendationScenario/load/' + id;

		return this.http.get<RecommendationScenario>(uri_);
	}
	
	//********************************************************************
	// gets all RecommendationScenario
	// returns the results untouched as JSON representation of an
	// Observable array of RecommendationScenario models
	// delegates via URI
	//********************************************************************
	getRecommendationScenarios() : Observable<RecommendationScenario[]> {
		const uri_ = this.apiUrl + '/RecommendationScenario/';

		return this
			.http.get<RecommendationScenario[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more modelsIds as a Models
	// to a RecommendationScenario
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addModels( recommendationScenarioId, modelsIds ): Observable<any> {

		// get the RecommendationScenario
		this.loadHelper( recommendationScenarioId );

	// split on a comma with no spaces
	var idList = modelsIds.split(',')

	// iterate over array of models ids
	idList.forEach(function (id) {
		// read the Model_
		var model_ = new Model_Service(this.http).getModel_(id);
		// add the Model_ if not already assigned
		if ( this.recommendationScenario.models.indexOf(model_) == -1 )
		this.recommendationScenario.models.push(model_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more modelsIds as a Models
	// from a RecommendationScenario
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeModels( recommendationScenarioId, modelsIds ): Observable<any> {

		// get the RecommendationScenario
		this.loadHelper( recommendationScenarioId );


	// split on a comma with no spaces
	var idList 					= modelsIds.split(',');
	var models 	= this.recommendationScenario.models;

	if ( models != null && modelsIds != null ) {

		// iterate over array of models ids
		models.forEach(function (obj) {
			if ( modelsIds.indexOf(obj._id) > -1 ) {
				// remove the Model_
				this.recommendationScenario.models.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a RecommendationScenario
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( recommendationScenarioId, datasetsIds ): Observable<any> {

		// get the RecommendationScenario
		this.loadHelper( recommendationScenarioId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.recommendationScenario.datasets.indexOf(dataSet) == -1 )
		this.recommendationScenario.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a RecommendationScenario
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( recommendationScenarioId, datasetsIds ): Observable<any> {

		// get the RecommendationScenario
		this.loadHelper( recommendationScenarioId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.recommendationScenario.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.recommendationScenario.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more experimentsIds as a Experiments
	// to a RecommendationScenario
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addExperiments( recommendationScenarioId, experimentsIds ): Observable<any> {

		// get the RecommendationScenario
		this.loadHelper( recommendationScenarioId );

	// split on a comma with no spaces
	var idList = experimentsIds.split(',')

	// iterate over array of experiments ids
	idList.forEach(function (id) {
		// read the Experiment
		var experiment = new ExperimentService(this.http).getExperiment(id);
		// add the Experiment if not already assigned
		if ( this.recommendationScenario.experiments.indexOf(experiment) == -1 )
		this.recommendationScenario.experiments.push(experiment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more experimentsIds as a Experiments
	// from a RecommendationScenario
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeExperiments( recommendationScenarioId, experimentsIds ): Observable<any> {

		// get the RecommendationScenario
		this.loadHelper( recommendationScenarioId );


	// split on a comma with no spaces
	var idList 					= experimentsIds.split(',');
	var experiments 	= this.recommendationScenario.experiments;

	if ( experiments != null && experimentsIds != null ) {

		// iterate over array of experiments ids
		experiments.forEach(function (obj) {
			if ( experimentsIds.indexOf(obj._id) > -1 ) {
				// remove the Experiment
				this.recommendationScenario.experiments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more alertsIds as a Alerts
	// to a RecommendationScenario
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAlerts( recommendationScenarioId, alertsIds ): Observable<any> {

		// get the RecommendationScenario
		this.loadHelper( recommendationScenarioId );

	// split on a comma with no spaces
	var idList = alertsIds.split(',')

	// iterate over array of alerts ids
	idList.forEach(function (id) {
		// read the Alert
		var alert = new AlertService(this.http).getAlert(id);
		// add the Alert if not already assigned
		if ( this.recommendationScenario.alerts.indexOf(alert) == -1 )
		this.recommendationScenario.alerts.push(alert);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more alertsIds as a Alerts
	// from a RecommendationScenario
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAlerts( recommendationScenarioId, alertsIds ): Observable<any> {

		// get the RecommendationScenario
		this.loadHelper( recommendationScenarioId );


	// split on a comma with no spaces
	var idList 					= alertsIds.split(',');
	var alerts 	= this.recommendationScenario.alerts;

	if ( alerts != null && alertsIds != null ) {

		// iterate over array of alerts ids
		alerts.forEach(function (obj) {
			if ( alertsIds.indexOf(obj._id) > -1 ) {
				// remove the Alert
				this.recommendationScenario.alerts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a RecommendationScenario
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/RecommendationScenario/update/' + this.recommendationScenario;

	return  this.http.post(uri_, this.recommendationScenario );
}

	//********************************************************************
	// loadHelper - internal helper to load a RecommendationScenario
	//********************************************************************	
	loadHelper( id ) {
		this.getRecommendationScenario(id)
			.subscribe((res : RecommendationScenario) => {
				this.recommendationScenario = res;
			});
	}
}