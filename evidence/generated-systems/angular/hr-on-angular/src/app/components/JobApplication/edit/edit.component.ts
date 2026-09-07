import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { JobApplicationService } from '../../../services/JobApplication.service';
import { SubBaseComponent } from '../../JobApplication/sub.base.component';


@Component({
    selector: 'app-edit-jobApplication',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditJobApplicationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit JobApplication';

    jobApplicationForm: FormGroup;
    jobApplication: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: JobApplicationService,
        private fb: FormBuilder
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

    
    updateJobApplication(applicationNumber, appliedDate, resumeUrl, Candidate, Requisition, Screenings, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateJobApplication(applicationNumber, appliedDate, resumeUrl, Candidate, Requisition, Screenings, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexJobApplication']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getJobApplication(params['id']).subscribe(res => {
                this.jobApplication = res;
            });
        });
    }
}