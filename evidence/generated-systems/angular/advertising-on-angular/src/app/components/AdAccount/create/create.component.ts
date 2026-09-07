import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AdAccountService } from '../../../services/AdAccount.service';
import { AdAccount } from '../../../models/AdAccount';
import { SubBaseComponent } from '../../AdAccount/sub.base.component';

@Component({
    selector: 'app-create-adAccount',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAdAccountComponent extends SubBaseComponent implements OnInit {

    title = 'Add AdAccount';

    adAccountForm: FormGroup;
    adAccount: AdAccount;

    constructor( http: HttpClient,
        private adAccountService: AdAccountService,
        private fb: FormBuilder,
        private router: Router
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

    
    addAdAccount(name, accountCode, defaultCurrency, defaultTimezone, Advertiser, Users, Campaigns, BillingProfile, Dsp, PerformanceMetrics): void {
        this.adAccountService
        .addAdAccount(name, accountCode, defaultCurrency, defaultTimezone, Advertiser, Users, Campaigns, BillingProfile, Dsp, PerformanceMetrics)
            .subscribe(() => {
                this.router.navigate(['/indexAdAccount']);
            });
    }

    ngOnInit(): void {
    }
}