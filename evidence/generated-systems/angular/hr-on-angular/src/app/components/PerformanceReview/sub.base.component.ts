import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all PerformanceReview Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initEmployeeList();
	this.initEmployeeList();
	this.initPerformanceCycleList();
	this.initCompetencyRatingList();
	this.initGoalList();
  }
}
