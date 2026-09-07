import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DataProcessingActivityService } from './DataProcessingActivity.service';

describe('DataProcessingActivityService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DataProcessingActivityService] });
	});

  it('should be created', () => {
    const service: DataProcessingActivityService = TestBed.get(DataProcessingActivityService);
    expect(service).toBeTruthy();
  });
});
