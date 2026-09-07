import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { SemanticModelService } from './SemanticModel.service';

describe('SemanticModelService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [SemanticModelService] });
	});

  it('should be created', () => {
    const service: SemanticModelService = TestBed.get(SemanticModelService);
    expect(service).toBeTruthy();
  });
});
