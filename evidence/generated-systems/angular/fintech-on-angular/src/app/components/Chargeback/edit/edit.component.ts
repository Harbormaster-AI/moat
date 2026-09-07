import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ChargebackService } from '../../../services/Chargeback.service';
import { SubBaseComponent } from '../../Chargeback/sub.base.component';


@Component({
    selector: 'app-edit-chargeback',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditChargebackComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Chargeback';

    chargebackForm: FormGroup;
    chargeback: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ChargebackService,
        private fb: FormBuilder
) {
        super(http);
        this.chargebackForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  chargebackReference: ['', Validators.required],
      amount: ['', Validators.required],
      postedAt: ['', Validators.required],
      Dispute: ['', ],
      Transaction: ['', ],
      Stage: ['', ],
      Status: ['', ]
        });
    }

    
    updateChargeback(chargebackReference, amount, postedAt, Dispute, Transaction, Stage, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateChargeback(chargebackReference, amount, postedAt, Dispute, Transaction, Stage, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexChargeback']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getChargeback(params['id']).subscribe(res => {
                this.chargeback = res;
            });
        });
    }
}