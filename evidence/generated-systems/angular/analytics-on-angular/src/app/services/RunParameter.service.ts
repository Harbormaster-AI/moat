import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {RunParameter} from '../models/RunParameter';
import {TrainingRunService} from '../services/TrainingRun.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class RunParameterService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	runParameter : RunParameter;

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
	// add a RunParameter
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addRunParameter(name, value, TrainingRun) : Observable<any> {
		const uri_ = this.apiUrl + '/RunParameter/create';
		const obj = {
			      		name: name,
      		value: value,
			TrainingRun: TrainingRun != null && TrainingRun.length > 0 ? TrainingRun : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a RunParameter
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateRunParameter(name, value, TrainingRun, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/RunParameter/update/' + id;
		const obj = {
				      		name: name,
      		value: value,
			TrainingRun: TrainingRun != null && TrainingRun.length > 0 ? TrainingRun : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a RunParameter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteRunParameter(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/RunParameter/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a RunParameter
	// returns the results untouched as an Observable RunParameter
	// RunParameter model
	// delegates via URI
	//********************************************************************
	getRunParameter(id) : Observable<RunParameter> {
		const uri_ = this.apiUrl + '/RunParameter/load/' + id;

		return this.http.get<RunParameter>(uri_);
	}
	
	//********************************************************************
	// gets all RunParameter
	// returns the results untouched as JSON representation of an
	// Observable array of RunParameter models
	// delegates via URI
	//********************************************************************
	getRunParameters() : Observable<RunParameter[]> {
		const uri_ = this.apiUrl + '/RunParameter/';

		return this
			.http.get<RunParameter[]>(uri_);
	}
	
			//********************************************************************
	// assigns a TrainingRun on a RunParameter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTrainingRun( runParameterId, _trainingRunId ): Observable<any> {

		// get the RunParameter from storage
		this.loadHelper( runParameterId );

	// get the TrainingRun from storage
	var tmp 	= new TrainingRunService(this.http).getTrainingRun(_trainingRunId);

	// assign the TrainingRun
	this.runParameter.trainingRun = tmp;

	// save the RunParameter
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a TrainingRun on a RunParameter
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTrainingRun( runParameterId ): Observable<any> {

		// get the RunParameter from storage
		this.loadHelper( runParameterId );

	// assign TrainingRun to null
	this.runParameter.trainingRun = null;

	// save the RunParameter
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a RunParameter
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/RunParameter/update/' + this.runParameter;

	return  this.http.post(uri_, this.runParameter );
}

	//********************************************************************
	// loadHelper - internal helper to load a RunParameter
	//********************************************************************	
	loadHelper( id ) {
		this.getRunParameter(id)
			.subscribe((res : RunParameter) => {
				this.runParameter = res;
			});
	}
}