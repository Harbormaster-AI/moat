import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { GeoRegionService } from './GeoRegion.service';

describe('GeoRegionService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [GeoRegionService] });
	});

  it('should be created', () => {
    const service: GeoRegionService = TestBed.get(GeoRegionService);
    expect(service).toBeTruthy();
  });
});
