import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {TypeCertificate} from '../models/TypeCertificate';
import {AircraftProgramService} from '../services/AircraftProgram.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TypeCertificateService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	typeCertificate : TypeCertificate;

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
	// add a TypeCertificate
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTypeCertificate(certificateNumber, authority, Program) : Observable<any> {
		const uri_ = this.apiUrl + '/TypeCertificate/create';
		const obj = {
			      		certificateNumber: certificateNumber,
      		authority: authority,
			Program: Program != null && Program.length > 0 ? Program : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a TypeCertificate
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTypeCertificate(certificateNumber, authority, Program, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/TypeCertificate/update/' + id;
		const obj = {
				      		certificateNumber: certificateNumber,
      		authority: authority,
			Program: Program != null && Program.length > 0 ? Program : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a TypeCertificate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTypeCertificate(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/TypeCertificate/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a TypeCertificate
	// returns the results untouched as an Observable TypeCertificate
	// TypeCertificate model
	// delegates via URI
	//********************************************************************
	getTypeCertificate(id) : Observable<TypeCertificate> {
		const uri_ = this.apiUrl + '/TypeCertificate/load/' + id;

		return this.http.get<TypeCertificate>(uri_);
	}
	
	//********************************************************************
	// gets all TypeCertificate
	// returns the results untouched as JSON representation of an
	// Observable array of TypeCertificate models
	// delegates via URI
	//********************************************************************
	getTypeCertificates() : Observable<TypeCertificate[]> {
		const uri_ = this.apiUrl + '/TypeCertificate/';

		return this
			.http.get<TypeCertificate[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Program on a TypeCertificate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignProgram( typeCertificateId, _programId ): Observable<any> {

		// get the TypeCertificate from storage
		this.loadHelper( typeCertificateId );

	// get the AircraftProgram from storage
	var tmp 	= new AircraftProgramService(this.http).getAircraftProgram(_programId);

	// assign the Program
	this.typeCertificate.program = tmp;

	// save the TypeCertificate
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Program on a TypeCertificate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignProgram( typeCertificateId ): Observable<any> {

		// get the TypeCertificate from storage
		this.loadHelper( typeCertificateId );

	// assign Program to null
	this.typeCertificate.program = null;

	// save the TypeCertificate
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a TypeCertificate
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/TypeCertificate/update/' + this.typeCertificate;

	return  this.http.post(uri_, this.typeCertificate );
}

	//********************************************************************
	// loadHelper - internal helper to load a TypeCertificate
	//********************************************************************	
	loadHelper( id ) {
		this.getTypeCertificate(id)
			.subscribe((res : TypeCertificate) => {
				this.typeCertificate = res;
			});
	}
}