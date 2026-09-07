import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { GoodsReceiptService } from './GoodsReceipt.service';

describe('GoodsReceiptService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [GoodsReceiptService] });
	});

  it('should be created', () => {
    const service: GoodsReceiptService = TestBed.get(GoodsReceiptService);
    expect(service).toBeTruthy();
  });
});
