import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { TypeCertificateService } from './TypeCertificate.service';

describe('TypeCertificateService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [TypeCertificateService] });
	});

  it('should be created', () => {
    const service: TypeCertificateService = TestBed.get(TypeCertificateService);
    expect(service).toBeTruthy();
  });
});
