import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DataPipelineService } from './DataPipeline.service';

describe('DataPipelineService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DataPipelineService] });
	});

  it('should be created', () => {
    const service: DataPipelineService = TestBed.get(DataPipelineService);
    expect(service).toBeTruthy();
  });
});
