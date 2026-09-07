import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ThirdPartyService } from './ThirdParty.service';

describe('ThirdPartyService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ThirdPartyService] });
	});

  it('should be created', () => {
    const service: ThirdPartyService = TestBed.get(ThirdPartyService);
    expect(service).toBeTruthy();
  });
});
