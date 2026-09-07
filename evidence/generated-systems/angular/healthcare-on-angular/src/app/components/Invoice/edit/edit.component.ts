import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InvoiceService } from '../../../services/Invoice.service';
import { SubBaseComponent } from '../../Invoice/sub.base.component';


@Component({
    selector: 'app-edit-invoice',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInvoiceComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Invoice';

    invoiceForm: FormGroup;
    invoice: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InvoiceService,
        private fb: FormBuilder
) {
        super(http);
        this.invoiceForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  invoiceNumber: ['', Validators.required],
      totalAmount: ['', Validators.required],
      dueDate: ['', Validators.required],
      Patient: ['', ],
      Claim: ['', ],
      Payments: ['', ],
      Status: ['', ]
        });
    }

    
    updateInvoice(invoiceNumber, totalAmount, dueDate, Patient, Claim, Payments, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInvoice(invoiceNumber, totalAmount, dueDate, Patient, Claim, Payments, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInvoice']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInvoice(params['id']).subscribe(res => {
                this.invoice = res;
            });
        });
    }
}