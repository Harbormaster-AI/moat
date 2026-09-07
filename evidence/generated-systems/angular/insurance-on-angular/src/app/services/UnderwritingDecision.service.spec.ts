import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { UnderwritingDecisionService } from './UnderwritingDecision.service';

describe('UnderwritingDecisionService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [UnderwritingDecisionService] });
	});

  it('should be created', () => {
    const service: UnderwritingDecisionService = TestBed.get(UnderwritingDecisionService);
    expect(service).toBeTruthy();
  });
});
