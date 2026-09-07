import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ChargebackService } from './Chargeback.service';

describe('ChargebackService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ChargebackService] });
	});

  it('should be created', () => {
    const service: ChargebackService = TestBed.get(ChargebackService);
    expect(service).toBeTruthy();
  });
});
