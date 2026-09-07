import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all Control Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initPolicyList();
	this.initControlTest_List();
	this.initEvidenceList();
	this.initRiskList();
	this.initObligationList();
	this.initProcedureList();
	this.initIssueList();
  }
}
