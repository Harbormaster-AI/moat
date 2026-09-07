import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InspectionResultService } from './InspectionResult.service';

describe('InspectionResultService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InspectionResultService] });
	});

  it('should be created', () => {
    const service: InspectionResultService = TestBed.get(InspectionResultService);
    expect(service).toBeTruthy();
  });
});
