import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { WorkAuthorizationService } from './WorkAuthorization.service';

describe('WorkAuthorizationService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [WorkAuthorizationService] });
	});

  it('should be created', () => {
    const service: WorkAuthorizationService = TestBed.get(WorkAuthorizationService);
    expect(service).toBeTruthy();
  });
});
