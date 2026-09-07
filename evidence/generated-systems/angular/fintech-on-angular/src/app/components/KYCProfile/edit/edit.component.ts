import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { KYCProfileService } from '../../../services/KYCProfile.service';
import { SubBaseComponent } from '../../KYCProfile/sub.base.component';


@Component({
    selector: 'app-edit-kYCProfile',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditKYCProfileComponent extends SubBaseComponent implements OnInit {

    title = 'Edit KYCProfile';

    kYCProfileForm: FormGroup;
    kYCProfile: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: KYCProfileService,
        private fb: FormBuilder
) {
        super(http);
        this.kYCProfileForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  profileId: ['', Validators.required],
      createdAt: ['', Validators.required],
      Customer: ['', ],
      Documents: ['', ],
      Screenings: ['', ],
      Addresses: ['', ],
      Status: ['', ],
      VerificationLevel: ['', ]
        });
    }

    
    updateKYCProfile(profileId, createdAt, Customer, Documents, Screenings, Addresses, Status, VerificationLevel): void {
        this.route.params.subscribe((params) => {

                        this.service.updateKYCProfile(profileId, createdAt, Customer, Documents, Screenings, Addresses, Status, VerificationLevel, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexKYCProfile']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getKYCProfile(params['id']).subscribe(res => {
                this.kYCProfile = res;
            });
        });
    }
}