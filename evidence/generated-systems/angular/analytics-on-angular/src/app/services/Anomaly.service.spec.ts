import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AnomalyService } from './Anomaly.service';

describe('AnomalyService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AnomalyService] });
	});

  it('should be created', () => {
    const service: AnomalyService = TestBed.get(AnomalyService);
    expect(service).toBeTruthy();
  });
});
