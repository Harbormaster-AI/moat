import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all Dashboard Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initAnalyticsWorkspaceList();
	this.initVisualizationList();
	this.initReportList();
	this.initDataSetList();
	this.initAlertList();
	this.initBIQueryList();
	this.initTagList();
  }
}
