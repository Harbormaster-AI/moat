import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DataProviderService } from './DataProvider.service';

describe('DataProviderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DataProviderService] });
	});

  it('should be created', () => {
    const service: DataProviderService = TestBed.get(DataProviderService);
    expect(service).toBeTruthy();
  });
});
