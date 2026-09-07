import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AdvertiserService } from '../../../services/Advertiser.service';
import { Advertiser } from '../../../models/Advertiser';
import { SubBaseComponent } from '../../Advertiser/sub.base.component';

@Component({
    selector: 'app-create-advertiser',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAdvertiserComponent extends SubBaseComponent implements OnInit {

    title = 'Add Advertiser';

    advertiserForm: FormGroup;
    advertiser: Advertiser;

    constructor( http: HttpClient,
        private advertiserService: AdvertiserService,
        private fb: FormBuilder,
        private router: Router
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

    
    addAdvertiser(name, legalName, industry, website, Agency, AdAccounts, BillingProfiles, Campaigns, TrackingPixels): void {
        this.advertiserService
        .addAdvertiser(name, legalName, industry, website, Agency, AdAccounts, BillingProfiles, Campaigns, TrackingPixels)
            .subscribe(() => {
                this.router.navigate(['/indexAdvertiser']);
            });
    }

    ngOnInit(): void {
    }
}