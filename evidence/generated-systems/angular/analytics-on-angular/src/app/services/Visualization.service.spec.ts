import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { VisualizationService } from './Visualization.service';

describe('VisualizationService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [VisualizationService] });
	});

  it('should be created', () => {
    const service: VisualizationService = TestBed.get(VisualizationService);
    expect(service).toBeTruthy();
  });
});
