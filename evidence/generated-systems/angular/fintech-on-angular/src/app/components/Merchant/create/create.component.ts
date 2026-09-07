import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MerchantService } from '../../../services/Merchant.service';
import { Merchant } from '../../../models/Merchant';
import { SubBaseComponent } from '../../Merchant/sub.base.component';

@Component({
    selector: 'app-create-merchant',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateMerchantComponent extends SubBaseComponent implements OnInit {

    title = 'Add Merchant';

    merchantForm: FormGroup;
    merchant: Merchant;

    constructor( http: HttpClient,
        private merchantService: MerchantService,
        private fb: FormBuilder,
        private router: Router
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

    
    addMerchant(name, mcc, url, country, settlementCurrency, Terminals, PaymentContracts, Payouts, Settlements, Disputes, Invoices): void {
        this.merchantService
        .addMerchant(name, mcc, url, country, settlementCurrency, Terminals, PaymentContracts, Payouts, Settlements, Disputes, Invoices)
            .subscribe(() => {
                this.router.navigate(['/indexMerchant']);
            });
    }

    ngOnInit(): void {
    }
}