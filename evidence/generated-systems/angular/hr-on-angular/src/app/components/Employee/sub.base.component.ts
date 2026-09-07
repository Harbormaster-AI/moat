import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all Employee Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initEmployeeList();
	this.initEmployeeList();
	this.initDepartmentList();
	this.initLocationList();
	this.initCostCenterList();
	this.initEmploymentAssignmentList();
	this.initEmploymentContractList();
	this.initBenefitEnrollmentList();
	this.initTimesheetList();
	this.initLeaveRequestList();
	this.initPerformanceReviewList();
	this.initTrainingEnrollmentList();
	this.initWorkAuthorizationList();
  }
}
