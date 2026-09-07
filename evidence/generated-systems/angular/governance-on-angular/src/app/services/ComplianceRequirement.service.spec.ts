import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ComplianceRequirementService } from './ComplianceRequirement.service';

describe('ComplianceRequirementService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ComplianceRequirementService] });
	});

  it('should be created', () => {
    const service: ComplianceRequirementService = TestBed.get(ComplianceRequirementService);
    expect(service).toBeTruthy();
  });
});
