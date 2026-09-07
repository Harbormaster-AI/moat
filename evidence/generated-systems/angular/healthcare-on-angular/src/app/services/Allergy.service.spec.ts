import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AllergyService } from './Allergy.service';

describe('AllergyService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AllergyService] });
	});

  it('should be created', () => {
    const service: AllergyService = TestBed.get(AllergyService);
    expect(service).toBeTruthy();
  });
});
