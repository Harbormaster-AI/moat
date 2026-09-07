import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { OnboardingTaskService } from './OnboardingTask.service';

describe('OnboardingTaskService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [OnboardingTaskService] });
	});

  it('should be created', () => {
    const service: OnboardingTaskService = TestBed.get(OnboardingTaskService);
    expect(service).toBeTruthy();
  });
});
