import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { TargetingProfileService } from './TargetingProfile.service';

describe('TargetingProfileService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [TargetingProfileService] });
	});

  it('should be created', () => {
    const service: TargetingProfileService = TestBed.get(TargetingProfileService);
    expect(service).toBeTruthy();
  });
});
