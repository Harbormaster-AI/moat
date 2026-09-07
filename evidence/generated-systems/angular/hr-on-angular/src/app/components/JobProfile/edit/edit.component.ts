import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { JobProfileService } from '../../../services/JobProfile.service';
import { SubBaseComponent } from '../../JobProfile/sub.base.component';


@Component({
    selector: 'app-edit-jobProfile',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditJobProfileComponent extends SubBaseComponent implements OnInit {

    title = 'Edit JobProfile';

    jobProfileForm: FormGroup;
    jobProfile: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: JobProfileService,
        private fb: FormBuilder
) {
        super(http);
        this.jobProfileForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  title: ['', Validators.required],
      jobCode: ['', Validators.required],
      JobFamily: ['', ],
      Competencies: ['', ],
      TrainingRecommendations: ['', ],
      Positions: ['', ],
      JobLevel: ['', ],
      ExemptStatus: ['', ]
        });
    }

    
    updateJobProfile(title, jobCode, JobFamily, Competencies, TrainingRecommendations, Positions, JobLevel, ExemptStatus): void {
        this.route.params.subscribe((params) => {

                        this.service.updateJobProfile(title, jobCode, JobFamily, Competencies, TrainingRecommendations, Positions, JobLevel, ExemptStatus, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexJobProfile']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getJobProfile(params['id']).subscribe(res => {
                this.jobProfile = res;
            });
        });
    }
}