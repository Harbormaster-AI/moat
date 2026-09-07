import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { BrandSafetyPolicyService } from './BrandSafetyPolicy.service';

describe('BrandSafetyPolicyService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [BrandSafetyPolicyService] });
	});

  it('should be created', () => {
    const service: BrandSafetyPolicyService = TestBed.get(BrandSafetyPolicyService);
    expect(service).toBeTruthy();
  });
});
