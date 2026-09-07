import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { NotebookService } from './Notebook.service';

describe('NotebookService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [NotebookService] });
	});

  it('should be created', () => {
    const service: NotebookService = TestBed.get(NotebookService);
    expect(service).toBeTruthy();
  });
});
