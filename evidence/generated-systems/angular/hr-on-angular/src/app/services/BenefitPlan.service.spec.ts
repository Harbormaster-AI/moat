import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { BenefitPlanService } from './BenefitPlan.service';

describe('BenefitPlanService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [BenefitPlanService] });
	});

  it('should be created', () => {
    const service: BenefitPlanService = TestBed.get(BenefitPlanService);
    expect(service).toBeTruthy();
  });
});
