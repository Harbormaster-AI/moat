import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PricingPlanService } from './PricingPlan.service';

describe('PricingPlanService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PricingPlanService] });
	});

  it('should be created', () => {
    const service: PricingPlanService = TestBed.get(PricingPlanService);
    expect(service).toBeTruthy();
  });
});
