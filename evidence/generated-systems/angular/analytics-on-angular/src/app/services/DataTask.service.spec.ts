import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DataTaskService } from './DataTask.service';

describe('DataTaskService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DataTaskService] });
	});

  it('should be created', () => {
    const service: DataTaskService = TestBed.get(DataTaskService);
    expect(service).toBeTruthy();
  });
});
