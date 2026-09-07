import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AuditProgramService } from './AuditProgram.service';

describe('AuditProgramService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AuditProgramService] });
	});

  it('should be created', () => {
    const service: AuditProgramService = TestBed.get(AuditProgramService);
    expect(service).toBeTruthy();
  });
});
