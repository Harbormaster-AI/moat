import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PaymentMethodService } from './PaymentMethod.service';

describe('PaymentMethodService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PaymentMethodService] });
	});

  it('should be created', () => {
    const service: PaymentMethodService = TestBed.get(PaymentMethodService);
    expect(service).toBeTruthy();
  });
});
