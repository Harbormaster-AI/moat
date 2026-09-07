import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { WorkCenterService } from './WorkCenter.service';

describe('WorkCenterService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [WorkCenterService] });
	});

  it('should be created', () => {
    const service: WorkCenterService = TestBed.get(WorkCenterService);
    expect(service).toBeTruthy();
  });
});
