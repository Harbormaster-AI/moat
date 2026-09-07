import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {RiskAssessment} from '../models/RiskAssessment';
import {RiskService} from '../services/Risk.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class RiskAssessmentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	riskAssessment : RiskAssessment;

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
	// add a RiskAssessment
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addRiskAssessment(assessmentDate, assessor, summary, Risk, AssessmentType) : Observable<any> {
		const uri_ = this.apiUrl + '/RiskAssessment/create';
		const obj = {
			      		assessmentDate: assessmentDate,
      		assessor: assessor,
      		summary: summary,
      		Risk: Risk != null && Risk.length > 0 ? Risk : null,
			AssessmentType: AssessmentType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a RiskAssessment
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateRiskAssessment(assessmentDate, assessor, summary, Risk, AssessmentType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/RiskAssessment/update/' + id;
		const obj = {
				      		assessmentDate: assessmentDate,
      		assessor: assessor,
      		summary: summary,
      		Risk: Risk != null && Risk.length > 0 ? Risk : null,
			AssessmentType: AssessmentType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a RiskAssessment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteRiskAssessment(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/RiskAssessment/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a RiskAssessment
	// returns the results untouched as an Observable RiskAssessment
	// RiskAssessment model
	// delegates via URI
	//********************************************************************
	getRiskAssessment(id) : Observable<RiskAssessment> {
		const uri_ = this.apiUrl + '/RiskAssessment/load/' + id;

		return this.http.get<RiskAssessment>(uri_);
	}
	
	//********************************************************************
	// gets all RiskAssessment
	// returns the results untouched as JSON representation of an
	// Observable array of RiskAssessment models
	// delegates via URI
	//********************************************************************
	getRiskAssessments() : Observable<RiskAssessment[]> {
		const uri_ = this.apiUrl + '/RiskAssessment/';

		return this
			.http.get<RiskAssessment[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Risk on a RiskAssessment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRisk( riskAssessmentId, _riskId ): Observable<any> {

		// get the RiskAssessment from storage
		this.loadHelper( riskAssessmentId );

	// get the Risk from storage
	var tmp 	= new RiskService(this.http).getRisk(_riskId);

	// assign the Risk
	this.riskAssessment.risk = tmp;

	// save the RiskAssessment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Risk on a RiskAssessment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRisk( riskAssessmentId ): Observable<any> {

		// get the RiskAssessment from storage
		this.loadHelper( riskAssessmentId );

	// assign Risk to null
	this.riskAssessment.risk = null;

	// save the RiskAssessment
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a RiskAssessment
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/RiskAssessment/update/' + this.riskAssessment;

	return  this.http.post(uri_, this.riskAssessment );
}

	//********************************************************************
	// loadHelper - internal helper to load a RiskAssessment
	//********************************************************************	
	loadHelper( id ) {
		this.getRiskAssessment(id)
			.subscribe((res : RiskAssessment) => {
				this.riskAssessment = res;
			});
	}
}