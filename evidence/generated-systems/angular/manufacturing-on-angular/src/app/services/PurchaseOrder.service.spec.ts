import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PurchaseOrderService } from './PurchaseOrder.service';

describe('PurchaseOrderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PurchaseOrderService] });
	});

  it('should be created', () => {
    const service: PurchaseOrderService = TestBed.get(PurchaseOrderService);
    expect(service).toBeTruthy();
  });
});
