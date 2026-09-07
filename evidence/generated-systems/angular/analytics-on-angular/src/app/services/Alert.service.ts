import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Alert} from '../models/Alert';
import {MetricService} from '../services/Metric.service';
import {DashboardService} from '../services/Dashboard.service';
import {DataSetService} from '../services/DataSet.service';
import {QualityRuleService} from '../services/QualityRule.service';
import {AnomalyService} from '../services/Anomaly.service';
import {SubscriberService} from '../services/Subscriber.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AlertService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	alert : Alert;

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
	// add a Alert
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAlert(title, createdAt, Metric, Dashboard, Dataset, Rule, Anomalies, Subscribers, Severity, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Alert/create';
		const obj = {
			      		title: title,
      		createdAt: createdAt,
      		Metric: Metric != null && Metric.length > 0 ? Metric : null,
      		Dashboard: Dashboard != null && Dashboard.length > 0 ? Dashboard : null,
      		Dataset: Dataset != null && Dataset.length > 0 ? Dataset : null,
      		Rule: Rule != null && Rule.length > 0 ? Rule : null,
      		Anomalies: Anomalies != null && Anomalies.length > 0 ? Anomalies : null,
      		Subscribers: Subscribers != null && Subscribers.length > 0 ? Subscribers : null,
      		Severity: Severity,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Alert
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAlert(title, createdAt, Metric, Dashboard, Dataset, Rule, Anomalies, Subscribers, Severity, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Alert/update/' + id;
		const obj = {
				      		title: title,
      		createdAt: createdAt,
      		Metric: Metric != null && Metric.length > 0 ? Metric : null,
      		Dashboard: Dashboard != null && Dashboard.length > 0 ? Dashboard : null,
      		Dataset: Dataset != null && Dataset.length > 0 ? Dataset : null,
      		Rule: Rule != null && Rule.length > 0 ? Rule : null,
      		Anomalies: Anomalies != null && Anomalies.length > 0 ? Anomalies : null,
      		Subscribers: Subscribers != null && Subscribers.length > 0 ? Subscribers : null,
      		Severity: Severity,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Alert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAlert(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Alert/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Alert
	// returns the results untouched as an Observable Alert
	// Alert model
	// delegates via URI
	//********************************************************************
	getAlert(id) : Observable<Alert> {
		const uri_ = this.apiUrl + '/Alert/load/' + id;

		return this.http.get<Alert>(uri_);
	}
	
	//********************************************************************
	// gets all Alert
	// returns the results untouched as JSON representation of an
	// Observable array of Alert models
	// delegates via URI
	//********************************************************************
	getAlerts() : Observable<Alert[]> {
		const uri_ = this.apiUrl + '/Alert/';

		return this
			.http.get<Alert[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Metric on a Alert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignMetric( alertId, _metricId ): Observable<any> {

		// get the Alert from storage
		this.loadHelper( alertId );

	// get the Metric from storage
	var tmp 	= new MetricService(this.http).getMetric(_metricId);

	// assign the Metric
	this.alert.metric = tmp;

	// save the Alert
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Metric on a Alert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignMetric( alertId ): Observable<any> {

		// get the Alert from storage
		this.loadHelper( alertId );

	// assign Metric to null
	this.alert.metric = null;

	// save the Alert
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Dashboard on a Alert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDashboard( alertId, _dashboardId ): Observable<any> {

		// get the Alert from storage
		this.loadHelper( alertId );

	// get the Dashboard from storage
	var tmp 	= new DashboardService(this.http).getDashboard(_dashboardId);

	// assign the Dashboard
	this.alert.dashboard = tmp;

	// save the Alert
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Dashboard on a Alert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDashboard( alertId ): Observable<any> {

		// get the Alert from storage
		this.loadHelper( alertId );

	// assign Dashboard to null
	this.alert.dashboard = null;

	// save the Alert
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Dataset on a Alert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDataset( alertId, _datasetId ): Observable<any> {

		// get the Alert from storage
		this.loadHelper( alertId );

	// get the DataSet from storage
	var tmp 	= new DataSetService(this.http).getDataSet(_datasetId);

	// assign the Dataset
	this.alert.dataset = tmp;

	// save the Alert
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Dataset on a Alert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDataset( alertId ): Observable<any> {

		// get the Alert from storage
		this.loadHelper( alertId );

	// assign Dataset to null
	this.alert.dataset = null;

	// save the Alert
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Rule on a Alert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRule( alertId, _ruleId ): Observable<any> {

		// get the Alert from storage
		this.loadHelper( alertId );

	// get the QualityRule from storage
	var tmp 	= new QualityRuleService(this.http).getQualityRule(_ruleId);

	// assign the Rule
	this.alert.rule = tmp;

	// save the Alert
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Rule on a Alert
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRule( alertId ): Observable<any> {

		// get the Alert from storage
		this.loadHelper( alertId );

	// assign Rule to null
	this.alert.rule = null;

	// save the Alert
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more anomaliesIds as a Anomalies
	// to a Alert
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAnomalies( alertId, anomaliesIds ): Observable<any> {

		// get the Alert
		this.loadHelper( alertId );

	// split on a comma with no spaces
	var idList = anomaliesIds.split(',')

	// iterate over array of anomalies ids
	idList.forEach(function (id) {
		// read the Anomaly
		var anomaly = new AnomalyService(this.http).getAnomaly(id);
		// add the Anomaly if not already assigned
		if ( this.alert.anomalies.indexOf(anomaly) == -1 )
		this.alert.anomalies.push(anomaly);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more anomaliesIds as a Anomalies
	// from a Alert
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAnomalies( alertId, anomaliesIds ): Observable<any> {

		// get the Alert
		this.loadHelper( alertId );


	// split on a comma with no spaces
	var idList 					= anomaliesIds.split(',');
	var anomalies 	= this.alert.anomalies;

	if ( anomalies != null && anomaliesIds != null ) {

		// iterate over array of anomalies ids
		anomalies.forEach(function (obj) {
			if ( anomaliesIds.indexOf(obj._id) > -1 ) {
				// remove the Anomaly
				this.alert.anomalies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more subscribersIds as a Subscribers
	// to a Alert
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSubscribers( alertId, subscribersIds ): Observable<any> {

		// get the Alert
		this.loadHelper( alertId );

	// split on a comma with no spaces
	var idList = subscribersIds.split(',')

	// iterate over array of subscribers ids
	idList.forEach(function (id) {
		// read the Subscriber
		var subscriber = new SubscriberService(this.http).getSubscriber(id);
		// add the Subscriber if not already assigned
		if ( this.alert.subscribers.indexOf(subscriber) == -1 )
		this.alert.subscribers.push(subscriber);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more subscribersIds as a Subscribers
	// from a Alert
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSubscribers( alertId, subscribersIds ): Observable<any> {

		// get the Alert
		this.loadHelper( alertId );


	// split on a comma with no spaces
	var idList 					= subscribersIds.split(',');
	var subscribers 	= this.alert.subscribers;

	if ( subscribers != null && subscribersIds != null ) {

		// iterate over array of subscribers ids
		subscribers.forEach(function (obj) {
			if ( subscribersIds.indexOf(obj._id) > -1 ) {
				// remove the Subscriber
				this.alert.subscribers.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Alert
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Alert/update/' + this.alert;

	return  this.http.post(uri_, this.alert );
}

	//********************************************************************
	// loadHelper - internal helper to load a Alert
	//********************************************************************	
	loadHelper( id ) {
		this.getAlert(id)
			.subscribe((res : Alert) => {
				this.alert = res;
			});
	}
}