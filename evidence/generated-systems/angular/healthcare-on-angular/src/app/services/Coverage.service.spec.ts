import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CoverageService } from './Coverage.service';

describe('CoverageService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CoverageService] });
	});

  it('should be created', () => {
    const service: CoverageService = TestBed.get(CoverageService);
    expect(service).toBeTruthy();
  });
});
