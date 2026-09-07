import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { TrainingCourseService } from './TrainingCourse.service';

describe('TrainingCourseService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [TrainingCourseService] });
	});

  it('should be created', () => {
    const service: TrainingCourseService = TestBed.get(TrainingCourseService);
    expect(service).toBeTruthy();
  });
});
