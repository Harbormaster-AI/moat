import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { EngineTypeService } from './EngineType.service';

describe('EngineTypeService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [EngineTypeService] });
	});

  it('should be created', () => {
    const service: EngineTypeService = TestBed.get(EngineTypeService);
    expect(service).toBeTruthy();
  });
});
