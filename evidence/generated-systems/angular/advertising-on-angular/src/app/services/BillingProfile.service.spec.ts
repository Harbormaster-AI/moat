import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { BillingProfileService } from './BillingProfile.service';

describe('BillingProfileService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [BillingProfileService] });
	});

  it('should be created', () => {
    const service: BillingProfileService = TestBed.get(BillingProfileService);
    expect(service).toBeTruthy();
  });
});
