import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Patient} from '../models/Patient';
import {AppointmentService} from '../services/Appointment.service';
import {EncounterService} from '../services/Encounter.service';
import {CarePlanService} from '../services/CarePlan.service';
import {AllergyService} from '../services/Allergy.service';
import {ConditionService} from '../services/Condition.service';
import {MedicationOrderService} from '../services/MedicationOrder.service';
import {LaboratoryOrderService} from '../services/LaboratoryOrder.service';
import {ImagingOrderService} from '../services/ImagingOrder.service';
import {CoverageService} from '../services/Coverage.service';
import {ClaimService} from '../services/Claim.service';
import {MedicalDeviceService} from '../services/MedicalDevice.service';
import {ObservationService} from '../services/Observation.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PatientService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	patient : Patient;

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
	// add a Patient
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPatient(firstName, lastName, mrn, dateOfBirth, address, primaryLanguage, Appointments, Encounters, CarePlans, Allergies, Conditions, MedicationOrders, LabOrders, ImagingOrders, Coverages, Claims, Devices, Observations, SexAtBirth, BloodType) : Observable<any> {
		const uri_ = this.apiUrl + '/Patient/create';
		const obj = {
			      		firstName: firstName,
      		lastName: lastName,
      		mrn: mrn,
      		dateOfBirth: dateOfBirth,
      		address: address,
      		primaryLanguage: primaryLanguage,
      		Appointments: Appointments != null && Appointments.length > 0 ? Appointments : null,
      		Encounters: Encounters != null && Encounters.length > 0 ? Encounters : null,
      		CarePlans: CarePlans != null && CarePlans.length > 0 ? CarePlans : null,
      		Allergies: Allergies != null && Allergies.length > 0 ? Allergies : null,
      		Conditions: Conditions != null && Conditions.length > 0 ? Conditions : null,
      		MedicationOrders: MedicationOrders != null && MedicationOrders.length > 0 ? MedicationOrders : null,
      		LabOrders: LabOrders != null && LabOrders.length > 0 ? LabOrders : null,
      		ImagingOrders: ImagingOrders != null && ImagingOrders.length > 0 ? ImagingOrders : null,
      		Coverages: Coverages != null && Coverages.length > 0 ? Coverages : null,
      		Claims: Claims != null && Claims.length > 0 ? Claims : null,
      		Devices: Devices != null && Devices.length > 0 ? Devices : null,
      		Observations: Observations != null && Observations.length > 0 ? Observations : null,
      		SexAtBirth: SexAtBirth,
			BloodType: BloodType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Patient
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePatient(firstName, lastName, mrn, dateOfBirth, address, primaryLanguage, Appointments, Encounters, CarePlans, Allergies, Conditions, MedicationOrders, LabOrders, ImagingOrders, Coverages, Claims, Devices, Observations, SexAtBirth, BloodType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Patient/update/' + id;
		const obj = {
				      		firstName: firstName,
      		lastName: lastName,
      		mrn: mrn,
      		dateOfBirth: dateOfBirth,
      		address: address,
      		primaryLanguage: primaryLanguage,
      		Appointments: Appointments != null && Appointments.length > 0 ? Appointments : null,
      		Encounters: Encounters != null && Encounters.length > 0 ? Encounters : null,
      		CarePlans: CarePlans != null && CarePlans.length > 0 ? CarePlans : null,
      		Allergies: Allergies != null && Allergies.length > 0 ? Allergies : null,
      		Conditions: Conditions != null && Conditions.length > 0 ? Conditions : null,
      		MedicationOrders: MedicationOrders != null && MedicationOrders.length > 0 ? MedicationOrders : null,
      		LabOrders: LabOrders != null && LabOrders.length > 0 ? LabOrders : null,
      		ImagingOrders: ImagingOrders != null && ImagingOrders.length > 0 ? ImagingOrders : null,
      		Coverages: Coverages != null && Coverages.length > 0 ? Coverages : null,
      		Claims: Claims != null && Claims.length > 0 ? Claims : null,
      		Devices: Devices != null && Devices.length > 0 ? Devices : null,
      		Observations: Observations != null && Observations.length > 0 ? Observations : null,
      		SexAtBirth: SexAtBirth,
			BloodType: BloodType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Patient
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePatient(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Patient/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Patient
	// returns the results untouched as an Observable Patient
	// Patient model
	// delegates via URI
	//********************************************************************
	getPatient(id) : Observable<Patient> {
		const uri_ = this.apiUrl + '/Patient/load/' + id;

		return this.http.get<Patient>(uri_);
	}
	
	//********************************************************************
	// gets all Patient
	// returns the results untouched as JSON representation of an
	// Observable array of Patient models
	// delegates via URI
	//********************************************************************
	getPatients() : Observable<Patient[]> {
		const uri_ = this.apiUrl + '/Patient/';

		return this
			.http.get<Patient[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more appointmentsIds as a Appointments
	// to a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAppointments( patientId, appointmentsIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );

	// split on a comma with no spaces
	var idList = appointmentsIds.split(',')

	// iterate over array of appointments ids
	idList.forEach(function (id) {
		// read the Appointment
		var appointment = new AppointmentService(this.http).getAppointment(id);
		// add the Appointment if not already assigned
		if ( this.patient.appointments.indexOf(appointment) == -1 )
		this.patient.appointments.push(appointment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more appointmentsIds as a Appointments
	// from a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAppointments( patientId, appointmentsIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );


	// split on a comma with no spaces
	var idList 					= appointmentsIds.split(',');
	var appointments 	= this.patient.appointments;

	if ( appointments != null && appointmentsIds != null ) {

		// iterate over array of appointments ids
		appointments.forEach(function (obj) {
			if ( appointmentsIds.indexOf(obj._id) > -1 ) {
				// remove the Appointment
				this.patient.appointments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more encountersIds as a Encounters
	// to a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEncounters( patientId, encountersIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );

	// split on a comma with no spaces
	var idList = encountersIds.split(',')

	// iterate over array of encounters ids
	idList.forEach(function (id) {
		// read the Encounter
		var encounter = new EncounterService(this.http).getEncounter(id);
		// add the Encounter if not already assigned
		if ( this.patient.encounters.indexOf(encounter) == -1 )
		this.patient.encounters.push(encounter);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more encountersIds as a Encounters
	// from a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEncounters( patientId, encountersIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );


	// split on a comma with no spaces
	var idList 					= encountersIds.split(',');
	var encounters 	= this.patient.encounters;

	if ( encounters != null && encountersIds != null ) {

		// iterate over array of encounters ids
		encounters.forEach(function (obj) {
			if ( encountersIds.indexOf(obj._id) > -1 ) {
				// remove the Encounter
				this.patient.encounters.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more carePlansIds as a CarePlans
	// to a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCarePlans( patientId, carePlansIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );

	// split on a comma with no spaces
	var idList = carePlansIds.split(',')

	// iterate over array of carePlans ids
	idList.forEach(function (id) {
		// read the CarePlan
		var carePlan = new CarePlanService(this.http).getCarePlan(id);
		// add the CarePlan if not already assigned
		if ( this.patient.carePlans.indexOf(carePlan) == -1 )
		this.patient.carePlans.push(carePlan);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more carePlansIds as a CarePlans
	// from a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCarePlans( patientId, carePlansIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );


	// split on a comma with no spaces
	var idList 					= carePlansIds.split(',');
	var carePlans 	= this.patient.carePlans;

	if ( carePlans != null && carePlansIds != null ) {

		// iterate over array of carePlans ids
		carePlans.forEach(function (obj) {
			if ( carePlansIds.indexOf(obj._id) > -1 ) {
				// remove the CarePlan
				this.patient.carePlans.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more allergiesIds as a Allergies
	// to a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAllergies( patientId, allergiesIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );

	// split on a comma with no spaces
	var idList = allergiesIds.split(',')

	// iterate over array of allergies ids
	idList.forEach(function (id) {
		// read the Allergy
		var allergy = new AllergyService(this.http).getAllergy(id);
		// add the Allergy if not already assigned
		if ( this.patient.allergies.indexOf(allergy) == -1 )
		this.patient.allergies.push(allergy);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more allergiesIds as a Allergies
	// from a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAllergies( patientId, allergiesIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );


	// split on a comma with no spaces
	var idList 					= allergiesIds.split(',');
	var allergies 	= this.patient.allergies;

	if ( allergies != null && allergiesIds != null ) {

		// iterate over array of allergies ids
		allergies.forEach(function (obj) {
			if ( allergiesIds.indexOf(obj._id) > -1 ) {
				// remove the Allergy
				this.patient.allergies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more conditionsIds as a Conditions
	// to a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addConditions( patientId, conditionsIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );

	// split on a comma with no spaces
	var idList = conditionsIds.split(',')

	// iterate over array of conditions ids
	idList.forEach(function (id) {
		// read the Condition
		var condition = new ConditionService(this.http).getCondition(id);
		// add the Condition if not already assigned
		if ( this.patient.conditions.indexOf(condition) == -1 )
		this.patient.conditions.push(condition);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more conditionsIds as a Conditions
	// from a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeConditions( patientId, conditionsIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );


	// split on a comma with no spaces
	var idList 					= conditionsIds.split(',');
	var conditions 	= this.patient.conditions;

	if ( conditions != null && conditionsIds != null ) {

		// iterate over array of conditions ids
		conditions.forEach(function (obj) {
			if ( conditionsIds.indexOf(obj._id) > -1 ) {
				// remove the Condition
				this.patient.conditions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more medicationOrdersIds as a MedicationOrders
	// to a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addMedicationOrders( patientId, medicationOrdersIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );

	// split on a comma with no spaces
	var idList = medicationOrdersIds.split(',')

	// iterate over array of medicationOrders ids
	idList.forEach(function (id) {
		// read the MedicationOrder
		var medicationOrder = new MedicationOrderService(this.http).getMedicationOrder(id);
		// add the MedicationOrder if not already assigned
		if ( this.patient.medicationOrders.indexOf(medicationOrder) == -1 )
		this.patient.medicationOrders.push(medicationOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more medicationOrdersIds as a MedicationOrders
	// from a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeMedicationOrders( patientId, medicationOrdersIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );


	// split on a comma with no spaces
	var idList 					= medicationOrdersIds.split(',');
	var medicationOrders 	= this.patient.medicationOrders;

	if ( medicationOrders != null && medicationOrdersIds != null ) {

		// iterate over array of medicationOrders ids
		medicationOrders.forEach(function (obj) {
			if ( medicationOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the MedicationOrder
				this.patient.medicationOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more labOrdersIds as a LabOrders
	// to a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLabOrders( patientId, labOrdersIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );

	// split on a comma with no spaces
	var idList = labOrdersIds.split(',')

	// iterate over array of labOrders ids
	idList.forEach(function (id) {
		// read the LaboratoryOrder
		var laboratoryOrder = new LaboratoryOrderService(this.http).getLaboratoryOrder(id);
		// add the LaboratoryOrder if not already assigned
		if ( this.patient.labOrders.indexOf(laboratoryOrder) == -1 )
		this.patient.labOrders.push(laboratoryOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more labOrdersIds as a LabOrders
	// from a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLabOrders( patientId, labOrdersIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );


	// split on a comma with no spaces
	var idList 					= labOrdersIds.split(',');
	var labOrders 	= this.patient.labOrders;

	if ( labOrders != null && labOrdersIds != null ) {

		// iterate over array of labOrders ids
		labOrders.forEach(function (obj) {
			if ( labOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the LaboratoryOrder
				this.patient.labOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more imagingOrdersIds as a ImagingOrders
	// to a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addImagingOrders( patientId, imagingOrdersIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );

	// split on a comma with no spaces
	var idList = imagingOrdersIds.split(',')

	// iterate over array of imagingOrders ids
	idList.forEach(function (id) {
		// read the ImagingOrder
		var imagingOrder = new ImagingOrderService(this.http).getImagingOrder(id);
		// add the ImagingOrder if not already assigned
		if ( this.patient.imagingOrders.indexOf(imagingOrder) == -1 )
		this.patient.imagingOrders.push(imagingOrder);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more imagingOrdersIds as a ImagingOrders
	// from a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeImagingOrders( patientId, imagingOrdersIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );


	// split on a comma with no spaces
	var idList 					= imagingOrdersIds.split(',');
	var imagingOrders 	= this.patient.imagingOrders;

	if ( imagingOrders != null && imagingOrdersIds != null ) {

		// iterate over array of imagingOrders ids
		imagingOrders.forEach(function (obj) {
			if ( imagingOrdersIds.indexOf(obj._id) > -1 ) {
				// remove the ImagingOrder
				this.patient.imagingOrders.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more coveragesIds as a Coverages
	// to a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCoverages( patientId, coveragesIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );

	// split on a comma with no spaces
	var idList = coveragesIds.split(',')

	// iterate over array of coverages ids
	idList.forEach(function (id) {
		// read the Coverage
		var coverage = new CoverageService(this.http).getCoverage(id);
		// add the Coverage if not already assigned
		if ( this.patient.coverages.indexOf(coverage) == -1 )
		this.patient.coverages.push(coverage);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more coveragesIds as a Coverages
	// from a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCoverages( patientId, coveragesIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );


	// split on a comma with no spaces
	var idList 					= coveragesIds.split(',');
	var coverages 	= this.patient.coverages;

	if ( coverages != null && coveragesIds != null ) {

		// iterate over array of coverages ids
		coverages.forEach(function (obj) {
			if ( coveragesIds.indexOf(obj._id) > -1 ) {
				// remove the Coverage
				this.patient.coverages.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more claimsIds as a Claims
	// to a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addClaims( patientId, claimsIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );

	// split on a comma with no spaces
	var idList = claimsIds.split(',')

	// iterate over array of claims ids
	idList.forEach(function (id) {
		// read the Claim
		var claim = new ClaimService(this.http).getClaim(id);
		// add the Claim if not already assigned
		if ( this.patient.claims.indexOf(claim) == -1 )
		this.patient.claims.push(claim);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more claimsIds as a Claims
	// from a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeClaims( patientId, claimsIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );


	// split on a comma with no spaces
	var idList 					= claimsIds.split(',');
	var claims 	= this.patient.claims;

	if ( claims != null && claimsIds != null ) {

		// iterate over array of claims ids
		claims.forEach(function (obj) {
			if ( claimsIds.indexOf(obj._id) > -1 ) {
				// remove the Claim
				this.patient.claims.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more devicesIds as a Devices
	// to a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDevices( patientId, devicesIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );

	// split on a comma with no spaces
	var idList = devicesIds.split(',')

	// iterate over array of devices ids
	idList.forEach(function (id) {
		// read the MedicalDevice
		var medicalDevice = new MedicalDeviceService(this.http).getMedicalDevice(id);
		// add the MedicalDevice if not already assigned
		if ( this.patient.devices.indexOf(medicalDevice) == -1 )
		this.patient.devices.push(medicalDevice);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more devicesIds as a Devices
	// from a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDevices( patientId, devicesIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );


	// split on a comma with no spaces
	var idList 					= devicesIds.split(',');
	var devices 	= this.patient.devices;

	if ( devices != null && devicesIds != null ) {

		// iterate over array of devices ids
		devices.forEach(function (obj) {
			if ( devicesIds.indexOf(obj._id) > -1 ) {
				// remove the MedicalDevice
				this.patient.devices.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more observationsIds as a Observations
	// to a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addObservations( patientId, observationsIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );

	// split on a comma with no spaces
	var idList = observationsIds.split(',')

	// iterate over array of observations ids
	idList.forEach(function (id) {
		// read the Observation
		var observation = new ObservationService(this.http).getObservation(id);
		// add the Observation if not already assigned
		if ( this.patient.observations.indexOf(observation) == -1 )
		this.patient.observations.push(observation);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more observationsIds as a Observations
	// from a Patient
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeObservations( patientId, observationsIds ): Observable<any> {

		// get the Patient
		this.loadHelper( patientId );


	// split on a comma with no spaces
	var idList 					= observationsIds.split(',');
	var observations 	= this.patient.observations;

	if ( observations != null && observationsIds != null ) {

		// iterate over array of observations ids
		observations.forEach(function (obj) {
			if ( observationsIds.indexOf(obj._id) > -1 ) {
				// remove the Observation
				this.patient.observations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Patient
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Patient/update/' + this.patient;

	return  this.http.post(uri_, this.patient );
}

	//********************************************************************
	// loadHelper - internal helper to load a Patient
	//********************************************************************	
	loadHelper( id ) {
		this.getPatient(id)
			.subscribe((res : Patient) => {
				this.patient = res;
			});
	}
}