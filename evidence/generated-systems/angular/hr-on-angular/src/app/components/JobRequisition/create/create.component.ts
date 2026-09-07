import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { JobRequisitionService } from '../../../services/JobRequisition.service';
import { JobRequisition } from '../../../models/JobRequisition';
import { SubBaseComponent } from '../../JobRequisition/sub.base.component';

@Component({
    selector: 'app-create-jobRequisition',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateJobRequisitionComponent extends SubBaseComponent implements OnInit {

    title = 'Add JobRequisition';

    jobRequisitionForm: FormGroup;
    jobRequisition: JobRequisition;

    constructor( http: HttpClient,
        private jobRequisitionService: JobRequisitionService,
        private fb: FormBuilder,
        private router: Router
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

    
    addJobRequisition(requisitionNumber, title, openings, targetStartDate, Department, HiringManager, Recruiter, JobProfile, Candidates, Interviews, Offers, Status, Priority): void {
        this.jobRequisitionService
        .addJobRequisition(requisitionNumber, title, openings, targetStartDate, Department, HiringManager, Recruiter, JobProfile, Candidates, Interviews, Offers, Status, Priority)
            .subscribe(() => {
                this.router.navigate(['/indexJobRequisition']);
            });
    }

    ngOnInit(): void {
    }
}