import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { TransferOrderLineService } from './TransferOrderLine.service';

describe('TransferOrderLineService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [TransferOrderLineService] });
	});

  it('should be created', () => {
    const service: TransferOrderLineService = TestBed.get(TransferOrderLineService);
    expect(service).toBeTruthy();
  });
});
