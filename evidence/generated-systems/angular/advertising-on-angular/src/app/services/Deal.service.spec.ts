import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DealService } from './Deal.service';

describe('DealService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DealService] });
	});

  it('should be created', () => {
    const service: DealService = TestBed.get(DealService);
    expect(service).toBeTruthy();
  });
});
