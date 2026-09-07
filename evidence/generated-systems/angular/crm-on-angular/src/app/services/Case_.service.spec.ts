import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { Case_Service } from './Case_.service';

describe('Case_Service', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [Case_Service] });
	});

  it('should be created', () => {
    const service: Case_Service = TestBed.get(Case_Service);
    expect(service).toBeTruthy();
  });
});
