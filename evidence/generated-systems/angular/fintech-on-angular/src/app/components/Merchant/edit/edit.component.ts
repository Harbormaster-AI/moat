import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { MerchantService } from '../../../services/Merchant.service';
import { SubBaseComponent } from '../../Merchant/sub.base.component';


@Component({
    selector: 'app-edit-merchant',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditMerchantComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Merchant';

    merchantForm: FormGroup;
    merchant: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: MerchantService,
        private fb: FormBuilder
) {
        super(http);
        this.merchantForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      mcc: ['', Validators.required],
      url: ['', Validators.required],
      country: ['', Validators.required],
      settlementCurrency: ['', Validators.required],
      Terminals: ['', ],
      PaymentContracts: ['', ],
      Payouts: ['', ],
      Settlements: ['', ],
      Disputes: ['', ],
      Invoices: ['', ]
        });
    }

    
    updateMerchant(name, mcc, url, country, settlementCurrency, Terminals, PaymentContracts, Payouts, Settlements, Disputes, Invoices): void {
        this.route.params.subscribe((params) => {

                        this.service.updateMerchant(name, mcc, url, country, settlementCurrency, Terminals, PaymentContracts, Payouts, Settlements, Disputes, Invoices, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexMerchant']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getMerchant(params['id']).subscribe(res => {
                this.merchant = res;
            });
        });
    }
}