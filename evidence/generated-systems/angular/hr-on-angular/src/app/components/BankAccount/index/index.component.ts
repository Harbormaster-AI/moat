
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BankAccountService } from '../../../services/BankAccount.service';
import { BankAccount } from '../../../models/BankAccount';

@Component({
    selector: 'app-index-bankAccount',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexBankAccountComponent implements OnInit {

    bankAccounts: BankAccount[] = [];

    constructor(
        private router: Router,
        private service: BankAccountService
) {}

    ngOnInit(): void {
        this.getBankAccounts();
}

    getBankAccounts(): void {
        this.service.getBankAccounts().subscribe((res) => {
        this.bankAccounts = res;
    });
}

    deleteBankAccount(id: any): void {
        this.service.deleteBankAccount(id)
            .subscribe(() => {
                this.getBankAccounts();
            });
    }
}