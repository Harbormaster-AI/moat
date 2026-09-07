import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AppliedFeeService } from './AppliedFee.service';

describe('AppliedFeeService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AppliedFeeService] });
	});

  it('should be created', () => {
    const service: AppliedFeeService = TestBed.get(AppliedFeeService);
    expect(service).toBeTruthy();
  });
});
