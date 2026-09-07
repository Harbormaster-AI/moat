import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ClaimPaymentService } from './ClaimPayment.service';

describe('ClaimPaymentService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ClaimPaymentService] });
	});

  it('should be created', () => {
    const service: ClaimPaymentService = TestBed.get(ClaimPaymentService);
    expect(service).toBeTruthy();
  });
});
