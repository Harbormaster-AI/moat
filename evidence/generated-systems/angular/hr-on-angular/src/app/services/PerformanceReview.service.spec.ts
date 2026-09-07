import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PerformanceReviewService } from './PerformanceReview.service';

describe('PerformanceReviewService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PerformanceReviewService] });
	});

  it('should be created', () => {
    const service: PerformanceReviewService = TestBed.get(PerformanceReviewService);
    expect(service).toBeTruthy();
  });
});
