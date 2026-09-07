import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ThirdPartyAssessmentService } from './ThirdPartyAssessment.service';

describe('ThirdPartyAssessmentService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ThirdPartyAssessmentService] });
	});

  it('should be created', () => {
    const service: ThirdPartyAssessmentService = TestBed.get(ThirdPartyAssessmentService);
    expect(service).toBeTruthy();
  });
});
