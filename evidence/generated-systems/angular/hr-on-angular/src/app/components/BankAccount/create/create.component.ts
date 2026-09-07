import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BankAccountService } from '../../../services/BankAccount.service';
import { BankAccount } from '../../../models/BankAccount';
import { SubBaseComponent } from '../../BankAccount/sub.base.component';

@Component({
    selector: 'app-create-bankAccount',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateBankAccountComponent extends SubBaseComponent implements OnInit {

    title = 'Add BankAccount';

    bankAccountForm: FormGroup;
    bankAccount: BankAccount;

    constructor( http: HttpClient,
        private bankAccountService: BankAccountService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.bankAccountForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  accountHolder: ['', Validators.required],
      bankName: ['', Validators.required],
      iban: ['', Validators.required],
      bic: ['', Validators.required],
      accountNumber: ['', Validators.required],
      routingNumber: ['', Validators.required]
        });
    }

    
    addBankAccount(accountHolder, bankName, iban, bic, accountNumber, routingNumber): void {
        this.bankAccountService
        .addBankAccount(accountHolder, bankName, iban, bic, accountNumber, routingNumber)
            .subscribe(() => {
                this.router.navigate(['/indexBankAccount']);
            });
    }

    ngOnInit(): void {
    }
}