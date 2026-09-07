import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { Model_Service } from './Model_.service';

describe('Model_Service', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [Model_Service] });
	});

  it('should be created', () => {
    const service: Model_Service = TestBed.get(Model_Service);
    expect(service).toBeTruthy();
  });
});
