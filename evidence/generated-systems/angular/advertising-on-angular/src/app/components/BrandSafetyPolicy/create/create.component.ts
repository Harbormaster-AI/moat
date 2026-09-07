import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BrandSafetyPolicyService } from '../../../services/BrandSafetyPolicy.service';
import { BrandSafetyPolicy } from '../../../models/BrandSafetyPolicy';
import { SubBaseComponent } from '../../BrandSafetyPolicy/sub.base.component';

@Component({
    selector: 'app-create-brandSafetyPolicy',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateBrandSafetyPolicyComponent extends SubBaseComponent implements OnInit {

    title = 'Add BrandSafetyPolicy';

    brandSafetyPolicyForm: FormGroup;
    brandSafetyPolicy: BrandSafetyPolicy;

    constructor( http: HttpClient,
        private brandSafetyPolicyService: BrandSafetyPolicyService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.brandSafetyPolicyForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  TargetingProfiles: ['', ],
      Level: ['', ],
      ContentRatingThreshold: ['', ]
        });
    }

    
    addBrandSafetyPolicy(TargetingProfiles, Level, ContentRatingThreshold): void {
        this.brandSafetyPolicyService
        .addBrandSafetyPolicy(TargetingProfiles, Level, ContentRatingThreshold)
            .subscribe(() => {
                this.router.navigate(['/indexBrandSafetyPolicy']);
            });
    }

    ngOnInit(): void {
    }
}