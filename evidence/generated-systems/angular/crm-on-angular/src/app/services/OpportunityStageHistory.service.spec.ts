import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { OpportunityStageHistoryService } from './OpportunityStageHistory.service';

describe('OpportunityStageHistoryService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [OpportunityStageHistoryService] });
	});

  it('should be created', () => {
    const service: OpportunityStageHistoryService = TestBed.get(OpportunityStageHistoryService);
    expect(service).toBeTruthy();
  });
});
