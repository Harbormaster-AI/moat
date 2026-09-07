import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PurchaseAgreementService } from '../../../services/PurchaseAgreement.service';
import { SubBaseComponent } from '../../PurchaseAgreement/sub.base.component';


@Component({
    selector: 'app-edit-purchaseAgreement',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPurchaseAgreementComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PurchaseAgreement';

    purchaseAgreementForm: FormGroup;
    purchaseAgreement: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PurchaseAgreementService,
        private fb: FormBuilder
) {
        super(http);
        this.purchaseAgreementForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  agreementNumber: ['', Validators.required],
      effectiveDate: ['', Validators.required],
      AircraftOrder: ['', ]
        });
    }

    
    updatePurchaseAgreement(agreementNumber, effectiveDate, AircraftOrder): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePurchaseAgreement(agreementNumber, effectiveDate, AircraftOrder, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPurchaseAgreement']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPurchaseAgreement(params['id']).subscribe(res => {
                this.purchaseAgreement = res;
            });
        });
    }
}