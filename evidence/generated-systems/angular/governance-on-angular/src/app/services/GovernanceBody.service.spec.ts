import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { GovernanceBodyService } from './GovernanceBody.service';

describe('GovernanceBodyService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [GovernanceBodyService] });
	});

  it('should be created', () => {
    const service: GovernanceBodyService = TestBed.get(GovernanceBodyService);
    expect(service).toBeTruthy();
  });
});
