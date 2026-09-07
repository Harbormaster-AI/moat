import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { BenefitEnrollmentService } from './BenefitEnrollment.service';

describe('BenefitEnrollmentService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [BenefitEnrollmentService] });
	});

  it('should be created', () => {
    const service: BenefitEnrollmentService = TestBed.get(BenefitEnrollmentService);
    expect(service).toBeTruthy();
  });
});
