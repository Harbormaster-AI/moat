import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { RateCardService } from './RateCard.service';

describe('RateCardService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [RateCardService] });
	});

  it('should be created', () => {
    const service: RateCardService = TestBed.get(RateCardService);
    expect(service).toBeTruthy();
  });
});
