import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { TrackingPixelService } from './TrackingPixel.service';

describe('TrackingPixelService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [TrackingPixelService] });
	});

  it('should be created', () => {
    const service: TrackingPixelService = TestBed.get(TrackingPixelService);
    expect(service).toBeTruthy();
  });
});
