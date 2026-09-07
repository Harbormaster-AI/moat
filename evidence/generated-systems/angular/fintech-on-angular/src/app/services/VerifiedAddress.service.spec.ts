import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { VerifiedAddressService } from './VerifiedAddress.service';

describe('VerifiedAddressService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [VerifiedAddressService] });
	});

  it('should be created', () => {
    const service: VerifiedAddressService = TestBed.get(VerifiedAddressService);
    expect(service).toBeTruthy();
  });
});
