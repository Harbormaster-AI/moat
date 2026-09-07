import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { BrandSafetyPolicyService } from '../../../services/BrandSafetyPolicy.service';
import { SubBaseComponent } from '../../BrandSafetyPolicy/sub.base.component';


@Component({
    selector: 'app-edit-brandSafetyPolicy',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditBrandSafetyPolicyComponent extends SubBaseComponent implements OnInit {

    title = 'Edit BrandSafetyPolicy';

    brandSafetyPolicyForm: FormGroup;
    brandSafetyPolicy: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: BrandSafetyPolicyService,
        private fb: FormBuilder
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

    
    updateBrandSafetyPolicy(TargetingProfiles, Level, ContentRatingThreshold): void {
        this.route.params.subscribe((params) => {

                        this.service.updateBrandSafetyPolicy(TargetingProfiles, Level, ContentRatingThreshold, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexBrandSafetyPolicy']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getBrandSafetyPolicy(params['id']).subscribe(res => {
                this.brandSafetyPolicy = res;
            });
        });
    }
}