import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InventoryTransactionService } from './InventoryTransaction.service';

describe('InventoryTransactionService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InventoryTransactionService] });
	});

  it('should be created', () => {
    const service: InventoryTransactionService = TestBed.get(InventoryTransactionService);
    expect(service).toBeTruthy();
  });
});
