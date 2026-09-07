import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { HealthSystemService } from './HealthSystem.service';

describe('HealthSystemService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [HealthSystemService] });
	});

  it('should be created', () => {
    const service: HealthSystemService = TestBed.get(HealthSystemService);
    expect(service).toBeTruthy();
  });
});
