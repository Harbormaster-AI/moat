import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { QualityCheckService } from './QualityCheck.service';

describe('QualityCheckService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [QualityCheckService] });
	});

  it('should be created', () => {
    const service: QualityCheckService = TestBed.get(QualityCheckService);
    expect(service).toBeTruthy();
  });
});
