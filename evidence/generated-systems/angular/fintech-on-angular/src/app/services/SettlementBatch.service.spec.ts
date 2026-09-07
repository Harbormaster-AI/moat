import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { SettlementBatchService } from './SettlementBatch.service';

describe('SettlementBatchService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [SettlementBatchService] });
	});

  it('should be created', () => {
    const service: SettlementBatchService = TestBed.get(SettlementBatchService);
    expect(service).toBeTruthy();
  });
});
