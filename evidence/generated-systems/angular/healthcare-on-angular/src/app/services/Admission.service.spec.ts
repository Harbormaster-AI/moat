import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AdmissionService } from './Admission.service';

describe('AdmissionService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AdmissionService] });
	});

  it('should be created', () => {
    const service: AdmissionService = TestBed.get(AdmissionService);
    expect(service).toBeTruthy();
  });
});
