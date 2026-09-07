import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AgentService } from './Agent.service';

describe('AgentService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AgentService] });
	});

  it('should be created', () => {
    const service: AgentService = TestBed.get(AgentService);
    expect(service).toBeTruthy();
  });
});
