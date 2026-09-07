import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DirectDebitMandateService } from '../../../services/DirectDebitMandate.service';
import { DirectDebitMandate } from '../../../models/DirectDebitMandate';
import { SubBaseComponent } from '../../DirectDebitMandate/sub.base.component';

@Component({
    selector: 'app-create-directDebitMandate',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDirectDebitMandateComponent extends SubBaseComponent implements OnInit {

    title = 'Add DirectDebitMandate';

    directDebitMandateForm: FormGroup;
    directDebitMandate: DirectDebitMandate;

    constructor( http: HttpClient,
        private directDebitMandateService: DirectDebitMandateService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.directDebitMandateForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  mandateId: ['', Validators.required],
      signedAt: ['', Validators.required],
      Account: ['', ],
      Creditor: ['', ],
      Scheme: ['', ],
      Status: ['', ]
        });
    }

    
    addDirectDebitMandate(mandateId, signedAt, Account, Creditor, Scheme, Status): void {
        this.directDebitMandateService
        .addDirectDebitMandate(mandateId, signedAt, Account, Creditor, Scheme, Status)
            .subscribe(() => {
                this.router.navigate(['/indexDirectDebitMandate']);
            });
    }

    ngOnInit(): void {
    }
}