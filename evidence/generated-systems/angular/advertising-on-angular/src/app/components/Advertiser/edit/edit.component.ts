import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AdvertiserService } from '../../../services/Advertiser.service';
import { SubBaseComponent } from '../../Advertiser/sub.base.component';


@Component({
    selector: 'app-edit-advertiser',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAdvertiserComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Advertiser';

    advertiserForm: FormGroup;
    advertiser: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AdvertiserService,
        private fb: FormBuilder
) {
        super(http);
        this.advertiserForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      legalName: ['', Validators.required],
      industry: ['', Validators.required],
      website: ['', Validators.required],
      Agency: ['', ],
      AdAccounts: ['', ],
      BillingProfiles: ['', ],
      Campaigns: ['', ],
      TrackingPixels: ['', ]
        });
    }

    
    updateAdvertiser(name, legalName, industry, website, Agency, AdAccounts, BillingProfiles, Campaigns, TrackingPixels): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAdvertiser(name, legalName, industry, website, Agency, AdAccounts, BillingProfiles, Campaigns, TrackingPixels, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAdvertiser']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAdvertiser(params['id']).subscribe(res => {
                this.advertiser = res;
            });
        });
    }
}