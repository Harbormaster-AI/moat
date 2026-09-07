import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ReplenishmentPolicyService } from './ReplenishmentPolicy.service';

describe('ReplenishmentPolicyService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ReplenishmentPolicyService] });
	});

  it('should be created', () => {
    const service: ReplenishmentPolicyService = TestBed.get(ReplenishmentPolicyService);
    expect(service).toBeTruthy();
  });
});
