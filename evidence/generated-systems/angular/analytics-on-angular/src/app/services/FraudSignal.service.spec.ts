import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { FraudSignalService } from './FraudSignal.service';

describe('FraudSignalService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [FraudSignalService] });
	});

  it('should be created', () => {
    const service: FraudSignalService = TestBed.get(FraudSignalService);
    expect(service).toBeTruthy();
  });
});
