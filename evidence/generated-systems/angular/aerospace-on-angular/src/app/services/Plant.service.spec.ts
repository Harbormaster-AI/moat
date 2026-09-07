import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PlantService } from './Plant.service';

describe('PlantService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PlantService] });
	});

  it('should be created', () => {
    const service: PlantService = TestBed.get(PlantService);
    expect(service).toBeTruthy();
  });
});
