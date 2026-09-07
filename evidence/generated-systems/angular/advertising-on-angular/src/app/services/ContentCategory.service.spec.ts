import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ContentCategoryService } from './ContentCategory.service';

describe('ContentCategoryService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ContentCategoryService] });
	});

  it('should be created', () => {
    const service: ContentCategoryService = TestBed.get(ContentCategoryService);
    expect(service).toBeTruthy();
  });
});
