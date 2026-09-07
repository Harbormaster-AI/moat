import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CompensationPackageService } from './CompensationPackage.service';

describe('CompensationPackageService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CompensationPackageService] });
	});

  it('should be created', () => {
    const service: CompensationPackageService = TestBed.get(CompensationPackageService);
    expect(service).toBeTruthy();
  });
});
