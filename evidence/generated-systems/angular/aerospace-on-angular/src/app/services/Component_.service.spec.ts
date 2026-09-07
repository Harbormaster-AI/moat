import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { Component_Service } from './Component_.service';

describe('Component_Service', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [Component_Service] });
	});

  it('should be created', () => {
    const service: Component_Service = TestBed.get(Component_Service);
    expect(service).toBeTruthy();
  });
});
