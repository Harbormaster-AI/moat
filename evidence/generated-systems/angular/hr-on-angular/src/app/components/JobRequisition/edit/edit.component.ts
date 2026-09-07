import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { JobRequisitionService } from '../../../services/JobRequisition.service';
import { SubBaseComponent } from '../../JobRequisition/sub.base.component';


@Component({
    selector: 'app-edit-jobRequisition',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditJobRequisitionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit JobRequisition';

    jobRequisitionForm: FormGroup;
    jobRequisition: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: JobRequisitionService,
        private fb: FormBuilder
) {
        super(http);
        this.jobRequisitionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  requisitionNumber: ['', Validators.required],
      title: ['', Validators.required],
      openings: ['', Validators.required],
      targetStartDate: ['', Validators.required],
      Department: ['', ],
      HiringManager: ['', ],
      Recruiter: ['', ],
      JobProfile: ['', ],
      Candidates: ['', ],
      Interviews: ['', ],
      Offers: ['', ],
      Status: ['', ],
      Priority: ['', ]
        });
    }

    
    updateJobRequisition(requisitionNumber, title, openings, targetStartDate, Department, HiringManager, Recruiter, JobProfile, Candidates, Interviews, Offers, Status, Priority): void {
        this.route.params.subscribe((params) => {

                        this.service.updateJobRequisition(requisitionNumber, title, openings, targetStartDate, Department, HiringManager, Recruiter, JobProfile, Candidates, Interviews, Offers, Status, Priority, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexJobRequisition']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getJobRequisition(params['id']).subscribe(res => {
                this.jobRequisition = res;
            });
        });
    }
}