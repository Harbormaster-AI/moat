import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { RateService } from './Rate.service';

describe('RateService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [RateService] });
	});

  it('should be created', () => {
    const service: RateService = TestBed.get(RateService);
    expect(service).toBeTruthy();
  });
});
