import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DataCategoryService } from './DataCategory.service';

describe('DataCategoryService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DataCategoryService] });
	});

  it('should be created', () => {
    const service: DataCategoryService = TestBed.get(DataCategoryService);
    expect(service).toBeTruthy();
  });
});
