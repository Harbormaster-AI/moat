import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { FXDealService } from './FXDeal.service';

describe('FXDealService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [FXDealService] });
	});

  it('should be created', () => {
    const service: FXDealService = TestBed.get(FXDealService);
    expect(service).toBeTruthy();
  });
});
