import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DeviceCriterionService } from './DeviceCriterion.service';

describe('DeviceCriterionService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DeviceCriterionService] });
	});

  it('should be created', () => {
    const service: DeviceCriterionService = TestBed.get(DeviceCriterionService);
    expect(service).toBeTruthy();
  });
});
