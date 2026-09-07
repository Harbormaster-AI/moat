import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { WalletService } from '../../../services/Wallet.service';
import { SubBaseComponent } from '../../Wallet/sub.base.component';


@Component({
    selector: 'app-edit-wallet',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditWalletComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Wallet';

    walletForm: FormGroup;
    wallet: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: WalletService,
        private fb: FormBuilder
) {
        super(http);
        this.walletForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  currency: ['', Validators.required],
      balance: ['', Validators.required],
      Customer: ['', ],
      Transactions: ['', ],
      Status: ['', ]
        });
    }

    
    updateWallet(currency, balance, Customer, Transactions, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateWallet(currency, balance, Customer, Transactions, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexWallet']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getWallet(params['id']).subscribe(res => {
                this.wallet = res;
            });
        });
    }
}