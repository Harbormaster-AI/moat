import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ComplianceAlertService } from './ComplianceAlert.service';

describe('ComplianceAlertService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ComplianceAlertService] });
	});

  it('should be created', () => {
    const service: ComplianceAlertService = TestBed.get(ComplianceAlertService);
    expect(service).toBeTruthy();
  });
});
