import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { StorageLocationService } from './StorageLocation.service';

describe('StorageLocationService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [StorageLocationService] });
	});

  it('should be created', () => {
    const service: StorageLocationService = TestBed.get(StorageLocationService);
    expect(service).toBeTruthy();
  });
});
