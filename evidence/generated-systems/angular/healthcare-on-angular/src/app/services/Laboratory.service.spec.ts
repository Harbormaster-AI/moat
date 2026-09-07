import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { LaboratoryService } from './Laboratory.service';

describe('LaboratoryService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [LaboratoryService] });
	});

  it('should be created', () => {
    const service: LaboratoryService = TestBed.get(LaboratoryService);
    expect(service).toBeTruthy();
  });
});
