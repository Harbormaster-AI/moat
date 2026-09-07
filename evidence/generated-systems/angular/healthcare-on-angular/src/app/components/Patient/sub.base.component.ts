import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all Patient Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initAppointmentList();
	this.initEncounterList();
	this.initCarePlanList();
	this.initAllergyList();
	this.initConditionList();
	this.initMedicationOrderList();
	this.initLaboratoryOrderList();
	this.initImagingOrderList();
	this.initCoverageList();
	this.initClaimList();
	this.initMedicalDeviceList();
	this.initObservationList();
  }
}
