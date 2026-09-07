import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all AuditFinding Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initAuditEngagementList();
	this.initAuditWorkpaperList();
	this.initCorrectiveActionList();
	this.initRiskList();
	this.initControlList();
	this.initIssueList();
  }
}
