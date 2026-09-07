import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ClaimReserveService } from './ClaimReserve.service';

describe('ClaimReserveService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ClaimReserveService] });
	});

  it('should be created', () => {
    const service: ClaimReserveService = TestBed.get(ClaimReserveService);
    expect(service).toBeTruthy();
  });
});
