import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ExpirationPolicyService } from './ExpirationPolicy.service';

describe('ExpirationPolicyService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ExpirationPolicyService] });
	});

  it('should be created', () => {
    const service: ExpirationPolicyService = TestBed.get(ExpirationPolicyService);
    expect(service).toBeTruthy();
  });
});
