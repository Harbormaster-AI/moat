import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { UsageLimitService } from './UsageLimit.service';

describe('UsageLimitService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [UsageLimitService] });
	});

  it('should be created', () => {
    const service: UsageLimitService = TestBed.get(UsageLimitService);
    expect(service).toBeTruthy();
  });
});
