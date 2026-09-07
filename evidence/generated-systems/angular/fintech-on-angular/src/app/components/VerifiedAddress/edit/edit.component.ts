import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { VerifiedAddressService } from '../../../services/VerifiedAddress.service';
import { SubBaseComponent } from '../../VerifiedAddress/sub.base.component';


@Component({
    selector: 'app-edit-verifiedAddress',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditVerifiedAddressComponent extends SubBaseComponent implements OnInit {

    title = 'Edit VerifiedAddress';

    verifiedAddressForm: FormGroup;
    verifiedAddress: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: VerifiedAddressService,
        private fb: FormBuilder
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

    
    updateVerifiedAddress(address, verifiedAt, KycProfile, VerificationStatus): void {
        this.route.params.subscribe((params) => {

                        this.service.updateVerifiedAddress(address, verifiedAt, KycProfile, VerificationStatus, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexVerifiedAddress']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getVerifiedAddress(params['id']).subscribe(res => {
                this.verifiedAddress = res;
            });
        });
    }
}