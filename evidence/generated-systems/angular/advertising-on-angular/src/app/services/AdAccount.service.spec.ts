import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AdAccountService } from './AdAccount.service';

describe('AdAccountService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AdAccountService] });
	});

  it('should be created', () => {
    const service: AdAccountService = TestBed.get(AdAccountService);
    expect(service).toBeTruthy();
  });
});
