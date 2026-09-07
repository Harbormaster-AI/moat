import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DependentService } from './Dependent.service';

describe('DependentService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DependentService] });
	});

  it('should be created', () => {
    const service: DependentService = TestBed.get(DependentService);
    expect(service).toBeTruthy();
  });
});
