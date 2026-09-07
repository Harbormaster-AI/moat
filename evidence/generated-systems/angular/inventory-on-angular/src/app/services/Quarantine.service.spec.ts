import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { QuarantineService } from './Quarantine.service';

describe('QuarantineService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [QuarantineService] });
	});

  it('should be created', () => {
    const service: QuarantineService = TestBed.get(QuarantineService);
    expect(service).toBeTruthy();
  });
});
