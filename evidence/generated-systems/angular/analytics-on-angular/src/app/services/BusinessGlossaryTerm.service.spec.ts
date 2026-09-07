import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { BusinessGlossaryTermService } from './BusinessGlossaryTerm.service';

describe('BusinessGlossaryTermService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [BusinessGlossaryTermService] });
	});

  it('should be created', () => {
    const service: BusinessGlossaryTermService = TestBed.get(BusinessGlossaryTermService);
    expect(service).toBeTruthy();
  });
});
