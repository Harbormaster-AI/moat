import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DiagnosisService } from './Diagnosis.service';

describe('DiagnosisService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DiagnosisService] });
	});

  it('should be created', () => {
    const service: DiagnosisService = TestBed.get(DiagnosisService);
    expect(service).toBeTruthy();
  });
});
