import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TargetingProfileService } from '../../../services/TargetingProfile.service';
import { TargetingProfile } from '../../../models/TargetingProfile';
import { SubBaseComponent } from '../../TargetingProfile/sub.base.component';

@Component({
    selector: 'app-create-targetingProfile',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTargetingProfileComponent extends SubBaseComponent implements OnInit {

    title = 'Add TargetingProfile';

    targetingProfileForm: FormGroup;
    targetingProfile: TargetingProfile;

    constructor( http: HttpClient,
        private targetingProfileService: TargetingProfileService,
        private fb: FormBuilder,
        private router: Router
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

    
    addTargetingProfile(name, AudienceSegments, GeoRegions, ContentCategories, BrandSafetyPolicy, DeviceCriteria): void {
        this.targetingProfileService
        .addTargetingProfile(name, AudienceSegments, GeoRegions, ContentCategories, BrandSafetyPolicy, DeviceCriteria)
            .subscribe(() => {
                this.router.navigate(['/indexTargetingProfile']);
            });
    }

    ngOnInit(): void {
    }
}