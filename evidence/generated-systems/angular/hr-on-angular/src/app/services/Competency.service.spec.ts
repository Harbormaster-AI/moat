import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CompetencyService } from './Competency.service';

describe('CompetencyService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CompetencyService] });
	});

  it('should be created', () => {
    const service: CompetencyService = TestBed.get(CompetencyService);
    expect(service).toBeTruthy();
  });
});
