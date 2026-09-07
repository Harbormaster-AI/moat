import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InspectionResult} from '../models/InspectionResult';
import {InspectionLotService} from '../services/InspectionLot.service';
import {InspectionCharacteristicService} from '../services/InspectionCharacteristic.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InspectionResultService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	inspectionResult : InspectionResult;

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
	// add a InspectionResult
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInspectionResult(resultValue, recordedOn, notes, InspectionLot, Characteristic, ResultStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/InspectionResult/create';
		const obj = {
			      		resultValue: resultValue,
      		recordedOn: recordedOn,
      		notes: notes,
      		InspectionLot: InspectionLot != null && InspectionLot.length > 0 ? InspectionLot : null,
      		Characteristic: Characteristic != null && Characteristic.length > 0 ? Characteristic : null,
			ResultStatus: ResultStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InspectionResult
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInspectionResult(resultValue, recordedOn, notes, InspectionLot, Characteristic, ResultStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InspectionResult/update/' + id;
		const obj = {
				      		resultValue: resultValue,
      		recordedOn: recordedOn,
      		notes: notes,
      		InspectionLot: InspectionLot != null && InspectionLot.length > 0 ? InspectionLot : null,
      		Characteristic: Characteristic != null && Characteristic.length > 0 ? Characteristic : null,
			ResultStatus: ResultStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InspectionResult
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInspectionResult(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InspectionResult/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InspectionResult
	// returns the results untouched as an Observable InspectionResult
	// InspectionResult model
	// delegates via URI
	//********************************************************************
	getInspectionResult(id) : Observable<InspectionResult> {
		const uri_ = this.apiUrl + '/InspectionResult/load/' + id;

		return this.http.get<InspectionResult>(uri_);
	}
	
	//********************************************************************
	// gets all InspectionResult
	// returns the results untouched as JSON representation of an
	// Observable array of InspectionResult models
	// delegates via URI
	//********************************************************************
	getInspectionResults() : Observable<InspectionResult[]> {
		const uri_ = this.apiUrl + '/InspectionResult/';

		return this
			.http.get<InspectionResult[]>(uri_);
	}
	
			//********************************************************************
	// assigns a InspectionLot on a InspectionResult
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInspectionLot( inspectionResultId, _inspectionLotId ): Observable<any> {

		// get the InspectionResult from storage
		this.loadHelper( inspectionResultId );

	// get the InspectionLot from storage
	var tmp 	= new InspectionLotService(this.http).getInspectionLot(_inspectionLotId);

	// assign the InspectionLot
	this.inspectionResult.inspectionLot = tmp;

	// save the InspectionResult
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a InspectionLot on a InspectionResult
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInspectionLot( inspectionResultId ): Observable<any> {

		// get the InspectionResult from storage
		this.loadHelper( inspectionResultId );

	// assign InspectionLot to null
	this.inspectionResult.inspectionLot = null;

	// save the InspectionResult
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Characteristic on a InspectionResult
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCharacteristic( inspectionResultId, _characteristicId ): Observable<any> {

		// get the InspectionResult from storage
		this.loadHelper( inspectionResultId );

	// get the InspectionCharacteristic from storage
	var tmp 	= new InspectionCharacteristicService(this.http).getInspectionCharacteristic(_characteristicId);

	// assign the Characteristic
	this.inspectionResult.characteristic = tmp;

	// save the InspectionResult
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Characteristic on a InspectionResult
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCharacteristic( inspectionResultId ): Observable<any> {

		// get the InspectionResult from storage
		this.loadHelper( inspectionResultId );

	// assign Characteristic to null
	this.inspectionResult.characteristic = null;

	// save the InspectionResult
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a InspectionResult
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InspectionResult/update/' + this.inspectionResult;

	return  this.http.post(uri_, this.inspectionResult );
}

	//********************************************************************
	// loadHelper - internal helper to load a InspectionResult
	//********************************************************************	
	loadHelper( id ) {
		this.getInspectionResult(id)
			.subscribe((res : InspectionResult) => {
				this.inspectionResult = res;
			});
	}
}