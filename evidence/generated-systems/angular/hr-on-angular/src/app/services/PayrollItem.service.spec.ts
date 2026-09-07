import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PayrollItemService } from './PayrollItem.service';

describe('PayrollItemService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PayrollItemService] });
	});

  it('should be created', () => {
    const service: PayrollItemService = TestBed.get(PayrollItemService);
    expect(service).toBeTruthy();
  });
});
