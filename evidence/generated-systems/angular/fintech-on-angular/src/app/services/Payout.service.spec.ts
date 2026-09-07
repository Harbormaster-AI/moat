import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PayoutService } from './Payout.service';

describe('PayoutService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PayoutService] });
	});

  it('should be created', () => {
    const service: PayoutService = TestBed.get(PayoutService);
    expect(service).toBeTruthy();
  });
});
