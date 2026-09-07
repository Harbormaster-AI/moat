import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CertificationService } from './Certification.service';

describe('CertificationService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CertificationService] });
	});

  it('should be created', () => {
    const service: CertificationService = TestBed.get(CertificationService);
    expect(service).toBeTruthy();
  });
});
