import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { GoodsReceiptLineService } from './GoodsReceiptLine.service';

describe('GoodsReceiptLineService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [GoodsReceiptLineService] });
	});

  it('should be created', () => {
    const service: GoodsReceiptLineService = TestBed.get(GoodsReceiptLineService);
    expect(service).toBeTruthy();
  });
});
