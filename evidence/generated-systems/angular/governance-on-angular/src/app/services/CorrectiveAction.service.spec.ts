import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CorrectiveActionService } from './CorrectiveAction.service';

describe('CorrectiveActionService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CorrectiveActionService] });
	});

  it('should be created', () => {
    const service: CorrectiveActionService = TestBed.get(CorrectiveActionService);
    expect(service).toBeTruthy();
  });
});
