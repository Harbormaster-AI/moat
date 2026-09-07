import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { TrainingRunService } from './TrainingRun.service';

describe('TrainingRunService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [TrainingRunService] });
	});

  it('should be created', () => {
    const service: TrainingRunService = TestBed.get(TrainingRunService);
    expect(service).toBeTruthy();
  });
});
