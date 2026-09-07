import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { StockAdjustmentLineService } from './StockAdjustmentLine.service';

describe('StockAdjustmentLineService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [StockAdjustmentLineService] });
	});

  it('should be created', () => {
    const service: StockAdjustmentLineService = TestBed.get(StockAdjustmentLineService);
    expect(service).toBeTruthy();
  });
});
