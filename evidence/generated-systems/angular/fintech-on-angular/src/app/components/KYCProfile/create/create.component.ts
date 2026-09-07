import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { KYCProfileService } from '../../../services/KYCProfile.service';
import { KYCProfile } from '../../../models/KYCProfile';
import { SubBaseComponent } from '../../KYCProfile/sub.base.component';

@Component({
    selector: 'app-create-kYCProfile',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateKYCProfileComponent extends SubBaseComponent implements OnInit {

    title = 'Add KYCProfile';

    kYCProfileForm: FormGroup;
    kYCProfile: KYCProfile;

    constructor( http: HttpClient,
        private kYCProfileService: KYCProfileService,
        private fb: FormBuilder,
        private router: Router
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

    
    addKYCProfile(profileId, createdAt, Customer, Documents, Screenings, Addresses, Status, VerificationLevel): void {
        this.kYCProfileService
        .addKYCProfile(profileId, createdAt, Customer, Documents, Screenings, Addresses, Status, VerificationLevel)
            .subscribe(() => {
                this.router.navigate(['/indexKYCProfile']);
            });
    }

    ngOnInit(): void {
    }
}