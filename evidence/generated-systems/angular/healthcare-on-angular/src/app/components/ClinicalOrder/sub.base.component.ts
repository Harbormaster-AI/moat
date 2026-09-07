import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all ClinicalOrder Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initPatientList();
	this.initEncounterList();
	this.initClinicianList();
	this.initMedicationOrderList();
	this.initLaboratoryOrderList();
	this.initImagingOrderList();
	this.initProcedureOrderList();
	this.initAuthorizationList();
  }
}
