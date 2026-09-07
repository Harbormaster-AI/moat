import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ExperimentVariantService } from './ExperimentVariant.service';

describe('ExperimentVariantService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ExperimentVariantService] });
	});

  it('should be created', () => {
    const service: ExperimentVariantService = TestBed.get(ExperimentVariantService);
    expect(service).toBeTruthy();
  });
});
