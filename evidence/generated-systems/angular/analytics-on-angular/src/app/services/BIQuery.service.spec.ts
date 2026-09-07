import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { BIQueryService } from './BIQuery.service';

describe('BIQueryService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [BIQueryService] });
	});

  it('should be created', () => {
    const service: BIQueryService = TestBed.get(BIQueryService);
    expect(service).toBeTruthy();
  });
});
