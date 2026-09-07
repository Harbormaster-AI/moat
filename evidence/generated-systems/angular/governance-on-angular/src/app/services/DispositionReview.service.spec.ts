import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DispositionReviewService } from './DispositionReview.service';

describe('DispositionReviewService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DispositionReviewService] });
	});

  it('should be created', () => {
    const service: DispositionReviewService = TestBed.get(DispositionReviewService);
    expect(service).toBeTruthy();
  });
});
