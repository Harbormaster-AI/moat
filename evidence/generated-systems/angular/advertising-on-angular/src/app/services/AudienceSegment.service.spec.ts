import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AudienceSegmentService } from './AudienceSegment.service';

describe('AudienceSegmentService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AudienceSegmentService] });
	});

  it('should be created', () => {
    const service: AudienceSegmentService = TestBed.get(AudienceSegmentService);
    expect(service).toBeTruthy();
  });
});
