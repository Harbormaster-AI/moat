import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { OperationService } from './Operation.service';

describe('OperationService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [OperationService] });
	});

  it('should be created', () => {
    const service: OperationService = TestBed.get(OperationService);
    expect(service).toBeTruthy();
  });
});
