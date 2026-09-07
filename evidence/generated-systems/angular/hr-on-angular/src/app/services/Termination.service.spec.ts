import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { TerminationService } from './Termination.service';

describe('TerminationService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [TerminationService] });
	});

  it('should be created', () => {
    const service: TerminationService = TestBed.get(TerminationService);
    expect(service).toBeTruthy();
  });
});
