import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { JobProfileService } from '../../../services/JobProfile.service';
import { JobProfile } from '../../../models/JobProfile';
import { SubBaseComponent } from '../../JobProfile/sub.base.component';

@Component({
    selector: 'app-create-jobProfile',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateJobProfileComponent extends SubBaseComponent implements OnInit {

    title = 'Add JobProfile';

    jobProfileForm: FormGroup;
    jobProfile: JobProfile;

    constructor( http: HttpClient,
        private jobProfileService: JobProfileService,
        private fb: FormBuilder,
        private router: Router
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

    
    addJobProfile(title, jobCode, JobFamily, Competencies, TrainingRecommendations, Positions, JobLevel, ExemptStatus): void {
        this.jobProfileService
        .addJobProfile(title, jobCode, JobFamily, Competencies, TrainingRecommendations, Positions, JobLevel, ExemptStatus)
            .subscribe(() => {
                this.router.navigate(['/indexJobProfile']);
            });
    }

    ngOnInit(): void {
    }
}