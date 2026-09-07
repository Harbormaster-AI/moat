import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ExperimentService } from './Experiment.service';

describe('ExperimentService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ExperimentService] });
	});

  it('should be created', () => {
    const service: ExperimentService = TestBed.get(ExperimentService);
    expect(service).toBeTruthy();
  });
});
