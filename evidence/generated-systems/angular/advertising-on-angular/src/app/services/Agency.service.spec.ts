import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AgencyService } from './Agency.service';

describe('AgencyService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AgencyService] });
	});

  it('should be created', () => {
    const service: AgencyService = TestBed.get(AgencyService);
    expect(service).toBeTruthy();
  });
});
