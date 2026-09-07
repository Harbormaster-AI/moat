import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InsuranceProductService } from './InsuranceProduct.service';

describe('InsuranceProductService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InsuranceProductService] });
	});

  it('should be created', () => {
    const service: InsuranceProductService = TestBed.get(InsuranceProductService);
    expect(service).toBeTruthy();
  });
});
