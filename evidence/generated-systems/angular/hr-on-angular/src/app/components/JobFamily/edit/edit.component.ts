import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { JobFamilyService } from '../../../services/JobFamily.service';
import { SubBaseComponent } from '../../JobFamily/sub.base.component';


@Component({
    selector: 'app-edit-jobFamily',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditJobFamilyComponent extends SubBaseComponent implements OnInit {

    title = 'Edit JobFamily';

    jobFamilyForm: FormGroup;
    jobFamily: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: JobFamilyService,
        private fb: FormBuilder
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

    
    updateJobFamily(name, description, Organization, JobProfiles): void {
        this.route.params.subscribe((params) => {

                        this.service.updateJobFamily(name, description, Organization, JobProfiles, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexJobFamily']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getJobFamily(params['id']).subscribe(res => {
                this.jobFamily = res;
            });
        });
    }
}