import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { BillingAccountService } from './BillingAccount.service';

describe('BillingAccountService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [BillingAccountService] });
	});

  it('should be created', () => {
    const service: BillingAccountService = TestBed.get(BillingAccountService);
    expect(service).toBeTruthy();
  });
});
