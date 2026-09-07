import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { TradeService } from './Trade.service';

describe('TradeService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [TradeService] });
	});

  it('should be created', () => {
    const service: TradeService = TestBed.get(TradeService);
    expect(service).toBeTruthy();
  });
});
