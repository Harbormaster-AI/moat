import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { JobApplicationService } from '../../../services/JobApplication.service';
import { JobApplication } from '../../../models/JobApplication';
import { SubBaseComponent } from '../../JobApplication/sub.base.component';

@Component({
    selector: 'app-create-jobApplication',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateJobApplicationComponent extends SubBaseComponent implements OnInit {

    title = 'Add JobApplication';

    jobApplicationForm: FormGroup;
    jobApplication: JobApplication;

    constructor( http: HttpClient,
        private jobApplicationService: JobApplicationService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.jobApplicationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  applicationNumber: ['', Validators.required],
      appliedDate: ['', Validators.required],
      resumeUrl: ['', Validators.required],
      Candidate: ['', ],
      Requisition: ['', ],
      Screenings: ['', ],
      Status: ['', ]
        });
    }

    
    addJobApplication(applicationNumber, appliedDate, resumeUrl, Candidate, Requisition, Screenings, Status): void {
        this.jobApplicationService
        .addJobApplication(applicationNumber, appliedDate, resumeUrl, Candidate, Requisition, Screenings, Status)
            .subscribe(() => {
                this.router.navigate(['/indexJobApplication']);
            });
    }

    ngOnInit(): void {
    }
}