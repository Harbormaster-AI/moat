import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TargetingProfileService } from '../../../services/TargetingProfile.service';
import { SubBaseComponent } from '../../TargetingProfile/sub.base.component';


@Component({
    selector: 'app-edit-targetingProfile',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTargetingProfileComponent extends SubBaseComponent implements OnInit {

    title = 'Edit TargetingProfile';

    targetingProfileForm: FormGroup;
    targetingProfile: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TargetingProfileService,
        private fb: FormBuilder
) {
        super(http);
        this.targetingProfileForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      AudienceSegments: ['', ],
      GeoRegions: ['', ],
      ContentCategories: ['', ],
      BrandSafetyPolicy: ['', ],
      DeviceCriteria: ['', ]
        });
    }

    
    updateTargetingProfile(name, AudienceSegments, GeoRegions, ContentCategories, BrandSafetyPolicy, DeviceCriteria): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTargetingProfile(name, AudienceSegments, GeoRegions, ContentCategories, BrandSafetyPolicy, DeviceCriteria, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTargetingProfile']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTargetingProfile(params['id']).subscribe(res => {
                this.targetingProfile = res;
            });
        });
    }
}