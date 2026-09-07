import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all LineageNode Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initAnalyticsWorkspaceList();
	this.initLineageNodeList();
	this.initLineageNodeList();
	this.initDataSetList();
	this.initModel_List();
	this.initDataPipelineList();
	this.initDashboardList();
	this.initReportList();
  }
}
