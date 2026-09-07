import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AnalyticsWorkspaceService } from './AnalyticsWorkspace.service';

describe('AnalyticsWorkspaceService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AnalyticsWorkspaceService] });
	});

  it('should be created', () => {
    const service: AnalyticsWorkspaceService = TestBed.get(AnalyticsWorkspaceService);
    expect(service).toBeTruthy();
  });
});
