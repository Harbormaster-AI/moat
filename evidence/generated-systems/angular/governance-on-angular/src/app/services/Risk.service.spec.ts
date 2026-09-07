import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { RiskService } from './Risk.service';

describe('RiskService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [RiskService] });
	});

  it('should be created', () => {
    const service: RiskService = TestBed.get(RiskService);
    expect(service).toBeTruthy();
  });
});
