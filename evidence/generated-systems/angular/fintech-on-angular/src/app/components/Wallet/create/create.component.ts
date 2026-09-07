import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { WalletService } from '../../../services/Wallet.service';
import { Wallet } from '../../../models/Wallet';
import { SubBaseComponent } from '../../Wallet/sub.base.component';

@Component({
    selector: 'app-create-wallet',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateWalletComponent extends SubBaseComponent implements OnInit {

    title = 'Add Wallet';

    walletForm: FormGroup;
    wallet: Wallet;

    constructor( http: HttpClient,
        private walletService: WalletService,
        private fb: FormBuilder,
        private router: Router
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

    
    addWallet(currency, balance, Customer, Transactions, Status): void {
        this.walletService
        .addWallet(currency, balance, Customer, Transactions, Status)
            .subscribe(() => {
                this.router.navigate(['/indexWallet']);
            });
    }

    ngOnInit(): void {
    }
}