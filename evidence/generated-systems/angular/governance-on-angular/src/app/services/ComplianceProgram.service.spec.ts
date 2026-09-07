import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ComplianceProgramService } from './ComplianceProgram.service';

describe('ComplianceProgramService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ComplianceProgramService] });
	});

  it('should be created', () => {
    const service: ComplianceProgramService = TestBed.get(ComplianceProgramService);
    expect(service).toBeTruthy();
  });
});
