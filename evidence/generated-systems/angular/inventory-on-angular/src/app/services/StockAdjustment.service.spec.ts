import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { StockAdjustmentService } from './StockAdjustment.service';

describe('StockAdjustmentService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [StockAdjustmentService] });
	});

  it('should be created', () => {
    const service: StockAdjustmentService = TestBed.get(StockAdjustmentService);
    expect(service).toBeTruthy();
  });
});
