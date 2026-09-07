import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CreativeFileService } from './CreativeFile.service';

describe('CreativeFileService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CreativeFileService] });
	});

  it('should be created', () => {
    const service: CreativeFileService = TestBed.get(CreativeFileService);
    expect(service).toBeTruthy();
  });
});
