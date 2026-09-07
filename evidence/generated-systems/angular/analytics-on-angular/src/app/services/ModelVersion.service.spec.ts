import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ModelVersionService } from './ModelVersion.service';

describe('ModelVersionService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ModelVersionService] });
	});

  it('should be created', () => {
    const service: ModelVersionService = TestBed.get(ModelVersionService);
    expect(service).toBeTruthy();
  });
});
