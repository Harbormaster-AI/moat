import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { System_Service } from './System_.service';

describe('System_Service', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [System_Service] });
	});

  it('should be created', () => {
    const service: System_Service = TestBed.get(System_Service);
    expect(service).toBeTruthy();
  });
});
