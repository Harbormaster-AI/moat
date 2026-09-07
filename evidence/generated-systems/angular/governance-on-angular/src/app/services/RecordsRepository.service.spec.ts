import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { RecordsRepositoryService } from './RecordsRepository.service';

describe('RecordsRepositoryService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [RecordsRepositoryService] });
	});

  it('should be created', () => {
    const service: RecordsRepositoryService = TestBed.get(RecordsRepositoryService);
    expect(service).toBeTruthy();
  });
});
