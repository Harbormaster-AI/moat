import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { QualitySpecificationService } from './QualitySpecification.service';

describe('QualitySpecificationService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [QualitySpecificationService] });
	});

  it('should be created', () => {
    const service: QualitySpecificationService = TestBed.get(QualitySpecificationService);
    expect(service).toBeTruthy();
  });
});
