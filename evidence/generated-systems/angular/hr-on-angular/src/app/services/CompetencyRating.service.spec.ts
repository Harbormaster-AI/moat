import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CompetencyRatingService } from './CompetencyRating.service';

describe('CompetencyRatingService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CompetencyRatingService] });
	});

  it('should be created', () => {
    const service: CompetencyRatingService = TestBed.get(CompetencyRatingService);
    expect(service).toBeTruthy();
  });
});
