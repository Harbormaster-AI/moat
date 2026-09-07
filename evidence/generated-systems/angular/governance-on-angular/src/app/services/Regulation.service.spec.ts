import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { RegulationService } from './Regulation.service';

describe('RegulationService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [RegulationService] });
	});

  it('should be created', () => {
    const service: RegulationService = TestBed.get(RegulationService);
    expect(service).toBeTruthy();
  });
});
