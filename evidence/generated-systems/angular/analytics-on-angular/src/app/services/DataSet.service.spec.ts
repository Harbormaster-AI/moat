import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DataSetService } from './DataSet.service';

describe('DataSetService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DataSetService] });
	});

  it('should be created', () => {
    const service: DataSetService = TestBed.get(DataSetService);
    expect(service).toBeTruthy();
  });
});
