import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { OperatorService } from './Operator.service';

describe('OperatorService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [OperatorService] });
	});

  it('should be created', () => {
    const service: OperatorService = TestBed.get(OperatorService);
    expect(service).toBeTruthy();
  });
});
