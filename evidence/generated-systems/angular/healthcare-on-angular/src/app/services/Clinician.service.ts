import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Clinician} from '../models/Clinician';
import {CareTeamService} from '../services/CareTeam.service';
import {AppointmentService} from '../services/Appointment.service';
import {EncounterService} from '../services/Encounter.service';
import {ProcedureService} from '../services/Procedure.service';
import {ImagingReportService} from '../services/ImagingReport.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ClinicianService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	clinician : Clinician;

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
	// add a Clinician
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addClinician(firstName, lastName, licenseNumber, CareTeams, Appointments, Encounters, Procedures, ImagingReports, ClinicianType, Specialty) : Observable<any> {
		const uri_ = this.apiUrl + '/Clinician/create';
		const obj = {
			      		firstName: firstName,
      		lastName: lastName,
      		licenseNumber: licenseNumber,
      		CareTeams: CareTeams != null && CareTeams.length > 0 ? CareTeams : null,
      		Appointments: Appointments != null && Appointments.length > 0 ? Appointments : null,
      		Encounters: Encounters != null && Encounters.length > 0 ? Encounters : null,
      		Procedures: Procedures != null && Procedures.length > 0 ? Procedures : null,
      		ImagingReports: ImagingReports != null && ImagingReports.length > 0 ? ImagingReports : null,
      		ClinicianType: ClinicianType,
			Specialty: Specialty
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Clinician
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateClinician(firstName, lastName, licenseNumber, CareTeams, Appointments, Encounters, Procedures, ImagingReports, ClinicianType, Specialty, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Clinician/update/' + id;
		const obj = {
				      		firstName: firstName,
      		lastName: lastName,
      		licenseNumber: licenseNumber,
      		CareTeams: CareTeams != null && CareTeams.length > 0 ? CareTeams : null,
      		Appointments: Appointments != null && Appointments.length > 0 ? Appointments : null,
      		Encounters: Encounters != null && Encounters.length > 0 ? Encounters : null,
      		Procedures: Procedures != null && Procedures.length > 0 ? Procedures : null,
      		ImagingReports: ImagingReports != null && ImagingReports.length > 0 ? ImagingReports : null,
      		ClinicianType: ClinicianType,
			Specialty: Specialty
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Clinician
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteClinician(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Clinician/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Clinician
	// returns the results untouched as an Observable Clinician
	// Clinician model
	// delegates via URI
	//********************************************************************
	getClinician(id) : Observable<Clinician> {
		const uri_ = this.apiUrl + '/Clinician/load/' + id;

		return this.http.get<Clinician>(uri_);
	}
	
	//********************************************************************
	// gets all Clinician
	// returns the results untouched as JSON representation of an
	// Observable array of Clinician models
	// delegates via URI
	//********************************************************************
	getClinicians() : Observable<Clinician[]> {
		const uri_ = this.apiUrl + '/Clinician/';

		return this
			.http.get<Clinician[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more careTeamsIds as a CareTeams
	// to a Clinician
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCareTeams( clinicianId, careTeamsIds ): Observable<any> {

		// get the Clinician
		this.loadHelper( clinicianId );

	// split on a comma with no spaces
	var idList = careTeamsIds.split(',')

	// iterate over array of careTeams ids
	idList.forEach(function (id) {
		// read the CareTeam
		var careTeam = new CareTeamService(this.http).getCareTeam(id);
		// add the CareTeam if not already assigned
		if ( this.clinician.careTeams.indexOf(careTeam) == -1 )
		this.clinician.careTeams.push(careTeam);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more careTeamsIds as a CareTeams
	// from a Clinician
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCareTeams( clinicianId, careTeamsIds ): Observable<any> {

		// get the Clinician
		this.loadHelper( clinicianId );


	// split on a comma with no spaces
	var idList 					= careTeamsIds.split(',');
	var careTeams 	= this.clinician.careTeams;

	if ( careTeams != null && careTeamsIds != null ) {

		// iterate over array of careTeams ids
		careTeams.forEach(function (obj) {
			if ( careTeamsIds.indexOf(obj._id) > -1 ) {
				// remove the CareTeam
				this.clinician.careTeams.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more appointmentsIds as a Appointments
	// to a Clinician
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAppointments( clinicianId, appointmentsIds ): Observable<any> {

		// get the Clinician
		this.loadHelper( clinicianId );

	// split on a comma with no spaces
	var idList = appointmentsIds.split(',')

	// iterate over array of appointments ids
	idList.forEach(function (id) {
		// read the Appointment
		var appointment = new AppointmentService(this.http).getAppointment(id);
		// add the Appointment if not already assigned
		if ( this.clinician.appointments.indexOf(appointment) == -1 )
		this.clinician.appointments.push(appointment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more appointmentsIds as a Appointments
	// from a Clinician
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAppointments( clinicianId, appointmentsIds ): Observable<any> {

		// get the Clinician
		this.loadHelper( clinicianId );


	// split on a comma with no spaces
	var idList 					= appointmentsIds.split(',');
	var appointments 	= this.clinician.appointments;

	if ( appointments != null && appointmentsIds != null ) {

		// iterate over array of appointments ids
		appointments.forEach(function (obj) {
			if ( appointmentsIds.indexOf(obj._id) > -1 ) {
				// remove the Appointment
				this.clinician.appointments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more encountersIds as a Encounters
	// to a Clinician
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEncounters( clinicianId, encountersIds ): Observable<any> {

		// get the Clinician
		this.loadHelper( clinicianId );

	// split on a comma with no spaces
	var idList = encountersIds.split(',')

	// iterate over array of encounters ids
	idList.forEach(function (id) {
		// read the Encounter
		var encounter = new EncounterService(this.http).getEncounter(id);
		// add the Encounter if not already assigned
		if ( this.clinician.encounters.indexOf(encounter) == -1 )
		this.clinician.encounters.push(encounter);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more encountersIds as a Encounters
	// from a Clinician
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEncounters( clinicianId, encountersIds ): Observable<any> {

		// get the Clinician
		this.loadHelper( clinicianId );


	// split on a comma with no spaces
	var idList 					= encountersIds.split(',');
	var encounters 	= this.clinician.encounters;

	if ( encounters != null && encountersIds != null ) {

		// iterate over array of encounters ids
		encounters.forEach(function (obj) {
			if ( encountersIds.indexOf(obj._id) > -1 ) {
				// remove the Encounter
				this.clinician.encounters.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more proceduresIds as a Procedures
	// to a Clinician
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addProcedures( clinicianId, proceduresIds ): Observable<any> {

		// get the Clinician
		this.loadHelper( clinicianId );

	// split on a comma with no spaces
	var idList = proceduresIds.split(',')

	// iterate over array of procedures ids
	idList.forEach(function (id) {
		// read the Procedure
		var procedure = new ProcedureService(this.http).getProcedure(id);
		// add the Procedure if not already assigned
		if ( this.clinician.procedures.indexOf(procedure) == -1 )
		this.clinician.procedures.push(procedure);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more proceduresIds as a Procedures
	// from a Clinician
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeProcedures( clinicianId, proceduresIds ): Observable<any> {

		// get the Clinician
		this.loadHelper( clinicianId );


	// split on a comma with no spaces
	var idList 					= proceduresIds.split(',');
	var procedures 	= this.clinician.procedures;

	if ( procedures != null && proceduresIds != null ) {

		// iterate over array of procedures ids
		procedures.forEach(function (obj) {
			if ( proceduresIds.indexOf(obj._id) > -1 ) {
				// remove the Procedure
				this.clinician.procedures.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more imagingReportsIds as a ImagingReports
	// to a Clinician
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addImagingReports( clinicianId, imagingReportsIds ): Observable<any> {

		// get the Clinician
		this.loadHelper( clinicianId );

	// split on a comma with no spaces
	var idList = imagingReportsIds.split(',')

	// iterate over array of imagingReports ids
	idList.forEach(function (id) {
		// read the ImagingReport
		var imagingReport = new ImagingReportService(this.http).getImagingReport(id);
		// add the ImagingReport if not already assigned
		if ( this.clinician.imagingReports.indexOf(imagingReport) == -1 )
		this.clinician.imagingReports.push(imagingReport);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more imagingReportsIds as a ImagingReports
	// from a Clinician
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeImagingReports( clinicianId, imagingReportsIds ): Observable<any> {

		// get the Clinician
		this.loadHelper( clinicianId );


	// split on a comma with no spaces
	var idList 					= imagingReportsIds.split(',');
	var imagingReports 	= this.clinician.imagingReports;

	if ( imagingReports != null && imagingReportsIds != null ) {

		// iterate over array of imagingReports ids
		imagingReports.forEach(function (obj) {
			if ( imagingReportsIds.indexOf(obj._id) > -1 ) {
				// remove the ImagingReport
				this.clinician.imagingReports.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Clinician
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Clinician/update/' + this.clinician;

	return  this.http.post(uri_, this.clinician );
}

	//********************************************************************
	// loadHelper - internal helper to load a Clinician
	//********************************************************************	
	loadHelper( id ) {
		this.getClinician(id)
			.subscribe((res : Clinician) => {
				this.clinician = res;
			});
	}
}