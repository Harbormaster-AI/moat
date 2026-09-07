import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { KYCDocumentService } from './KYCDocument.service';

describe('KYCDocumentService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [KYCDocumentService] });
	});

  it('should be created', () => {
    const service: KYCDocumentService = TestBed.get(KYCDocumentService);
    expect(service).toBeTruthy();
  });
});
