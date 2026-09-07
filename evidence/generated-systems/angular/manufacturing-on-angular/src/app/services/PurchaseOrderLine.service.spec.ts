import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PurchaseOrderLineService } from './PurchaseOrderLine.service';

describe('PurchaseOrderLineService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PurchaseOrderLineService] });
	});

  it('should be created', () => {
    const service: PurchaseOrderLineService = TestBed.get(PurchaseOrderLineService);
    expect(service).toBeTruthy();
  });
});
