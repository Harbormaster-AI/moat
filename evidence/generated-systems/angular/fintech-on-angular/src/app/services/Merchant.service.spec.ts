import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { MerchantService } from './Merchant.service';

describe('MerchantService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [MerchantService] });
	});

  it('should be created', () => {
    const service: MerchantService = TestBed.get(MerchantService);
    expect(service).toBeTruthy();
  });
});
