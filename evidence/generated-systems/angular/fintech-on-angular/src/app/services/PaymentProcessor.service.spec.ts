import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PaymentProcessorService } from './PaymentProcessor.service';

describe('PaymentProcessorService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PaymentProcessorService] });
	});

  it('should be created', () => {
    const service: PaymentProcessorService = TestBed.get(PaymentProcessorService);
    expect(service).toBeTruthy();
  });
});
