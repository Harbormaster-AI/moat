import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DataSourceService } from './DataSource.service';

describe('DataSourceService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DataSourceService] });
	});

  it('should be created', () => {
    const service: DataSourceService = TestBed.get(DataSourceService);
    expect(service).toBeTruthy();
  });
});
