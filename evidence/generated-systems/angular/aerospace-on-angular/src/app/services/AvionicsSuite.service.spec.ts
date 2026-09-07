import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AvionicsSuiteService } from './AvionicsSuite.service';

describe('AvionicsSuiteService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AvionicsSuiteService] });
	});

  it('should be created', () => {
    const service: AvionicsSuiteService = TestBed.get(AvionicsSuiteService);
    expect(service).toBeTruthy();
  });
});
