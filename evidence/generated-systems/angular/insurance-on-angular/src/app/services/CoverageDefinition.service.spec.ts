import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CoverageDefinitionService } from './CoverageDefinition.service';

describe('CoverageDefinitionService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CoverageDefinitionService] });
	});

  it('should be created', () => {
    const service: CoverageDefinitionService = TestBed.get(CoverageDefinitionService);
    expect(service).toBeTruthy();
  });
});
