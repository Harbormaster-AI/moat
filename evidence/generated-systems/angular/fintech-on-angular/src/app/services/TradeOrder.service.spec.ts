import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { TradeOrderService } from './TradeOrder.service';

describe('TradeOrderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [TradeOrderService] });
	});

  it('should be created', () => {
    const service: TradeOrderService = TestBed.get(TradeOrderService);
    expect(service).toBeTruthy();
  });
});
