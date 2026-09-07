import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { JobFamilyService } from '../../../services/JobFamily.service';
import { JobFamily } from '../../../models/JobFamily';
import { SubBaseComponent } from '../../JobFamily/sub.base.component';

@Component({
    selector: 'app-create-jobFamily',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateJobFamilyComponent extends SubBaseComponent implements OnInit {

    title = 'Add JobFamily';

    jobFamilyForm: FormGroup;
    jobFamily: JobFamily;

    constructor( http: HttpClient,
        private jobFamilyService: JobFamilyService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.jobFamilyForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      description: ['', Validators.required],
      Organization: ['', ],
      JobProfiles: ['', ]
        });
    }

    
    addJobFamily(name, description, Organization, JobProfiles): void {
        this.jobFamilyService
        .addJobFamily(name, description, Organization, JobProfiles)
            .subscribe(() => {
                this.router.navigate(['/indexJobFamily']);
            });
    }

    ngOnInit(): void {
    }
}