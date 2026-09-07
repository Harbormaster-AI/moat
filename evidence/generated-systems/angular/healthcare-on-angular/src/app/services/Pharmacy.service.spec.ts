import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PharmacyService } from './Pharmacy.service';

describe('PharmacyService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PharmacyService] });
	});

  it('should be created', () => {
    const service: PharmacyService = TestBed.get(PharmacyService);
    expect(service).toBeTruthy();
  });
});
