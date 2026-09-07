import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ProductionOrderService } from './ProductionOrder.service';

describe('ProductionOrderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ProductionOrderService] });
	});

  it('should be created', () => {
    const service: ProductionOrderService = TestBed.get(ProductionOrderService);
    expect(service).toBeTruthy();
  });
});
