import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PaymentOrderService } from './PaymentOrder.service';

describe('PaymentOrderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PaymentOrderService] });
	});

  it('should be created', () => {
    const service: PaymentOrderService = TestBed.get(PaymentOrderService);
    expect(service).toBeTruthy();
  });
});
