import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Anomaly} from '../models/Anomaly';
import {TimeSeriesService} from '../services/TimeSeries.service';
import {AlertService} from '../services/Alert.service';
import {DataSetService} from '../services/DataSet.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AnomalyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	anomaly : Anomaly;

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
	// add a Anomaly
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAnomaly(occurredAt, details, TimeSeries, Alert, Dataset, AnomalyType, Severity) : Observable<any> {
		const uri_ = this.apiUrl + '/Anomaly/create';
		const obj = {
			      		occurredAt: occurredAt,
      		details: details,
      		TimeSeries: TimeSeries != null && TimeSeries.length > 0 ? TimeSeries : null,
      		Alert: Alert != null && Alert.length > 0 ? Alert : null,
      		Dataset: Dataset != null && Dataset.length > 0 ? Dataset : null,
      		AnomalyType: AnomalyType,
			Severity: Severity
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Anomaly
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAnomaly(occurredAt, details, TimeSeries, Alert, Dataset, AnomalyType, Severity, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Anomaly/update/' + id;
		const obj = {
				      		occurredAt: occurredAt,
      		details: details,
      		TimeSeries: TimeSeries != null && TimeSeries.length > 0 ? TimeSeries : null,
      		Alert: Alert != null && Alert.length > 0 ? Alert : null,
      		Dataset: Dataset != null && Dataset.length > 0 ? Dataset : null,
      		AnomalyType: AnomalyType,
			Severity: Severity
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Anomaly
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAnomaly(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Anomaly/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Anomaly
	// returns the results untouched as an Observable Anomaly
	// Anomaly model
	// delegates via URI
	//********************************************************************
	getAnomaly(id) : Observable<Anomaly> {
		const uri_ = this.apiUrl + '/Anomaly/load/' + id;

		return this.http.get<Anomaly>(uri_);
	}
	
	//********************************************************************
	// gets all Anomaly
	// returns the results untouched as JSON representation of an
	// Observable array of Anomaly models
	// delegates via URI
	//********************************************************************
	getAnomalys() : Observable<Anomaly[]> {
		const uri_ = this.apiUrl + '/Anomaly/';

		return this
			.http.get<Anomaly[]>(uri_);
	}
	
			//********************************************************************
	// assigns a TimeSeries on a Anomaly
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTimeSeries( anomalyId, _timeSeriesId ): Observable<any> {

		// get the Anomaly from storage
		this.loadHelper( anomalyId );

	// get the TimeSeries from storage
	var tmp 	= new TimeSeriesService(this.http).getTimeSeries(_timeSeriesId);

	// assign the TimeSeries
	this.anomaly.timeSeries = tmp;

	// save the Anomaly
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a TimeSeries on a Anomaly
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTimeSeries( anomalyId ): Observable<any> {

		// get the Anomaly from storage
		this.loadHelper( anomalyId );

	// assign TimeSeries to null
	this.anomaly.timeSeries = null;

	// save the Anomaly
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Alert on a Anomaly
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAlert( anomalyId, _alertId ): Observable<any> {

		// get the Anomaly from storage
		this.loadHelper( anomalyId );

	// get the Alert from storage
	var tmp 	= new AlertService(this.http).getAlert(_alertId);

	// assign the Alert
	this.anomaly.alert = tmp;

	// save the Anomaly
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Alert on a Anomaly
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAlert( anomalyId ): Observable<any> {

		// get the Anomaly from storage
		this.loadHelper( anomalyId );

	// assign Alert to null
	this.anomaly.alert = null;

	// save the Anomaly
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Dataset on a Anomaly
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDataset( anomalyId, _datasetId ): Observable<any> {

		// get the Anomaly from storage
		this.loadHelper( anomalyId );

	// get the DataSet from storage
	var tmp 	= new DataSetService(this.http).getDataSet(_datasetId);

	// assign the Dataset
	this.anomaly.dataset = tmp;

	// save the Anomaly
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Dataset on a Anomaly
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDataset( anomalyId ): Observable<any> {

		// get the Anomaly from storage
		this.loadHelper( anomalyId );

	// assign Dataset to null
	this.anomaly.dataset = null;

	// save the Anomaly
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Anomaly
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Anomaly/update/' + this.anomaly;

	return  this.http.post(uri_, this.anomaly );
}

	//********************************************************************
	// loadHelper - internal helper to load a Anomaly
	//********************************************************************	
	loadHelper( id ) {
		this.getAnomaly(id)
			.subscribe((res : Anomaly) => {
				this.anomaly = res;
			});
	}
}