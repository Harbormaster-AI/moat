import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { StockKeepingUnitService } from './StockKeepingUnit.service';

describe('StockKeepingUnitService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [StockKeepingUnitService] });
	});

  it('should be created', () => {
    const service: StockKeepingUnitService = TestBed.get(StockKeepingUnitService);
    expect(service).toBeTruthy();
  });
});
