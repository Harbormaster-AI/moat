import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { LabResultService } from './LabResult.service';

describe('LabResultService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [LabResultService] });
	});

  it('should be created', () => {
    const service: LabResultService = TestBed.get(LabResultService);
    expect(service).toBeTruthy();
  });
});
