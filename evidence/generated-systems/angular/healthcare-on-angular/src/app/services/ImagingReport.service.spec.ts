import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ImagingReportService } from './ImagingReport.service';

describe('ImagingReportService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ImagingReportService] });
	});

  it('should be created', () => {
    const service: ImagingReportService = TestBed.get(ImagingReportService);
    expect(service).toBeTruthy();
  });
});
