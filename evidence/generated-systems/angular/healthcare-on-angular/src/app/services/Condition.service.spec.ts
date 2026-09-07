import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ConditionService } from './Condition.service';

describe('ConditionService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ConditionService] });
	});

  it('should be created', () => {
    const service: ConditionService = TestBed.get(ConditionService);
    expect(service).toBeTruthy();
  });
});
