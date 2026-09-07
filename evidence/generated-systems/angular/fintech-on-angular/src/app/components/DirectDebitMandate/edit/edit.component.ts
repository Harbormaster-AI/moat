import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DirectDebitMandateService } from '../../../services/DirectDebitMandate.service';
import { SubBaseComponent } from '../../DirectDebitMandate/sub.base.component';


@Component({
    selector: 'app-edit-directDebitMandate',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDirectDebitMandateComponent extends SubBaseComponent implements OnInit {

    title = 'Edit DirectDebitMandate';

    directDebitMandateForm: FormGroup;
    directDebitMandate: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DirectDebitMandateService,
        private fb: FormBuilder
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

    
    updateDirectDebitMandate(mandateId, signedAt, Account, Creditor, Scheme, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDirectDebitMandate(mandateId, signedAt, Account, Creditor, Scheme, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDirectDebitMandate']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDirectDebitMandate(params['id']).subscribe(res => {
                this.directDebitMandate = res;
            });
        });
    }
}