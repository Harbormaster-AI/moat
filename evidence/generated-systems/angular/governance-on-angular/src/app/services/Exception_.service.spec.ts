import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { Exception_Service } from './Exception_.service';

describe('Exception_Service', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [Exception_Service] });
	});

  it('should be created', () => {
    const service: Exception_Service = TestBed.get(Exception_Service);
    expect(service).toBeTruthy();
  });
});
