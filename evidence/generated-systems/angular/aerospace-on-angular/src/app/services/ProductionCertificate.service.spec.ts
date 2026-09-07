import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ProductionCertificateService } from './ProductionCertificate.service';

describe('ProductionCertificateService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ProductionCertificateService] });
	});

  it('should be created', () => {
    const service: ProductionCertificateService = TestBed.get(ProductionCertificateService);
    expect(service).toBeTruthy();
  });
});
