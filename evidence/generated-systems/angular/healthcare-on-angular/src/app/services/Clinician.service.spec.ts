import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ClinicianService } from './Clinician.service';

describe('ClinicianService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ClinicianService] });
	});

  it('should be created', () => {
    const service: ClinicianService = TestBed.get(ClinicianService);
    expect(service).toBeTruthy();
  });
});
