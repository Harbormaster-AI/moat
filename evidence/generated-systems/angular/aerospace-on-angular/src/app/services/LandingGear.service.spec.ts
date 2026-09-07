import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { LandingGearService } from './LandingGear.service';

describe('LandingGearService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [LandingGearService] });
	});

  it('should be created', () => {
    const service: LandingGearService = TestBed.get(LandingGearService);
    expect(service).toBeTruthy();
  });
});
