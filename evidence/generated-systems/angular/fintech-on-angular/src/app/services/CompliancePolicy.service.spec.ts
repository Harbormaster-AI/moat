import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CompliancePolicyService } from './CompliancePolicy.service';

describe('CompliancePolicyService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CompliancePolicyService] });
	});

  it('should be created', () => {
    const service: CompliancePolicyService = TestBed.get(CompliancePolicyService);
    expect(service).toBeTruthy();
  });
});
