import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AdAccountService } from '../../../services/AdAccount.service';
import { SubBaseComponent } from '../../AdAccount/sub.base.component';


@Component({
    selector: 'app-edit-adAccount',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAdAccountComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AdAccount';

    adAccountForm: FormGroup;
    adAccount: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AdAccountService,
        private fb: FormBuilder
) {
        super(http);
        this.adAccountForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      accountCode: ['', Validators.required],
      defaultCurrency: ['', Validators.required],
      defaultTimezone: ['', Validators.required],
      Advertiser: ['', ],
      Users: ['', ],
      Campaigns: ['', ],
      BillingProfile: ['', ],
      Dsp: ['', ],
      PerformanceMetrics: ['', ]
        });
    }

    
    updateAdAccount(name, accountCode, defaultCurrency, defaultTimezone, Advertiser, Users, Campaigns, BillingProfile, Dsp, PerformanceMetrics): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAdAccount(name, accountCode, defaultCurrency, defaultTimezone, Advertiser, Users, Campaigns, BillingProfile, Dsp, PerformanceMetrics, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAdAccount']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAdAccount(params['id']).subscribe(res => {
                this.adAccount = res;
            });
        });
    }
}