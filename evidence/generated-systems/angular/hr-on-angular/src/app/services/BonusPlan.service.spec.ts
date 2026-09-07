import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { BonusPlanService } from './BonusPlan.service';

describe('BonusPlanService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [BonusPlanService] });
	});

  it('should be created', () => {
    const service: BonusPlanService = TestBed.get(BonusPlanService);
    expect(service).toBeTruthy();
  });
});
