import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { MROFacilityService } from './MROFacility.service';

describe('MROFacilityService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [MROFacilityService] });
	});

  it('should be created', () => {
    const service: MROFacilityService = TestBed.get(MROFacilityService);
    expect(service).toBeTruthy();
  });
});
