import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PlacementService } from './Placement.service';

describe('PlacementService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PlacementService] });
	});

  it('should be created', () => {
    const service: PlacementService = TestBed.get(PlacementService);
    expect(service).toBeTruthy();
  });
});
