import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { TransferOrderService } from './TransferOrder.service';

describe('TransferOrderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [TransferOrderService] });
	});

  it('should be created', () => {
    const service: TransferOrderService = TestBed.get(TransferOrderService);
    expect(service).toBeTruthy();
  });
});
