import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PaymentContractService } from './PaymentContract.service';

describe('PaymentContractService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PaymentContractService] });
	});

  it('should be created', () => {
    const service: PaymentContractService = TestBed.get(PaymentContractService);
    expect(service).toBeTruthy();
  });
});
