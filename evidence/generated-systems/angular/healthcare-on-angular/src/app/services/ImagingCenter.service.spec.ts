import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ImagingCenterService } from './ImagingCenter.service';

describe('ImagingCenterService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ImagingCenterService] });
	});

  it('should be created', () => {
    const service: ImagingCenterService = TestBed.get(ImagingCenterService);
    expect(service).toBeTruthy();
  });
});
