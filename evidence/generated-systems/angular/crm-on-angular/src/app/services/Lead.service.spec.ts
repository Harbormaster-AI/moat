import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { LeadService } from './Lead.service';

describe('LeadService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [LeadService] });
	});

  it('should be created', () => {
    const service: LeadService = TestBed.get(LeadService);
    expect(service).toBeTruthy();
  });
});
