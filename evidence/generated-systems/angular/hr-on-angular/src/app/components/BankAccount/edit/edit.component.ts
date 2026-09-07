import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { BankAccountService } from '../../../services/BankAccount.service';
import { SubBaseComponent } from '../../BankAccount/sub.base.component';


@Component({
    selector: 'app-edit-bankAccount',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditBankAccountComponent extends SubBaseComponent implements OnInit {

    title = 'Edit BankAccount';

    bankAccountForm: FormGroup;
    bankAccount: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: BankAccountService,
        private fb: FormBuilder
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

    
    updateBankAccount(accountHolder, bankName, iban, bic, accountNumber, routingNumber): void {
        this.route.params.subscribe((params) => {

                        this.service.updateBankAccount(accountHolder, bankName, iban, bic, accountNumber, routingNumber, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexBankAccount']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getBankAccount(params['id']).subscribe(res => {
                this.bankAccount = res;
            });
        });
    }
}