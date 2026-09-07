import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CareTaskService } from './CareTask.service';

describe('CareTaskService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CareTaskService] });
	});

  it('should be created', () => {
    const service: CareTaskService = TestBed.get(CareTaskService);
    expect(service).toBeTruthy();
  });
});
