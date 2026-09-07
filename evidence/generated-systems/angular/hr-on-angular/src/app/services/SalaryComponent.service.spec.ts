import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { SalaryComponentService } from './SalaryComponent.service';

describe('SalaryComponentService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [SalaryComponentService] });
	});

  it('should be created', () => {
    const service: SalaryComponentService = TestBed.get(SalaryComponentService);
    expect(service).toBeTruthy();
  });
});
