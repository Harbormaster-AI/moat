import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { FacilityService } from './Facility.service';

describe('FacilityService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [FacilityService] });
	});

  it('should be created', () => {
    const service: FacilityService = TestBed.get(FacilityService);
    expect(service).toBeTruthy();
  });
});
