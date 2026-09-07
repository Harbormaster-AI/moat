import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { EnterpriseService } from './Enterprise.service';

describe('EnterpriseService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [EnterpriseService] });
	});

  it('should be created', () => {
    const service: EnterpriseService = TestBed.get(EnterpriseService);
    expect(service).toBeTruthy();
  });
});
