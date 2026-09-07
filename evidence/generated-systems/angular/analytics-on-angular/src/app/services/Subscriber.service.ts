import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Subscriber} from '../models/Subscriber';
import {AlertService} from '../services/Alert.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class SubscriberService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	subscriber : Subscriber;

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
	// add a Subscriber
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addSubscriber(name, address, Alerts, Channel) : Observable<any> {
		const uri_ = this.apiUrl + '/Subscriber/create';
		const obj = {
			      		name: name,
      		address: address,
      		Alerts: Alerts != null && Alerts.length > 0 ? Alerts : null,
			Channel: Channel
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Subscriber
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateSubscriber(name, address, Alerts, Channel, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Subscriber/update/' + id;
		const obj = {
				      		name: name,
      		address: address,
      		Alerts: Alerts != null && Alerts.length > 0 ? Alerts : null,
			Channel: Channel
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Subscriber
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteSubscriber(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Subscriber/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Subscriber
	// returns the results untouched as an Observable Subscriber
	// Subscriber model
	// delegates via URI
	//********************************************************************
	getSubscriber(id) : Observable<Subscriber> {
		const uri_ = this.apiUrl + '/Subscriber/load/' + id;

		return this.http.get<Subscriber>(uri_);
	}
	
	//********************************************************************
	// gets all Subscriber
	// returns the results untouched as JSON representation of an
	// Observable array of Subscriber models
	// delegates via URI
	//********************************************************************
	getSubscribers() : Observable<Subscriber[]> {
		const uri_ = this.apiUrl + '/Subscriber/';

		return this
			.http.get<Subscriber[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more alertsIds as a Alerts
	// to a Subscriber
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAlerts( subscriberId, alertsIds ): Observable<any> {

		// get the Subscriber
		this.loadHelper( subscriberId );

	// split on a comma with no spaces
	var idList = alertsIds.split(',')

	// iterate over array of alerts ids
	idList.forEach(function (id) {
		// read the Alert
		var alert = new AlertService(this.http).getAlert(id);
		// add the Alert if not already assigned
		if ( this.subscriber.alerts.indexOf(alert) == -1 )
		this.subscriber.alerts.push(alert);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more alertsIds as a Alerts
	// from a Subscriber
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAlerts( subscriberId, alertsIds ): Observable<any> {

		// get the Subscriber
		this.loadHelper( subscriberId );


	// split on a comma with no spaces
	var idList 					= alertsIds.split(',');
	var alerts 	= this.subscriber.alerts;

	if ( alerts != null && alertsIds != null ) {

		// iterate over array of alerts ids
		alerts.forEach(function (obj) {
			if ( alertsIds.indexOf(obj._id) > -1 ) {
				// remove the Alert
				this.subscriber.alerts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Subscriber
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Subscriber/update/' + this.subscriber;

	return  this.http.post(uri_, this.subscriber );
}

	//********************************************************************
	// loadHelper - internal helper to load a Subscriber
	//********************************************************************	
	loadHelper( id ) {
		this.getSubscriber(id)
			.subscribe((res : Subscriber) => {
				this.subscriber = res;
			});
	}
}