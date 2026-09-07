import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { VerifiedAddressService } from '../../../services/VerifiedAddress.service';
import { VerifiedAddress } from '../../../models/VerifiedAddress';
import { SubBaseComponent } from '../../VerifiedAddress/sub.base.component';

@Component({
    selector: 'app-create-verifiedAddress',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateVerifiedAddressComponent extends SubBaseComponent implements OnInit {

    title = 'Add VerifiedAddress';

    verifiedAddressForm: FormGroup;
    verifiedAddress: VerifiedAddress;

    constructor( http: HttpClient,
        private verifiedAddressService: VerifiedAddressService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.verifiedAddressForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  address: ['', Validators.required],
      verifiedAt: ['', Validators.required],
      KycProfile: ['', ],
      VerificationStatus: ['', ]
        });
    }

    
    addVerifiedAddress(address, verifiedAt, KycProfile, VerificationStatus): void {
        this.verifiedAddressService
        .addVerifiedAddress(address, verifiedAt, KycProfile, VerificationStatus)
            .subscribe(() => {
                this.router.navigate(['/indexVerifiedAddress']);
            });
    }

    ngOnInit(): void {
    }
}