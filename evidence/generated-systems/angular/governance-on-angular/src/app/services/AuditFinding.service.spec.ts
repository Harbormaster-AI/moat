import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AuditFindingService } from './AuditFinding.service';

describe('AuditFindingService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AuditFindingService] });
	});

  it('should be created', () => {
    const service: AuditFindingService = TestBed.get(AuditFindingService);
    expect(service).toBeTruthy();
  });
});
