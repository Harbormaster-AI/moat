import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { EquityGrantService } from './EquityGrant.service';

describe('EquityGrantService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [EquityGrantService] });
	});

  it('should be created', () => {
    const service: EquityGrantService = TestBed.get(EquityGrantService);
    expect(service).toBeTruthy();
  });
});
