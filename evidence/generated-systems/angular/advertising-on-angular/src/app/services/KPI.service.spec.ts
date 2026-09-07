import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { KPIService } from './KPI.service';

describe('KPIService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [KPIService] });
	});

  it('should be created', () => {
    const service: KPIService = TestBed.get(KPIService);
    expect(service).toBeTruthy();
  });
});
