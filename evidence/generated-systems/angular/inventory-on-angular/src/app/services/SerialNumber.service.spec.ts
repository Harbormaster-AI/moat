import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { SerialNumberService } from './SerialNumber.service';

describe('SerialNumberService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [SerialNumberService] });
	});

  it('should be created', () => {
    const service: SerialNumberService = TestBed.get(SerialNumberService);
    expect(service).toBeTruthy();
  });
});
