import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { OpportunityService } from './Opportunity.service';

describe('OpportunityService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [OpportunityService] });
	});

  it('should be created', () => {
    const service: OpportunityService = TestBed.get(OpportunityService);
    expect(service).toBeTruthy();
  });
});
