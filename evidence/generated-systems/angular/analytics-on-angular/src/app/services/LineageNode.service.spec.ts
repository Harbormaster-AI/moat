import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { LineageNodeService } from './LineageNode.service';

describe('LineageNodeService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [LineageNodeService] });
	});

  it('should be created', () => {
    const service: LineageNodeService = TestBed.get(LineageNodeService);
    expect(service).toBeTruthy();
  });
});
