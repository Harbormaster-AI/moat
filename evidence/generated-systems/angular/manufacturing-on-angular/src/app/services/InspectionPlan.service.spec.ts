import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InspectionPlanService } from './InspectionPlan.service';

describe('InspectionPlanService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InspectionPlanService] });
	});

  it('should be created', () => {
    const service: InspectionPlanService = TestBed.get(InspectionPlanService);
    expect(service).toBeTruthy();
  });
});
