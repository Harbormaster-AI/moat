import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all ModelVersion Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initModel_List();
	this.initTrainingRunList();
	this.initEvaluationMetricList();
	this.initInferenceEndpointList();
	this.initFeatureSetList();
	this.initDataSetList();
  }
}
